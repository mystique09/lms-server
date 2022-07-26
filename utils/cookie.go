package utils

import "net/http"

func NewCookie(name, value string, mxAge int) http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.MaxAge = mxAge
	cookie.SameSite = http.SameSiteStrictMode
	cookie.Path = "/"
	cookie.HttpOnly = true
	return *cookie
}
