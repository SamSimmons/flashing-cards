package main

import "net/http"

// Route is route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a bunch of routes
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"CardIndex",
		"GET",
		"/cards",
		CardIndex,
	},
	Route{
		"CardShow",
		"GET",
		"/cards/{cardID}",
		CardShow,
	},
	Route{
		"HandleCors",
		"OPTIONS",
		"/cards/create",
		HandleCors,
	},
	Route{
		"CardCreate",
		"POST",
		"/cards/create",
		CardCreate,
	},
	Route{
		"CardDestroy",
		"GET",
		"/cards/{cardID}/delete",
		CardDestroy,
	},
}
