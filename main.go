package main

import (
	"github.com/mariogao/golang-unit-test/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/mariogao/golang-unit-test/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	/*
		Or can use EchoWrapHandler func with configurations.
		url := echoSwagger.URL("http://localhost:1323/swagger/doc.json") //The url pointing to API definition
		e.GET("/swagger/*", echoSwagger.EchoWrapHandler(url))
	*/

	// Routes
	userController := controller.NewUserControl()
	e.GET("/get", userController.GetUser)
	e.POST("/update", userController.UpdateUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
