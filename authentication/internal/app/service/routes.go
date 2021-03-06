package service

import (
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
		"Login",
		"POST",
		"/login",
		HandleLogin,
	},
	Route{
		"Registration",
		"POST",
		"/register",
		HandleRegistration,
	},
	// Route{
	// 	"ValidateToken",
	// 	"POST",
	// 	"/validateToken",
	// 	HandleValidateToken,
	// },
	Route{
		"AddBatch",
		"POST",
		"/addbatch",
		isAuthorized(isAccountType(HandleAddBatch, "teacher")),
	},
	Route{
		"GetBatches",
		"GET",
		"/getbatch",
		isAuthorized(HandleGetBatches),
	},
	Route{
		"AddAssignment",
		"POST",
		"/addassignment/{bID}",
		isAuthorized(isAccountType(HandleAddAssignment, "teacher")),
	},
	Route{
		"GetAssignment",
		"GET",
		"/getassignment/{aID}",
		isAuthorized(HandleGetAssignment),
	},
	Route{
		"Submission",
		"POST",
		"/submit/{aID}",
		isAuthorized(isAccountType(HandleSubmission, "student")),
	},
	//	Route{
	//		"GetAssignments",
	//		"POST",
	//		"/getassignments/{bID}",
	//		isAuthorized(HandleGetAssignments),
	//	},
}
