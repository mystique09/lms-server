package routes

import (
	"server/utils"
)

// auth errors
var LOGIN_FAILED utils.Response = utils.NewResponse(nil, "Incorrect username or password.")
var USER_NOTFOUND utils.Response = utils.NewResponse(nil, "User doesn't exist.")
var MISSING_FIELDS utils.Response = utils.NewResponse(nil, "Missing required fields.")
var USER_ALREADY_EXIST utils.Response = utils.NewResponse(nil, "User already exist.")

// role validation errors
var UNAUTHORIZED utils.Response = utils.NewResponse(nil, "You are not authorized to perform this action.")
var NOT_A_MEMBER utils.Response = utils.NewResponse(nil, "You are not a member of this classroom.")

// classroom errors
var CLASSROOM_NOTFOUND utils.Response = utils.NewResponse(nil, "Classroom doesn't exist.")

// classworks errors
var CLASSWORK_NOTFOUND utils.Response = utils.NewResponse(nil, "Classwork doesn't exist.")

// post errors
var POST_NOTFOUND utils.Response = utils.NewResponse(nil, "Post doesn't exist.")

// field errors
var UNKNOWN_FIELD utils.Response = utils.NewResponse(nil, "Unknown field in query params.")
var MISSING_ID_FIELD utils.Response = utils.NewResponse(nil, "Missing id field in path parameter.")
var EMPTY_QUERY_PARAM utils.Response = utils.NewResponse(nil, "Empty query parameter.")
var NEGATIVE_OFFSET utils.Response = utils.NewResponse(nil, "Offset must not be negative.")
