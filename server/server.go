package server

import (
	"prod/config"
	"prod/docs"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	gin    *gin.Engine
	config config.Config
}

func NewServer(config config.Config) *Server {
	return &Server{
		gin:    gin.Default(),
		config: config,
	}
}

func (s *Server) Run() {
	docs.SwaggerInfo.Title = "Prod API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	v1 := s.gin.Group("/api/v1")
	{
		// PingExample godoc
		// @Summary ping example
		// @Schemes
		// @Description do ping
		// @Tags example
		// @Accept json
		// @Produce json
		// @Success 200 {string} Helloworld
		// @Router /example/helloworld [get]
		v1.GET("/ping", func(c *gin.Context) {

		})
	}
	// version 1
	v1.Group("/user")
	{
		v1.POST("/", func(c *gin.Context) {})
		v1.GET("/", func(c *gin.Context) {})
		v1.GET("/:id", func(c *gin.Context) {})
		v1.PUT("/:id", func(c *gin.Context) {})
		v1.DELETE("/:id", func(c *gin.Context) {})
		v1.PATCH("/:id", func(c *gin.Context) {})
		v1.OPTIONS("/:id", func(c *gin.Context) {})

	}

	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	s.gin.Run(s.config.ServerAddress)

}
