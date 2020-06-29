package src

import "github.com/kataras/iris"

func IsLog(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("IsLog"); !auth {
		RtData := flag {
			Success: "0",
		}
		ctx.ContentType("application/json")
		_, _ = ctx.JSON(RtData)
		return
	} else {
		RtData := flag {
			Success: "1",
		}
		ctx.ContentType("application/json")
		_, _ = ctx.JSON(RtData)
	}
}
