package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// Config 应用配置结构
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	JWT      JWTConfig      `yaml:"jwt"`
	Log      LogConfig      `yaml:"log"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string `yaml:"port"`
	Mode string `yaml:"mode"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	DBName          string `yaml:"dbname"`
	Charset         string `yaml:"charset"`
	ParseTime       bool   `yaml:"parse_time"`
	Loc             string `yaml:"loc"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	SecretKey   string `yaml:"secret_key"`
	ExpireHours int    `yaml:"expire_hours"`
	Issuer      string `yaml:"issuer"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `yaml:"level"`
	FilePath   string `yaml:"file_path"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
}

var AppConfig *Config

// LoadConfig 加载配置文件
func LoadConfig(configPath string) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	config := &Config{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return fmt.Errorf("解析配置文件失败: %v", err)
	}

	AppConfig = config
	return nil
}

// GetDSN 获取数据库连接字符串
func (db *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.DBName,
		db.Charset,
		db.ParseTime,
		db.Loc,
	)
}
