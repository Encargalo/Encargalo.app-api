package customers

import (
	dto "CaliYa/core/domain/dto/customers"
	"context"

	"github.com/google/uuid"
)

type CustomersSessionsApp interface {
	Sign_In(ctx context.Context, sign_in dto.SignIn) (uuid.UUID, error)
}
