package utils

import (
	"context"
	"fmt"

	"yet-another-itsm/internal/constants"
)

type ContextKey string

const (
	TenantIDKey    ContextKey = "tenant_id"
	UserIDKey      ContextKey = "user_id"
	UserNameKey    ContextKey = "user_name"
	AccessTokenKey ContextKey = "access_token"
)

func GetTenantID(ctx context.Context) (string, error) {
	tenantID, ok := ctx.Value(TenantIDKey).(string)
	if !ok || tenantID == "" {
		return "", fmt.Errorf(constants.ErrTenantIDNotFound)
	}
	return tenantID, nil
}

func GetUserID(ctx context.Context) (string, error) {
	userID, ok := ctx.Value(UserIDKey).(string)
	if !ok || userID == "" {
		return "", fmt.Errorf(constants.ErrUserIDNotFound)
	}
	return userID, nil
}

func GetUserName(ctx context.Context) (string, error) {
	userName, ok := ctx.Value(UserNameKey).(string)
	if !ok || userName == "" {
		return "", fmt.Errorf(constants.ErrUserNameNotFound)
	}
	return userName, nil
}

func GetAccessToken(ctx context.Context) (string, error) {
	accessToken, ok := ctx.Value(AccessTokenKey).(string)
	if !ok || accessToken == "" {
		return "", fmt.Errorf(constants.ErrAccessTokenNotFound)
	}
	return accessToken, nil
}

func SetTenantContext(ctx context.Context, tenantID, userID, userName, accessToken string) context.Context {
	ctx = context.WithValue(ctx, TenantIDKey, tenantID)
	ctx = context.WithValue(ctx, UserIDKey, userID)
	ctx = context.WithValue(ctx, UserNameKey, userName)
	ctx = context.WithValue(ctx, AccessTokenKey, accessToken)
	return ctx
}
