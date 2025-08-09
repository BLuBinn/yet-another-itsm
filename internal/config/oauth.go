package config

import (
	"fmt"
	"time"

	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"

	"yet-another-itsm/internal/constants"

	"github.com/MicahParks/keyfunc"
	"github.com/rs/zerolog/log"

	"golang.org/x/oauth2"
)

const microsoftLoginBaseURL = "https://login.microsoftonline.com/"

func initOAuth(oauthConfig *OAuthConfig) error {
	if oauthConfig.ClientID == "" {
		return fmt.Errorf(constants.ErrEntraClientIDRequiredMsg)
	}
	if oauthConfig.ClientSecret == "" {
		return fmt.Errorf(constants.ErrEntraClientSecretRequiredMsg)
	}
	if oauthConfig.TenantID == "" {
		return fmt.Errorf(constants.ErrEntraTenantIDRequiredMsg)
	}

	oauthConfig.EntraConfig = &oauth2.Config{
		ClientID:     oauthConfig.ClientID,
		ClientSecret: oauthConfig.ClientSecret,
		RedirectURL:  oauthConfig.RedirectURI,
		// Scopes:       []string{"openid", "profile", "offline_access", "https://graph.microsoft.com/User.Read"},
		Scopes: []string{"openid", "profile", "offline_access", "api://" + oauthConfig.ClientID + "/user_impersonation"},

		Endpoint: oauth2.Endpoint{
			AuthURL:  microsoftLoginBaseURL + oauthConfig.TenantID + "/oauth2/v2.0/authorize",
			TokenURL: microsoftLoginBaseURL + oauthConfig.TenantID + "/oauth2/v2.0/token",
		},
	}

	jwksURLEntra := microsoftLoginBaseURL + oauthConfig.TenantID + "/discovery/v2.0/keys"
	var err error
	oauthConfig.JWKSEntra, err = keyfunc.Get(jwksURLEntra, keyfunc.Options{
		RefreshInterval: time.Hour,
		RefreshErrorHandler: func(err error) {
			log.Printf(constants.ErrJWKSRefreshErrorMsg, err)
		},
	})
	if err != nil {
		return fmt.Errorf(constants.ErrFailedToLoadJWKSMsg, err)
	}

	log.Info().
		Str("client_id", oauthConfig.ClientID).
		Str("tenant_id", oauthConfig.TenantID).
		Str("redirect_uri", oauthConfig.RedirectURI).
		Msg(constants.ErrOAuthConfigInitializedMsg)

	return nil
}

func GenerateRandomState() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)
}

func GenPKCE() (codeVerifier, codeChallenge string) {
	b := make([]byte, 32)
	rand.Read(b)
	codeVerifier = base64.RawURLEncoding.EncodeToString(b)
	h := sha256.Sum256([]byte(codeVerifier))
	codeChallenge = base64.RawURLEncoding.EncodeToString(h[:])
	return
}
