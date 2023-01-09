package event

import "time"

type OrderCreated struct {
	Name    string
	Payload interface{}
}

func NewOrderCreated() *OrderCreated {
	return &OrderCreated{
		Name: "OrderCreated",
	}
}

func (e *OrderCreated) GetName() string {
	return e.Name
}

func (e *OrderCreated) GetPayload() interface{} {
	return e.Payload
}

func (e *OrderCreated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *OrderCreated) GetDateTime() time.Time {
	return time.Now()
}
