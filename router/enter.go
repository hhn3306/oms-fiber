package router

import "oms-fiber/router/system"

type AllGroupRouter struct {
	System system.GroupRouter
}

var AllGroupRouterApp = new(AllGroupRouter)
