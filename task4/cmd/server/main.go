package main

import (
	"log"
	"time"
	"web3_study_task/task4/config"
	"web3_study_task/task4/internal/handler"
	"web3_study_task/task4/internal/middleware"
	"web3_study_task/task4/internal/repository"
	"web3_study_task/task4/internal/service"
	"web3_study_task/task4/pkg/entity"
	"web3_study_task/task4/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// 加载配置
	if err := config.LoadConfig("D:/code/go/web3_study_task/task4/config/config.yaml"); err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}

	// 初始化数据库
	db, err := initDB()
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	// 自动迁移
	if err := db.AutoMigrate(&entity.User{}, &entity.Post{}, &entity.Comment{}); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	// 初始化JWT工具
	jwtUtil := utils.New(
		config.AppConfig.JWT.SecretKey,
		utils.WithIssuer(config.AppConfig.JWT.Issuer),
	)

	// 初始化仓库层
	userRepo := repository.NewUserRepository(db)
	postRepo := repository.NewPostRepository(db)
	commentRepo := repository.NewCommentRepository(db)

	// 初始化服务层
	userService := service.NewUserService(userRepo, jwtUtil)
	postService := service.NewPostService(postRepo, userRepo)
	commentService := service.NewCommentService(commentRepo)

	// 初始化处理器层
	userHandler := handler.NewUserHandler(userService)
	postHandler := handler.NewPostHandler(postService)
	commentHandler := handler.NewCommentHandler(commentService)

	// 初始化路由
	router := setupRouter(userHandler, postHandler, commentHandler, jwtUtil)

	// 启动服务器
	log.Printf("服务器启动在端口 %s", config.AppConfig.Server.Port)
	if err := router.Run(config.AppConfig.Server.Port); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}

// initDB 初始化数据库连接
func initDB() (*gorm.DB, error) {
	dbConfig := config.AppConfig.Database

	db, err := gorm.Open(mysql.Open(dbConfig.GetDSN()), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束
	})
	if err != nil {
		return nil, err
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifetime) * time.Second)

	return db, nil
}

// setupRouter 设置路由
func setupRouter(userHandler *handler.UserHandler, postHandler *handler.PostHandler, commentHandler *handler.CommentHandler, jwtUtil *utils.Util) *gin.Engine {
	// 设置Gin模式
	gin.SetMode(config.AppConfig.Server.Mode)

	router := gin.Default()
	//所有api异常统一处理
	api := router.Group("/api", middleware.ErrHandler())

	// 无需登录认证的接口
	unAuthGroup := api.Group("/")
	{
		//用户登录
		unAuthGroup.POST("/user/login", userHandler.Login)
		//文章查询
		unAuthGroup.GET("/post/:id", postHandler.GetPost)
		unAuthGroup.GET("/post/list", postHandler.GetAllPosts)
		//评论相关
		unAuthGroup.GET("/comment/all/:postId", commentHandler.GetAllByPostId)
	}

	// 需要认证的接口
	authGroup := api.Group("/", middleware.AuthMiddleware(jwtUtil))
	{
		// 用户相关接口
		authGroup.GET("/user/getUserInfo", userHandler.GetUserInfo)
		authGroup.GET("/hello", userHandler.Hello)

		// 文章相关接口
		authGroup.POST("/post/create", postHandler.CreatePost)
		authGroup.GET("/post/my", postHandler.GetUserPosts)
		authGroup.POST("/post/update", postHandler.UpdatePost)
		authGroup.POST("/post/delete/:id", postHandler.DeletePost)

		//评论相关
		authGroup.POST("/comment/publish", commentHandler.PublishComment)
	}

	return router
}
