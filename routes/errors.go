package routes

// auth errors
var LOGIN_FAILED Response = newResponse(nil, "Incorrect username or password.")
var USER_NOTFOUND Response = newResponse(nil, "User doesn't exist.")
var MISSING_FIELDS Response = newResponse(nil, "Missing required fields.")
var USER_ALREADY_EXIST Response = newResponse(nil, "User already exist.")

// role validation errors
var UNAUTHORIZED Response = newResponse(nil, "You are not authorized to perform this action.")
var NOT_A_MEMBER Response = newResponse(nil, "You are not a member of this classroom.")

// classroom errors
var CLASSROOM_NOTFOUND Response = newResponse(nil, "Classroom doesn't exist.")

// classworks errors
var CLASSWORK_NOTFOUND Response = newResponse(nil, "Classwork doesn't exist.")

// post errors
var POST_NOTFOUND Response = newResponse(nil, "Post doesn't exist.")

// field errors
var UNKNOWN_FIELD Response = newResponse(nil, "Unknown field in query params.")
var MISSING_ID_FIELD Response = newResponse(nil, "Missing id field in path parameter.")
var EMPTY_QUERY_PARAM Response = newResponse(nil, "Empty query parameter.")
var NEGATIVE_OFFSET Response = newResponse(nil, "Offset must not be negative.")
