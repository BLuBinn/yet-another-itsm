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

func AuthMiddleWare(oauthConfig *config.OAuthConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		if oauthConfig == nil || oauthConfig.JWKSEntra == nil {
			log.Error().Msg("OAuth not initialized")
			utils.SendInternalServerError(c, constants.ErrAuthServiceNotAvailableMsg)
			c.Abort()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			log.Warn().Str("ip", c.ClientIP()).Msg(constants.ErrMissingAuthHeaderMsg)
			utils.SendUnauthorized(c, constants.ErrMissingAuthHeaderMsg)
			c.Abort()
			return
		}

		rawToken := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(rawToken, &AzureADClaims{}, oauthConfig.JWKSEntra.Keyfunc)

		if err != nil {
			log.Warn().Err(err).Str("ip", c.ClientIP()).Msg(constants.ErrFailedToParseJWTTokenMsg)
			utils.SendUnauthorized(c, constants.ErrFailedToParseJWTTokenMsg)
			c.Abort()
			return
		}

		if !token.Valid {
			log.Warn().Str("ip", c.ClientIP()).Msg(constants.ErrInvalidTokenMsg)
			utils.SendUnauthorized(c, constants.ErrInvalidTokenMsg)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*AzureADClaims)
		if !ok {
			log.Warn().Str("ip", c.ClientIP()).Msg(constants.ErrInvalidTokenClaimsMsg)
			utils.SendUnauthorized(c, constants.ErrInvalidTokenClaimsMsg)
			c.Abort()
			return
		}

		if err := validateTokenExpiry(claims); err != nil {
			log.Warn().Err(err).Str("ip", c.ClientIP()).Msg(constants.ErrTokenExpiredMsg)
			utils.SendUnauthorized(c, constants.ErrTokenExpiredMsg)
			c.Abort()
			return
		}

		expectedIssuer := fmt.Sprintf("https://sts.windows.net/%s/", oauthConfig.TenantID)

		if claims.Issuer != expectedIssuer {
			log.Warn().Str("issuer", claims.Issuer).Str("expected_v1", expectedIssuer).Msg(constants.ErrInvalidTokenIssuerMsg)
			utils.SendUnauthorized(c, constants.ErrInvalidTokenIssuerMsg)
			c.Abort()
			return
		}

		expectedAudience := oauthConfig.ClientID

		if !validateAudience(claims.Audience, expectedAudience) {
			log.Warn().Interface("audience", claims.Audience).Str("expected", expectedAudience).Msg(constants.ErrInvalidTokenAudienceMsg)
			utils.SendUnauthorized(c, constants.ErrInvalidTokenAudienceMsg)
			c.Abort()
			return
		}

		if claims.TID != oauthConfig.TenantID {
			log.Warn().Str("tid", claims.TID).Str("expected", oauthConfig.TenantID).Msg(constants.ErrInvalidTenantIDMsg)
			utils.SendUnauthorized(c, constants.ErrInvalidTenantIDMsg)
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
