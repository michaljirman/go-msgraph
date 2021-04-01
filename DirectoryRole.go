package msgraph

import (
	"fmt"
)

type DirectoryRole struct {
	ID             string
	Description    string
	DisplayName    string
	RoleTemplateID string

	graphClient *GraphClient // the graphClient that called the group
}

func (r *DirectoryRole) setGraphClient(gC *GraphClient) {
	r.graphClient = gC
}

func (r *DirectoryRole) String() string {
	return fmt.Sprintf("DirectoryRole(DisplayName: \"%v\", ID: \"%v\", RoleTemplateID: \"%v\", Description: \"%v\")",
		r.DisplayName, r.ID, r.RoleTemplateID, r.Description)
}
