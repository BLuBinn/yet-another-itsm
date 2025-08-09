package middleware

import (
	"fmt"
	"strings"
	"time"

	"yet-another-itsm/internal/config"
	"yet-another-itsm/internal/constants"
	"yet-another-itsm/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
)

type AzureADClaims struct {
	Audience interface{} `json:"aud"`
	Issuer   string      `json:"iss"`
	Subject  string      `json:"sub"`
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	OID      string      `json:"oid"`
	TID      string      `json:"tid"`
	UPN      string      `json:"upn"`
	Roles    []string    `json:"roles"`
	Scp      string      `json:"scp"`
	AppID    string      `json:"appid"`
	IPAddr   string      `json:"ipaddr"`
	Expiry   int64       `json:"exp"`
	IssuedAt int64       `json:"iat"`
	jwt.RegisteredClaims
}

func validateAuthHeader(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		log.Warn().Str("ip", c.ClientIP()).Msg(constants.ErrMissingAuthHeaderMsg)
		return "", fmt.Errorf(constants.ErrMissingAuthHeaderMsg)
	}
	return strings.TrimPrefix(authHeader, "Bearer "), nil
}

func validateAndExtractClaims(rawToken string, keyfunc jwt.Keyfunc, c *gin.Context) (*AzureADClaims, error) {
	token, err := jwt.ParseWithClaims(rawToken, &AzureADClaims{}, keyfunc)
	if err != nil {
		log.Warn().Err(err).Str("ip", c.ClientIP()).Msg(constants.ErrFailedToParseJWTTokenMsg)
		return nil, fmt.Errorf(constants.ErrFailedToParseJWTTokenMsg)
	}

	if !token.Valid {
		log.Warn().Str("ip", c.ClientIP()).Msg(constants.ErrInvalidTokenMsg)
		return nil, fmt.Errorf(constants.ErrInvalidTokenMsg)
	}

	claims, ok := token.Claims.(*AzureADClaims)
	if !ok {
		log.Warn().Str("ip", c.ClientIP()).Msg(constants.ErrInvalidTokenClaimsMsg)
		return nil, fmt.Errorf(constants.ErrInvalidTokenClaimsMsg)
	}

	return claims, nil
}

func validateTokenMetadata(claims *AzureADClaims, oauthConfig *config.OAuthConfig, c *gin.Context) error {
	if err := validateTokenExpiry(claims); err != nil {
		log.Warn().Err(err).Str("ip", c.ClientIP()).Msg(constants.ErrTokenExpiredMsg)
		return fmt.Errorf(constants.ErrTokenExpiredMsg)
	}

	expectedIssuer := fmt.Sprintf("https://sts.windows.net/%s/", oauthConfig.TenantID)
	if claims.Issuer != expectedIssuer {
		log.Warn().Str("issuer", claims.Issuer).Str("expected_v1", expectedIssuer).Msg(constants.ErrInvalidTokenIssuerMsg)
		return fmt.Errorf(constants.ErrInvalidTokenIssuerMsg)
	}

	if !validateAudience(claims.Audience, oauthConfig.ClientID) {
		log.Warn().Interface("audience", claims.Audience).Str("expected", oauthConfig.ClientID).Msg(constants.ErrInvalidTokenAudienceMsg)
		return fmt.Errorf(constants.ErrInvalidTokenAudienceMsg)
	}

	if claims.TID != oauthConfig.TenantID {
		log.Warn().Str("tid", claims.TID).Str("expected", oauthConfig.TenantID).Msg(constants.ErrInvalidTenantIDMsg)
		return fmt.Errorf(constants.ErrInvalidTenantIDMsg)
	}

	return nil
}

func AuthMiddleWare(oauthConfig *config.OAuthConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		if oauthConfig == nil || oauthConfig.JWKSEntra == nil {
			log.Error().Msg("OAuth not initialized")
			utils.SendInternalServerError(c, constants.ErrAuthServiceNotAvailableMsg)
			c.Abort()
			return
		}

		rawToken, err := validateAuthHeader(c)
		if err != nil {
			utils.SendUnauthorized(c, err.Error())
			c.Abort()
			return
		}

		claims, err := validateAndExtractClaims(rawToken, oauthConfig.JWKSEntra.Keyfunc, c)
		if err != nil {
			utils.SendUnauthorized(c, err.Error())
			c.Abort()
			return
		}

		if err := validateTokenMetadata(claims, oauthConfig, c); err != nil {
			utils.SendUnauthorized(c, err.Error())
			c.Abort()
			return
		}

		ctx := utils.SetTenantContext(
			c.Request.Context(),
			claims.TID,
			claims.OID,
			claims.Name,
			rawToken,
		)
		c.Request = c.Request.WithContext(ctx)

		log.Info().
			Str("user_id", claims.OID).
			Str("user_name", claims.Name).
			Msg(constants.ErrUserAuthenticatedSuccessfullyMsg)

		c.Next()
	}
}

func validateAudience(audience interface{}, expectedAudience string) bool {
	switch aud := audience.(type) {
	case string:
		return aud == expectedAudience
	case []interface{}:
		for _, a := range aud {
			if audStr, ok := a.(string); ok && audStr == expectedAudience {
				return true
			}
		}
		return false
	case []string:
		for _, a := range aud {
			if a == expectedAudience {
				return true
			}
		}
		return false
	default:
		return false
	}
}

func validateTokenExpiry(claims *AzureADClaims) error {
	if claims.Expiry == 0 {
		return fmt.Errorf(constants.ErrTokenExpiryNotSetMsg)
	}

	expiryTime := time.Unix(claims.Expiry, 0)
	currentTime := time.Now()

	if currentTime.After(expiryTime) {
		return fmt.Errorf(constants.ErrTokenExpiredAtMsg, expiryTime, currentTime)
	}

	return nil
}
