package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims = 标准声明 + 你的自定义数据
type Claims struct {
	jwt.RegisteredClaims
	Data map[string]any `json:"data,omitempty"`
}

// Util 持有你签名用的密钥与简单配置
type Util struct {
	secret []byte
	issuer string
	leeway time.Duration // 时钟漂移容忍，默认 30s
}

// Option 可选配置
type Option func(*Util)

func WithIssuer(iss string) Option      { return func(u *Util) { u.issuer = iss } }
func WithLeeway(d time.Duration) Option { return func(u *Util) { u.leeway = d } }

// New：只需要一个 HS256 密钥，其他都可选
func New(secret string, opts ...Option) *Util {
	u := &Util{
		secret: []byte(secret),
		leeway: 30 * time.Second,
	}
	for _, opt := range opts {
		opt(u)
	}
	return u
}

// Sign：签发一个带自定义数据的 Token
// subject：放用户ID等标识；ttl：有效期；data：自定义载荷（role、scopes等）
func (u *Util) Sign(subject string, ttl time.Duration, data map[string]any) (string, error) {
	if ttl <= 0 {
		return "", errors.New("ttl must be > 0")
	}
	now := time.Now()
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    u.issuer,
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now.Add(-u.leeway)),
		},
		Data: data,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(u.secret)
}

// Verify：校验 Token（签名、过期时间、算法、可选发行方），返回 Claims
func (u *Util) Verify(tokenStr string) (*Claims, error) {
	parser := jwt.NewParser(
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithLeeway(u.leeway),
	)

	claims := &Claims{}
	_, err := parser.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		// 防算法混淆
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %s", t.Method.Alg())
		}
		return u.secret, nil
	})
	if err != nil {
		return nil, err
	}
	if u.issuer != "" && claims.Issuer != u.issuer {
		return nil, errors.New("invalid issuer")
	}
	return claims, nil
}

// Data：直接从 Token 里拿自定义数据（内部会做一次 Verify）
func (u *Util) Data(tokenStr string) (map[string]any, error) {
	c, err := u.Verify(tokenStr)
	if err != nil {
		return nil, err
	}
	if c.Data == nil {
		return map[string]any{}, nil
	}
	return c.Data, nil
}
