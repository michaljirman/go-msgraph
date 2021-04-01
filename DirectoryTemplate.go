package msgraph

import "fmt"

type DirectoryRoleTemplate struct {
	ID          string
	Description string
	DisplayName string

	graphClient *GraphClient // the graphClient that called the group
}

func (r *DirectoryRoleTemplate) setGraphClient(gC *GraphClient) {
	r.graphClient = gC
}

func (r *DirectoryRoleTemplate) String() string {
	return fmt.Sprintf("DirectoryRoleTemplate(DisplayName: \"%v\", ID: \"%v\", Description: \"%v\")",
		r.DisplayName, r.ID, r.Description)
}
