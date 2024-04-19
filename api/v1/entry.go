package v1

import "oms-fiber/api/v1/system"

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	//ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
