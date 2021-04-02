package msgraph

import (
	"encoding/json"
	"fmt"
)

type DirectoryRoleTemplate struct {
	ID          string
	Description string
	DisplayName string

	graphClient *GraphClient // the graphClient that called the group
}

func (t *DirectoryRoleTemplate) setGraphClient(gC *GraphClient) {
	t.graphClient = gC
}

func (t *DirectoryRoleTemplate) String() string {
	return fmt.Sprintf("DirectoryRoleTemplate(DisplayName: \"%v\", ID: \"%v\", Description: \"%v\")",
		t.DisplayName, t.ID, t.Description)
}

// Activate a directory role template. To read a directory role or update its members, it must first be activated in the tenant.
// https://docs.microsoft.com/en-us/graph/api/directoryrole-post-directoryroles?view=graph-rest-1.0&tabs=http
func (t *DirectoryRoleTemplate) Activate() (DirectoryRole, error) {
	if t.graphClient == nil {
		return DirectoryRole{}, ErrNotGraphClientSourced
	}

	resource := fmt.Sprintf("/directoryRoles")
	var marsh struct {
		DirectoryRole DirectoryRole `json:"value"`
	}

	requestData := map[string]string{
		"roleTemplateId": t.ID,
	}
	requestDataJson, err := json.Marshal(requestData)
	if err != nil {
		return DirectoryRole{}, fmt.Errorf("failed to marshal request data: %w", err)
	}

	err = t.graphClient.makePOSTAPICall(resource, nil, requestDataJson, &marsh)
	marsh.DirectoryRole.setGraphClient(t.graphClient)
	return marsh.DirectoryRole, err
}
