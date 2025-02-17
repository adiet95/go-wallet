package auth

import (
	models "go-wallet/src/models/entity"

	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock mock.Mock
}

func (m *RepoMock) FindByPhone(email string) (*models.User, error) {
	args := m.mock.Called(email)
	return args.Get(0).(*models.User), nil
}

func (m *RepoMock) RegisterPhone(data *models.User) (*models.User, error) {
	args := m.mock.Called(data)
	return args.Get(0).(*models.User), nil
}
