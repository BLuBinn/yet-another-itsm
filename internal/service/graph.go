package service

import (
	"context"
	"fmt"

	"yet-another-itsm/internal/config"
	"yet-another-itsm/internal/constants"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	msgraphsdkUser "github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/rs/zerolog/log"
)

type GraphService struct {
	config *config.OAuthConfig
}

func NewGraphService(cfg *config.OAuthConfig) *GraphService {
	return &GraphService{config: cfg}
}

// GetGraphTokenOnBehalfOf performs the OAuth 2.0 OBO flow.
func (g *GraphService) GetGraphTokenOnBehalfOf(tokenStr string) (*msgraphsdk.GraphServiceClient, error) {
	log.Info().
		Str("service", "GraphService").
		Str("endpoint", "GetGraphTokenOnBehalfOf").
		Msg("Getting graph token on behalf of")

	cred, err := azidentity.NewOnBehalfOfCredentialWithSecret(
		g.config.TenantID,
		g.config.ClientID,
		tokenStr,
		g.config.ClientSecret,
		&azidentity.OnBehalfOfCredentialOptions{},
	)
	if err != nil {
		return nil, fmt.Errorf(constants.ErrorFormat, constants.ErrCouldNotCreateOBOCredential, err)
	}

	scopes := []string{g.config.GraphScope}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf(constants.ErrorFormat, constants.ErrCouldNotCreateGraphClient, err)
	}

	return client, nil
}

// GetCurrentUserFromGraph gets the current user from Microsoft Graph.
func (g *GraphService) GetCurrentUserFromGraph(tokenStr string) (models.Userable, error) {
	log.Info().
		Str("service", "GraphService").
		Str("endpoint", "GetCurrentUserFromGraph").
		Msg("Getting current user from graph")
	client, err := g.GetGraphTokenOnBehalfOf(tokenStr)
	if err != nil {
		return nil, fmt.Errorf(constants.ErrorFormat, constants.ErrCouldNotCreateGraphClient, err)
	}

	user, err := client.Me().Get(context.Background(), &msgraphsdkUser.UserItemRequestBuilderGetRequestConfiguration{
		QueryParameters: &msgraphsdkUser.UserItemRequestBuilderGetQueryParameters{
			Select: []string{"id", "displayName", "surname", "givenName", "mail", "mobilePhone", "jobTitle", "officeLocation", "department", "manager"},
		},
	})
	if err != nil {
		return nil, fmt.Errorf(constants.ErrorFormat, constants.ErrFailedToGetCurrentUser, err)
	}

	return user, nil
}
