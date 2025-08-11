package service

import (
	"fmt"

	"yet-another-itsm/internal/config"
	"yet-another-itsm/internal/constants"
	"yet-another-itsm/internal/utils"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"

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
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrCouldNotCreateOBOCredential, err)
	}

	scopes := []string{g.config.GraphScope}
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, scopes)
	if err != nil {
		return nil, fmt.Errorf(utils.ErrorFormat, constants.ErrCouldNotCreateGraphClient, err)
	}

	return client, nil
}
