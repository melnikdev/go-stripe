package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/melnikdev/go-stripe/internal/config"
	"github.com/melnikdev/go-stripe/internal/controller"
	"github.com/melnikdev/go-stripe/internal/service"
	"github.com/melnikdev/go-stripe/internal/service/stripe"
	"gorm.io/gorm"
)

type Server struct {
	config *config.Config
	db     *gorm.DB
}

func NewServer(config *config.Config, db *gorm.DB) *http.Server {
	NewServer := &Server{
		config: config,
		db:     db,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.config.Server.Port),
		Handler:      NewServer.initServer(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func (s *Server) initServer() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	productService := service.NewProductService(s.db)
	client := stripe.NewClient(s.config.Stripe.SecretKey)
	productController := controller.NewProductController(productService, client)

	e.POST("/products", productController.Create)
	e.GET("/products", productController.GetAll)
	e.GET("/products/:id", productController.GetById)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return e
}
