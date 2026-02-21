package domain

import "errors"


var (
	ErrCustomerNotFound = errors.New("Customer not found")
	ErrCustomerAlreadyExiting = errors.New("Customer already exiting")
	
)
