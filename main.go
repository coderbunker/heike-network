package main

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/coderbunker/heikenet-backend/controllers"
	mid "github.com/coderbunker/heikenet-backend/middleware"
	"github.com/coderbunker/heikenet-backend/models"
)

type app_config struct {
	Port         string `env:"PORT,required"`
	Database_url string `env:"DATABASE_URL,required"`
	Secret       string `env:"SECRET,required"`
	Key          string `env:"KEY,required"`
	Node         string `env:"NODE,required"`
	Dai          string `env:"DAI,required"`
	Symbol       string `env:"SYMBOL,required"`
	Retainer     string `env:"RETAINER,required"`
}

func main() {
	config := app_config{}
	err := env.Parse(&config)
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open("postgres", config.Database_url)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// this is for dev
	db.AutoMigrate(&models.User{})
	db.DropTable(&models.User{})
	db.CreateTable(&models.User{})

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(mid.SetSecret(config.Secret))
	e.Use(mid.SetDB(db))

	api := e.Group("/api/v1")
	api.POST("/login", controllers.Login)
	// api.POST("/contract/approve", controllers.Approve)
	// api.POST("/contract/fund", controllers.Fund)
	// api.POST("/contract/withdraw", controllers.Withdraw)

	users := api.Group("/users")
	users.POST("", controllers.CreateUser)
	// users.Use(middleware.JWT([]byte(secret)))
	users.GET("/:id", controllers.GetUser)
	users.PUT("/:id", controllers.UpdateUser)
	users.DELETE("/:id", controllers.DeleteUser)

	e.Logger.Fatal(e.Start(":" + config.Port))
}
