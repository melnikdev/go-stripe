package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/melnikdev/go-stripe/internal/request"
	"github.com/melnikdev/go-stripe/internal/service"
	"github.com/melnikdev/go-stripe/internal/service/stripe"
)

type CustomerController struct {
	customerService service.CustomerService
	client          *stripe.Client
	validate        *validator.Validate
}

func NewCustomerController(customerService service.CustomerService, client *stripe.Client) *CustomerController {
	return &CustomerController{
		customerService: customerService,
		client:          client,
		validate:        validator.New(),
	}
}

func (c *CustomerController) Create(ctx echo.Context) error {
	request := request.CreateCustomerRequest{}
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	err := c.validate.Struct(request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	stripeCustomer, err := c.client.CreateCustomer(request.Email, request.Name)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	customer, err := c.customerService.CreateCustomer(request, stripeCustomer.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, customer)
}
