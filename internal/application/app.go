package application

import "github.com/dwnGnL/validation/internal/repository"

type Core interface {
	TestService() string
}

type RegistrationCore interface {
	Registration(users *repository.Users) error
}

type LoginCore interface {
	Login(users *repository.Users) (string, error)
}
