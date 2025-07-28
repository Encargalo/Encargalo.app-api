package customers

import (
	dto "CaliYa/core/domain/dto/customers"
	"CaliYa/core/domain/ports/customers"
	ports "CaliYa/core/domain/ports/customers"
	"CaliYa/core/domain/ports/sessions"
	"CaliYa/core/utils"
	"context"
	"errors"

	"github.com/google/uuid"
)

type customersSessionsApp struct {
	svc        customers.CustomersApp
	sessionSVC sessions.SessionsApp
	pass       utils.Password
}

func NewCustomersSessionsApp(
	svc customers.CustomersApp,
	sessionSVC sessions.SessionsApp,
	pass utils.Password) ports.CustomersSessionsApp {
	return &customersSessionsApp{
		svc,
		sessionSVC,
		pass}
}

func (c *customersSessionsApp) Sign_In(ctx context.Context, sign_in dto.SignIn) (uuid.UUID, error) {

	criteria := dto.SearchCustomerBy{
		Phone: sign_in.PhoneNumber,
	}

	customer, err := c.svc.SearchCustomerBy(ctx, criteria)
	if err != nil {
		if err.Error() == "not found." {
			return uuid.Nil, errors.New("incorrect access data")
		}
		return uuid.Nil, err
	}

	if !(c.pass.CheckPasswordHash([]byte(customer.Password), sign_in.Password)) {
		return uuid.Nil, errors.New("incorrect access data")
	}

	sessionID, err := c.sessionSVC.RegisterSessions(ctx, customer.ID, "Client")
	if err != nil {
		return uuid.Nil, err
	}

	return sessionID, nil
}
