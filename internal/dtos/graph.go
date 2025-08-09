package dtos

import (
	"yet-another-itsm/internal/utils"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

type UserInfoResponse struct {
	ID             string `json:"id"`
	DisplayName    string `json:"display_name"`
	Surname        string `json:"surname"`
	GivenName      string `json:"given_name"`
	Email          string `json:"email"`
	MobilePhone    string `json:"mobile_phone"`
	JobTitle       string `json:"job_title"`
	OfficeLocation string `json:"office_location"`
	Department     string `json:"department"`
	Manager        string `json:"manager"`
}

func GetManagerId(user models.Userable) string {
	if user.GetManager() != nil {
		return utils.GetStringValue(user.GetManager().GetId())
	}
	return ""
}
