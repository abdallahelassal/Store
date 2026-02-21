package repository

import (
	"context"

	"github.com/abdallahelassal/Store/internal/modules/customer/domain"
)

type CustomerRepository interface {
	// Define methods for customer repository here
	CreateCustomer(ctx context.Context,customer *domain.Customer)error
	DeleteCustomer(ctx context.Context,customerUUID string)error
	GetCustomerByName(ctx context.Context, name string)(*domain.Customer,error)
	GetCustomers(ctx context.Context,limit,offset int)(*[]domain.Customer,int64,error)
}