package controller

import (
	"github.com/riete/aliyun-monitor-operator/pkg/controller/slbmonitor"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, slbmonitor.Add)
}
