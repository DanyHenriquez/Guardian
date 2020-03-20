package main

import (
	"Guardian/database"
	"Guardian/handlers"
	"Guardian/models"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	db := database.Connect()

	defer func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}()

	db.AutoMigrate(&models.User{})

	userHandler := handlers.UserHandler{}
	userHandler.Database = db

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(os.Getenv("secret")),
		TokenLookup: "header:token",
	}))

	userGroup := e.Group("/api/user")
	userGroup.GET("", userHandler.Index)
	userGroup.GET("/:id", userHandler.Get)
	userGroup.POST("", userHandler.Post)
	userGroup.PUT(":id", userHandler.Put)
	userGroup.DELETE(":id", userHandler.Delete)

	e.Static("/", "./public/dist")
	e.Logger.Debug(e.Start(":1323"))
}
