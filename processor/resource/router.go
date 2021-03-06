package resource

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"StatisticCreate",
		"POST",
		"/statistics",
		CreateNewStatistic,
	},
	Route{
		"StatisticsIndex",
		"GET",
		"/",
		GetStatistics,
	},
	Route{
		"IPSend",
		"POST",
		"/usage",
		GetIp,
	},
	Route{
		"SendMail",
		"Get",
		"/sendmail",
		SendStatisticsForced,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
