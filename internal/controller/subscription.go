package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/melnikdev/go-stripe/internal/request"
	"github.com/melnikdev/go-stripe/internal/service"
	"github.com/melnikdev/go-stripe/internal/service/stripe"
)

type SubscriptionController struct {
	subscriptionService service.SubscriptionService
	stripeClient        *stripe.Client
	validate            *validator.Validate
}

func NewSubscriptionController(subscriptionService service.SubscriptionService, stripeClient *stripe.Client) *SubscriptionController {
	return &SubscriptionController{
		subscriptionService: subscriptionService,
		stripeClient:        stripeClient,
		validate:            validator.New(),
	}
}

func (c *SubscriptionController) Create(ctx echo.Context) error {
	request := request.CreateSubscriptionRequest{}
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.validate.Struct(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	subscription, err := c.stripeClient.SubscribeCustomerToPrice(request.CustomerID, request.PriceID)

	//add to db
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, subscription)
}
