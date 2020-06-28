package src

import "github.com/kataras/iris"

func PostLogout(ctx iris.Context) {
	session := sess.Start(ctx)
	session.Set("IsLog", false)
	RtData := flag {
		Success: "1",
	}
	ctx.ContentType("application/json")
	_, _ = ctx.JSON(RtData)
}
