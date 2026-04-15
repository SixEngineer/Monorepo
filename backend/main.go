package main

import (
	"openbridge/backend/internal/config"
	"openbridge/backend/internal/domain/entity"
	"openbridge/backend/internal/handler"
	"openbridge/backend/internal/middleware"
	"openbridge/backend/internal/pkg/logger"
	"openbridge/backend/internal/repository"
	"openbridge/backend/internal/tool"
	"openbridge/backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	// 读取配置
	allConfig := config.ReadConfig()
	if err := logger.Init(allConfig.Log.Level, allConfig.Log.Format); err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.L().Info("service starting",
		zap.String("app", allConfig.App.Name),
		zap.String("env", allConfig.App.Env),
		zap.String("port", allConfig.App.Port),
	)

	// 数据库连接
	db, err := gorm.Open(sqlite.Open(allConfig.DB.Path), &gorm.Config{})
	if err != nil {
		logger.L().Fatal("db connect failed", zap.Error(err), zap.String("db_path", allConfig.DB.Path))
	}

	// 自动迁移
	err = db.AutoMigrate(&entity.Token{})
	if err != nil {
		logger.L().Fatal("db migrate failed", zap.Error(err))
	}

	err = db.AutoMigrate(&entity.ProviderAccount{})
	if err != nil {
		logger.L().Fatal("db migrate failed", zap.Error(err))
	}

	tokenRepo := repository.NewTokenRepository(db)
	tokenUsecase := usecase.NewTokenUseCase(tokenRepo)
	tokenHandler := handler.NewTokenHandler(tokenUsecase)

	providerRepo := repository.NewProviderRepository(db)
	providerRegistry := tool.NewRegistry()
	providerUsecase := usecase.NewProviderUseCase(providerRepo, providerRegistry)
	providerHandler := handler.NewProviderHandler(providerUsecase)

	r := gin.New()
	r.Use(middleware.RequestID())
	r.Use(middleware.AccessLog())
	r.Use(gin.Recovery())

	tokenGroup := r.Group("/api/v1/token")
	{
		tokenGroup.POST("", tokenHandler.UploadToken)
	}

	// 注册 Provider 相关路由
	providerGroup := r.Group("/api/v1/provider")
	{
		providerGroup.POST("", providerHandler.RegisterProvider)
		providerGroup.DELETE("", providerHandler.DeleteProvider)
		providerGroup.PUT("/", providerHandler.UpdateProvider)
		providerGroup.GET("/info", providerHandler.GetProvider)
		providerGroup.GET("/list", providerHandler.ListProvider)
	}

	if err := r.Run(":" + allConfig.App.Port); err != nil {
		logger.L().Fatal("http server run failed", zap.Error(err))
	}
}
