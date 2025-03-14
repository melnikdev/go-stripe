package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/melnikdev/go-stripe/internal/request"
	"github.com/melnikdev/go-stripe/internal/service"
	"github.com/melnikdev/go-stripe/internal/service/stripe"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (c *ProductController) Create(ctx echo.Context) error {
	request := request.CreateProductRequest{}
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	product, err := c.productService.Create(request)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	client := stripe.NewClient("sk_test_51")
	_, err = client.CreateProduct(product.Name, int64(product.Price))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, "Product created successfully")
}
