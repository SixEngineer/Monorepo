package main

import (
	"openbridge/backend/internal/handler"
	"openbridge/backend/internal/repository"
	"openbridge/backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	// 数据库连接
	db, err := gorm.Open(sqlite.Open("openbridge.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	tokenRepo := repository.NewTokenRepository(db)
	tokenUsecase := usecase.NewTokenUseCase(tokenRepo)
	baiduHandler := handler.NewTokenHandler(tokenUsecase)

	r := gin.Default()

	tokenGroup := r.Group("/api/v1/token")
	{
		tokenGroup.POST("/access", baiduHandler.UploadAccessToken)
		tokenGroup.POST("/refresh", baiduHandler.UploadRefreshToken)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}