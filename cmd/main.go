package main

import (
	"github.com/KiranPawar0/coditas-test/pkg/helper/structvalidator"
	"github.com/KiranPawar0/coditas-test/pkg/middleware"
	"github.com/KiranPawar0/coditas-test/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		structvalidator.RegisterCustomValidations(v)
	}
	r.Use(middleware.LatencyLogger())

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
