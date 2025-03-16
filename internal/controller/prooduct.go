package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/melnikdev/go-stripe/internal/request"
	"github.com/melnikdev/go-stripe/internal/service"
	"github.com/melnikdev/go-stripe/internal/service/stripe"
)

type ProductController struct {
	productService service.ProductService
	client         *stripe.Client
	validate       *validator.Validate
}

func NewProductController(productService service.ProductService, client *stripe.Client) *ProductController {
	return &ProductController{
		productService: productService,
		client:         client,
		validate:       validator.New(),
	}
}

func (c *ProductController) Create(ctx echo.Context) error {
	request := request.CreateProductRequest{}
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err := c.validate.Struct(request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	product, err := c.productService.Create(request)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	stripeProduct, err := c.client.CreateProduct(product.Name, int64(request.Price))

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	_, err = c.productService.UpdateStripeId(product, stripeProduct.ID)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	_, err = c.productService.CreatePrice(product.ID, stripeProduct.DefaultPrice.ID, request.Price)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, "Product created successfully")
}

func (c *ProductController) GetAll(ctx echo.Context) error {
	products, err := c.productService.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, products)
}

func (c *ProductController) GetById(ctx echo.Context) error {
	id := ctx.Param("id")
	product, err := c.productService.GetById(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, product)
}
