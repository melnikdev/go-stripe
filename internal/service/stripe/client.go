package stripe

import (
	"encoding/json"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/customer"
	"github.com/stripe/stripe-go/v81/paymentintent"
	"github.com/stripe/stripe-go/v81/price"
	"github.com/stripe/stripe-go/v81/product"
	"github.com/stripe/stripe-go/v81/subscription"
	"github.com/stripe/stripe-go/v81/webhook"
)

type Client struct {
	secretKey     string
	webhookSecret string
}

func NewClient(secretKey, webhookSecret string) *Client {
	stripe.Key = secretKey
	return &Client{
		secretKey:     secretKey,
		webhookSecret: webhookSecret,
	}
}

func (c *Client) HandleWebhook(payload []byte, signature string) (*stripe.Event, error) {
	event, err := webhook.ConstructEvent(payload, signature, c.webhookSecret)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (c *Client) ParseSubscriptionEvent(event *stripe.Event) (interface{}, error) {
	switch event.Type {
	case "customer.subscription.created":
		var subscription stripe.Subscription
		err := json.Unmarshal(event.Data.Raw, &subscription)
		return &subscription, err

	case "customer.subscription.updated":
		var subscription stripe.Subscription
		err := json.Unmarshal(event.Data.Raw, &subscription)
		return &subscription, err

	case "customer.subscription.deleted":
		var subscription stripe.Subscription
		err := json.Unmarshal(event.Data.Raw, &subscription)
		return &subscription, err

	default:
		return nil, nil
	}
}

func (c *Client) CreateCustomer(email, name string) (*stripe.Customer, error) {
	params := &stripe.CustomerParams{
		Email: stripe.String(email),
		Name:  stripe.String(name),
	}
	return customer.New(params)
}

func (c *Client) SubscribeCustomerToPrice(customerID, priceID string) (*stripe.Subscription, error) {
	params := &stripe.SubscriptionParams{
		Customer: stripe.String(customerID),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String(priceID),
			},
		},
		PaymentBehavior: stripe.String("default_incomplete"),
	}
	// params.AddExpand("pending_setup_intent")
	return subscription.New(params)
}

func (c *Client) CreatePaymentIntent(amount int64, currency string, customerID string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(currency),
		Customer: stripe.String(customerID),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	return paymentintent.New(params)
}

func (c *Client) CreateProduct(name string, price int64) (*stripe.Product, error) {
	params := &stripe.ProductParams{
		Name: stripe.String(name),
		DefaultPriceData: &stripe.ProductDefaultPriceDataParams{
			UnitAmount: stripe.Int64(price),
			Currency:   stripe.String(string(stripe.CurrencyUSD)),
			Recurring: &stripe.ProductDefaultPriceDataRecurringParams{
				Interval: stripe.String("month"),
			},
		},
	}
	return product.New(params)

}
func (c *Client) CreatePrice(productID string, amount int64, currency string) (*stripe.Price, error) {
	params := &stripe.PriceParams{
		Product:    stripe.String(productID),
		Currency:   stripe.String(currency),
		UnitAmount: stripe.Int64(amount),
	}
	return price.New(params)
}

func (c *Client) GetCustomer(customerID string) (*stripe.Customer, error) {
	return customer.Get(customerID, nil)
}

func (c *Client) GetPaymentIntent(paymentIntentID string) (*stripe.PaymentIntent, error) {
	return paymentintent.Get(paymentIntentID, nil)
}

func (c *Client) GetProduct(productID string) (*stripe.Product, error) {
	return product.Get(productID, nil)
}

func (c *Client) GetPrice(priceID string) (*stripe.Price, error) {
	return price.Get(priceID, nil)
}
