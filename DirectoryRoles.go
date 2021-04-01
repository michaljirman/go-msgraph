package msgraph

import (
	"fmt"
	"strings"
)

type DirectoryRoles []DirectoryRole

func (r DirectoryRoles) setGraphClient(gC *GraphClient) DirectoryRoles {
	for i := range r {
		r[i].setGraphClient(gC)
	}
	return r
}

func (r DirectoryRoles) String() string {
	var dirRoles = make([]string, len(r))
	for i, role := range r {
		dirRoles[i] = role.String()
	}
	return fmt.Sprintf("DirectoryRoles(%v)", strings.Join(dirRoles, ", "))
}
