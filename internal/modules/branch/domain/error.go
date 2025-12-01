package domain
import (
	
	"errors"
)

var (
		ErrBranchNotFound = errors.New("Branch not found")
		ErrBranchAlreadyExiting = errors.New("Branch already exiting")
)