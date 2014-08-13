package auth

import (
	"fmt"
	"net/url"

	"github.com/motemen/go-pocket/api"
)

type OAuthRequestAPIResponse struct {
	Code string `json:"code"`
}

type OAuthAuthorizeAPIResponse struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
}

func ObtainRequestToken(consumerKey, redirectURL string) (*OAuthRequestAPIResponse, error) {
	res := &OAuthRequestAPIResponse{}
	err := api.RequestJSON(
		"/v3/oauth/request",
		map[string]string{
			"consumer_key": consumerKey,
			"redirect_uri": redirectURL,
		},
		res,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func ObtainAccessToken(consumerKey string, requestToken *OAuthRequestAPIResponse) (*OAuthAuthorizeAPIResponse, error) {
	res := &OAuthAuthorizeAPIResponse{}
	err := api.RequestJSON(
		"/v3/oauth/authorize",
		map[string]string{
			"consumer_key": consumerKey,
			"code":         requestToken.Code,
		},
		res,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GenerateAuthorizationURL(requestToken *OAuthRequestAPIResponse, redirectURL string) string {
	values := url.Values{"request_token": {requestToken.Code}, "redirect_uri": {redirectURL}}
	return fmt.Sprintf("%s/auth/authorize?%s", api.Origin, values.Encode())
}
