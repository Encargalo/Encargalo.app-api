package customers

import "github.com/google/uuid"

type SearchCustomerBy struct {
	ID    uuid.UUID `query:"customer_id"`
	Phone string    `query:"phone"`
}
