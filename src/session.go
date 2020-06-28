package src

import "github.com/kataras/iris/sessions"

var (
	cookieNameForSessionID = "LogToken"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)
