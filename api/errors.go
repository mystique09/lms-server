package api

// auth errors
var LOGIN_FAILED = newError("Incorrect username or password.")
var USER_NOTFOUND = newError("User doesn't exist.")
var MISSING_FIELDS = newError("Missing required fields.")
var USER_ALREADY_EXIST = newError("User already exist.")

// role validation errors
var UNAUTHORIZED = newError("You are not authorized to perform this action.")
var NOT_A_MEMBER = newError("You are not a member of this classroom.")

// classroom errors
var CLASSROOM_NOTFOUND = newError("Classroom doesn't exist.")

// classworks errors
var CLASSWORK_NOTFOUND = newError("Classwork doesn't exist.")

// post errors
var POST_NOTFOUND = newError("Post doesn't exist.")

// field errors
var UNKNOWN_FIELD = newError("Unknown field in query params.")
var MISSING_ID_FIELD = newError("Missing id field in path parameter.")
var EMPTY_QUERY_PARAM = newError("Empty query parameter.")
var NEGATIVE_OFFSET = newError("Offset must not be negative.")
