package request

type CreateSubscriptionRequest struct {
	CustomerID string `validate:"required,min=1,max=200" json:"customer_id"`
	PriceID    string `validate:"required,min=1,max=200" json:"price_id"`
}
