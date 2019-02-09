package controller

import (
	"github.com/sitewhere/sitewhere-operator/pkg/controller/sitewhere"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, sitewhere.Add)
}
