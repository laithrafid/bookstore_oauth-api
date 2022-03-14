package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/laithrafid/utils-go/crypto_utils"
	"github.com/laithrafid/utils-go/errors_utils"
)

//TODO: Swagger API Documentation
// Need different Client ID to Give different results depends on frontend APP
// Web frontend App: ClientId: 123
// Android APP: ClientId: 234
// IOS APP: ClientId: 234
// Users API:
// {
// "email": "emailaddress@email.com",
// "paasword": "123abc"
// }
// Oauth API:
// {
//	"grant_type": "password",
//	"email": "emailaddress@email.com",
//	"paasword": "123abc"
// }
// {
//	"grant_type": "client_credentials",
//	"client_id": "id-123",
//	"client_secret": "secret-123"
// }
const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grandTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for client_credentials grant type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() errors_utils.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		break

	case grandTypeClientCredentials:
		break

	default:
		return errors_utils.NewBadRequestError("invalid grant_type parameter")
	}

	//TODO: Validate parameters for each grant_type
	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() errors_utils.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors_utils.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors_utils.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors_utils.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors_utils.NewBadRequestError("invalid expiration time")
	}
	return nil
}
func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
