package main

import (
	"openbridge/backend/internal/config"
	"openbridge/backend/internal/domain/entity"
	"openbridge/backend/internal/handler"
	"openbridge/backend/internal/repository"
	"openbridge/backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	// 读取配置
	allConfig := config.ReadConfig()

	// 数据库连接
	db, err := gorm.Open(sqlite.Open(allConfig.DB.Path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移
	err = db.AutoMigrate(&entity.Token{})
	if err != nil {
		panic("failed to migrate database")
	}

	tokenRepo := repository.NewTokenRepository(db)
	tokenUsecase := usecase.NewTokenUseCase(tokenRepo)
	baiduHandler := handler.NewTokenHandler(tokenUsecase)

	r := gin.Default()

	tokenGroup := r.Group("/api/v1/token")
	{
		tokenGroup.POST("", baiduHandler.UploadToken)
	}

	r.Run(":" + allConfig.App.Port)
}