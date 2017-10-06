// +build codegen

package api

import "strings"

// ExportableName a name which is exportable as a value or name in Go code
func (a *API) ExportableName(name string) string {
	if name == "" {
		return name
	}

	return strings.ToUpper(name[0:1]) + name[1:]
}
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
