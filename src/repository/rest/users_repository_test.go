package rest

import (
	"net/http"
	"os"
	"testing"

	"github.com/laithrafid/bookstore_oauth-api/src/utils/config_utils"
	"github.com/laithrafid/bookstore_oauth-api/src/utils/logger_utils"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	config, err := config_utils.LoadConfig(".")
	if err != nil {
		logger_utils.Error("cannot load config of user_repository_test:", err)
	}
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          config.UsersApiAddress + "/users/login",
		ReqBody:      `{"email":"laith.rafid@gmail.com","password":"New-password"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("laith.rafid@gmail.com", "New-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err)
	assert.EqualValues(t, "invalid restclient response when trying to login user", err)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	config, err := config_utils.LoadConfig(".")
	if err != nil {
		logger_utils.Error("cannot load config of user_repository_test:", err)
	}
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          config.UsersApiAddress + "/users/login",
		ReqBody:      `{"email":"laith.rafid@gmail.com","password":"New-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials", "status": "404", "error": "not_found"}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("laith.rafid@gmail.com", "New-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err)
	assert.EqualValues(t, "invalid error interface when trying to login user", err)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	config, err := config_utils.LoadConfig(".")
	if err != nil {
		logger_utils.Error("cannot load config of user_repository_test:", err)
	}
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          config.UsersApiAddress + "/users/login",
		ReqBody:      `{"email":"laith.rafid@gmail.com","password":"New-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials", "status": 404, "error": "not_found"}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("laith.rafid@gmail.com", "New-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err)
	assert.EqualValues(t, "invalid login credentials", err)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	config, err := config_utils.LoadConfig(".")
	if err != nil {
		logger_utils.Error("cannot load config of user_repository_test:", err)
	}
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          config.UsersApiAddress + "/users/login",
		ReqBody:      `{"email":"laith.rafid@gmail.com","password":"New-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 1, "first_name": "laith", "last_name": "ahmed", "email": "laith.rafid@gmail.com"}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("laith.rafid@gmail.com", "New-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err)
	assert.EqualValues(t, "error when trying to unmarshal users login response", err)
}

func TestLoginUserNoError(t *testing.T) {
	config, err := config_utils.LoadConfig(".")
	if err != nil {
		logger_utils.Error("cannot load config of user_repository_test:", err)
	}
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          config.UsersApiAddress + "/users/login",
		ReqBody:      `{"email":"laith.rafid@gmail.com","password":"New-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 1, "first_name": "laith", "last_name": "ahmed", "email": "laith.rafid@gmail.com"}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("laith.rafid@gmail.com", "New-password")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "laith", user.FirstName)
	assert.EqualValues(t, "ahmed", user.LastName)
	assert.EqualValues(t, "laith.rafid@gmail.com", user.Email)
}
