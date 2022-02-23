package db

import (
	"github.com/laithrafid/bookstore_oauth-api/src/domain/access_token"
	"github.com/laithrafid/bookstore_oauth-api/utils/errors_utils"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors_utils.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors_utils.RestErr) {
	return nil, nil
}
