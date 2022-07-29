package routes

import "server/utils"

var LOGIN_FAILED utils.Response = utils.NewResponse(nil, "Incorrect username or password.")
var USER_NOTFOUND utils.Response = utils.NewResponse(nil, "User doesn't exist.")
var UNAUTHORIZED utils.Response = utils.NewResponse(nil, "You are not authorized to perform this action.")
var CLASSROOM_NOTFOUND utils.Response = utils.NewResponse(nil, "Classroom doesn't exist.")
var POST_NOTFOUND utils.Response = utils.NewResponse(nil, "Post doesn't exist.")
