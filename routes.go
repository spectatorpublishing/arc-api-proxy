package main

import (
	"arc-api-proxy/controllers"
	"net/http"
)

type Route struct {
	Name        string
	Methods     []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// Namespace Pattern with a prefix path. For example r.Namespace("/v1")
func (r Route) Namespace(namespace string) string {
	// Prefix the namespace to the pattern
	return namespace + r.Pattern
}

var ProxyRoutes = Routes{

	Route{
		"ProxyArcApi",
		[]string{"GET"},
		"/{rest:.*}",
		controllers.ProxyArcApi,
	},
}
