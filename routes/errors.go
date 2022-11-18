package routes

// auth errors
var LOGIN_FAILED Response[any] = newResponse[any](nil, "Incorrect username or password.")
var USER_NOTFOUND Response[any] = newResponse[any](nil, "User doesn't exist.")
var MISSING_FIELDS Response[any] = newResponse[any](nil, "Missing required fields.")
var USER_ALREADY_EXIST Response[any] = newResponse[any](nil, "User already exist.")

// role validation errors
var UNAUTHORIZED Response[any] = newResponse[any](nil, "You are not authorized to perform this action.")
var NOT_A_MEMBER Response[any] = newResponse[any](nil, "You are not a member of this classroom.")

// classroom errors
var CLASSROOM_NOTFOUND Response[any] = newResponse[any](nil, "Classroom doesn't exist.")

// classworks errors
var CLASSWORK_NOTFOUND Response[any] = newResponse[any](nil, "Classwork doesn't exist.")

// post errors
var POST_NOTFOUND Response[any] = newResponse[any](nil, "Post doesn't exist.")

// field errors
var UNKNOWN_FIELD Response[any] = newResponse[any](nil, "Unknown field in query params.")
var MISSING_ID_FIELD Response[any] = newResponse[any](nil, "Missing id field in path parameter.")
var EMPTY_QUERY_PARAM Response[any] = newResponse[any](nil, "Empty query parameter.")
var NEGATIVE_OFFSET Response[any] = newResponse[any](nil, "Offset must not be negative.")
