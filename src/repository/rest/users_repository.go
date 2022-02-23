package rest

import (
	"github.com/laithrafid/bookstore_oauth-api/src/utils/config_utils"
	"github.com/laithrafid/bookstore_oauth-api/src/utils/errors_utils"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: config_utils.ServerAddress//myapiusers,
		Timeout: 100 * time.Millisecond,
	}
)
type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, errors_utils.RestErr)
}

type usersRepository struct{}

func NewRestUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, errors_utils.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response := usersRestClient.Post("/users/login", request)

	if response == nil || response.Response == nil {
		return nil, errors_utils.NewInternalServerError("invalid restclient response when trying to login user", errors.New("restclient error"))
	}

	if response.StatusCode > 299 {
		apiErr, err := errors_utils.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, errors_utils.NewInternalServerError("invalid error interface when trying to login user", err)
		}
		return nil, apiErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors_utils.NewInternalServerError("error when trying to unmarshal users login response", errors.New("json parsing error"))
	}
	return &user, nil
}