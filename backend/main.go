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
	"os"
	"path/filepath"

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

	// TODO: 创建数据库所在目录
	dbDir := filepath.Dir(allConfig.DB.Path)
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		logger.L().Fatal("failed to create db directory", zap.Error(err), zap.String("dir", dbDir))
	}

	// 数据库连接
	db, err := gorm.Open(sqlite.Open(allConfig.DB.Path), &gorm.Config{})
	if err != nil {
		logger.L().Fatal("db connect failed", zap.Error(err), zap.String("db_path", allConfig.DB.Path))
	}

	// 自动迁移数据库表
	err = db.AutoMigrate(&entity.ProviderAccount{})
	if err != nil {
		logger.L().Fatal("db migrate failed", zap.Error(err))
	}

	err = db.AutoMigrate(&entity.QuotaSnapshot{})
	if err != nil {
		logger.L().Fatal("db migrate failed", zap.Error(err))
	}

	err = db.AutoMigrate(&entity.MountPoint{})
	if err != nil {
		logger.L().Fatal("db migrate failed", zap.Error(err))
	}

	quotaRepo := repository.NewQuotaRepository(db)

	providerRegistry := tool.NewRegistry()

	// 初始化 Provider相关组件
	providerRepo := repository.NewProviderRepository(db)
	providerUsecase := usecase.NewProviderUseCase(providerRepo, providerRegistry)
	providerHandler := handler.NewProviderHandler(providerUsecase)

	// 初始化 Mount相关组件
	mountRepo := repository.NewMountRepository(db)
	mountUsecase := usecase.NewMountUseCase(mountRepo, providerRepo, quotaRepo, providerRegistry)
	mountHandler := handler.NewMountHandler(mountUsecase)

	// Gin引擎设置
	r := gin.New()
	r.Use(middleware.RequestID())
	r.Use(middleware.AccessLog())
	r.Use(gin.Recovery())

	// 注册 Provider 相关路由
	providerGroup := r.Group("/api/v1/provider")
	{
		providerGroup.POST("", providerHandler.RegisterProvider)
		providerGroup.DELETE("", providerHandler.DeleteProvider)
		providerGroup.PUT("/", providerHandler.UpdateProvider)
		providerGroup.GET("/info", providerHandler.GetProvider)
		providerGroup.GET("/list", providerHandler.ListProvider)
	}

	// 注册 Mount 相关路由
	mountGroup := r.Group("/api/v1/mount")
	{
		mountGroup.POST("", mountHandler.CreateMount)
		mountGroup.GET("/:id/quota", mountHandler.GetMountQuota)
		mountGroup.POST("/:id/quota/sync", mountHandler.SyncMountQuota)
	}

	if err := r.Run(":" + allConfig.App.Port); err != nil {
		logger.L().Fatal("http server run failed", zap.Error(err))
	}
}
