package msgraph

import (
	"fmt"
	"strings"
)

type DirectoryRoleTemplates []DirectoryRoleTemplate

func (r DirectoryRoleTemplates) setGraphClient(gC *GraphClient) DirectoryRoleTemplates {
	for i := range r {
		r[i].setGraphClient(gC)
	}
	return r
}

func (r DirectoryRoleTemplates) String() string {
	var dirRoles = make([]string, len(r))
	for i, role := range r {
		dirRoles[i] = role.String()
	}
	return fmt.Sprintf("DirectoryRoleTemplates(%v)", strings.Join(dirRoles, ", "))
}
