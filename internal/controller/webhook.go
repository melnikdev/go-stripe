package controller

import (
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/melnikdev/go-stripe/internal/service/stripe"
)

type WebhookController struct {
	stripeClient *stripe.Client
}

func NewWebhookController(stripeClient *stripe.Client) *WebhookController {
	return &WebhookController{
		stripeClient: stripeClient,
	}
}

type WebhookResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (c *WebhookController) HandleWebhook(ctx echo.Context) error {

	log.Println("Received webhook")
	payload, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, WebhookResponse{
			Status:  "error",
			Message: "Error reading request body",
		})
	}

	signature := ctx.Request().Header.Get("Stripe-Signature")
	if signature == "" {
		return ctx.JSON(http.StatusBadRequest, WebhookResponse{
			Status:  "error",
			Message: "No signature found",
		})
	}

	event, err := c.stripeClient.HandleWebhook(payload, signature)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, WebhookResponse{
			Status:  "error",
			Message: "Error verifying webhook signature",
		})
	}

	log.Println("Event:", event.Type)
	switch event.Type {
	case "customer.subscription.created":
		if _, err := c.stripeClient.ParseSubscriptionEvent(event); err != nil {
			return ctx.JSON(http.StatusBadRequest, WebhookResponse{
				Status:  "error",
				Message: "Error parsing subscription event",
			})
		}
		// TODO: Update database

	case "customer.subscription.updated":
		if _, err := c.stripeClient.ParseSubscriptionEvent(event); err != nil {
			return ctx.JSON(http.StatusBadRequest, WebhookResponse{
				Status:  "error",
				Message: "Error parsing subscription event",
			})
		}
		// TODO: Update database

	case "customer.subscription.deleted":
		if _, err := c.stripeClient.ParseSubscriptionEvent(event); err != nil {
			return ctx.JSON(http.StatusBadRequest, WebhookResponse{
				Status:  "error",
				Message: "Error parsing subscription event",
			})
		}
		// TODO: Update database

	default:
		return ctx.JSON(http.StatusOK, WebhookResponse{
			Status:  "ignored",
			Message: "Unhandled event type: " + string(event.Type),
		})
	}

	// Return success response
	return ctx.JSON(http.StatusOK, WebhookResponse{
		Status:  "success",
		Message: "Webhook processed successfully",
	})
}
