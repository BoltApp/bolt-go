// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/BoltApp/bolt-go/internal/utils"
)

type TokenRequestType string

const (
	TokenRequestTypeAuthorizationCodeRequest TokenRequestType = "authorization-code-request"
	TokenRequestTypeRefreshTokenRequest      TokenRequestType = "refresh-token-request"
)

type TokenRequest struct {
	AuthorizationCodeRequest *AuthorizationCodeRequest `queryParam:"inline"`
	RefreshTokenRequest      *RefreshTokenRequest      `queryParam:"inline"`

	Type TokenRequestType
}

func CreateTokenRequestAuthorizationCodeRequest(authorizationCodeRequest AuthorizationCodeRequest) TokenRequest {
	typ := TokenRequestTypeAuthorizationCodeRequest

	return TokenRequest{
		AuthorizationCodeRequest: &authorizationCodeRequest,
		Type:                     typ,
	}
}

func CreateTokenRequestRefreshTokenRequest(refreshTokenRequest RefreshTokenRequest) TokenRequest {
	typ := TokenRequestTypeRefreshTokenRequest

	return TokenRequest{
		RefreshTokenRequest: &refreshTokenRequest,
		Type:                typ,
	}
}

func (u *TokenRequest) UnmarshalJSON(data []byte) error {

	var authorizationCodeRequest AuthorizationCodeRequest = AuthorizationCodeRequest{}
	if err := utils.UnmarshalJSON(data, &authorizationCodeRequest, "", true, true); err == nil {
		u.AuthorizationCodeRequest = &authorizationCodeRequest
		u.Type = TokenRequestTypeAuthorizationCodeRequest
		return nil
	}

	var refreshTokenRequest RefreshTokenRequest = RefreshTokenRequest{}
	if err := utils.UnmarshalJSON(data, &refreshTokenRequest, "", true, true); err == nil {
		u.RefreshTokenRequest = &refreshTokenRequest
		u.Type = TokenRequestTypeRefreshTokenRequest
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for TokenRequest", string(data))
}

func (u TokenRequest) MarshalJSON() ([]byte, error) {
	if u.AuthorizationCodeRequest != nil {
		return utils.MarshalJSON(u.AuthorizationCodeRequest, "", true)
	}

	if u.RefreshTokenRequest != nil {
		return utils.MarshalJSON(u.RefreshTokenRequest, "", true)
	}

	return nil, errors.New("could not marshal union type TokenRequest: all fields are null")
}
