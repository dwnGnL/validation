package service

import (
	"github.com/dwnGnL/validation/internal/config"
	"github.com/dwnGnL/validation/internal/repository"
)

type repositoryIter interface {
	GetTestTable() (*repository.TestTable, error)
}

type ServiceImpl struct {
	conf *config.Config
	repo repositoryIter
}

type Option func(*ServiceImpl)

func New(conf *config.Config, repo repositoryIter, opts ...Option) *ServiceImpl {
	s := ServiceImpl{
		conf: conf,
		repo: repo,
	}

	for _, opt := range opts {
		opt(&s)
	}

	return &s
}

func (s ServiceImpl) TestService() string {
	return "it`s test"
}
