package src

import "github.com/kataras/iris"

func IsLog(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("IsLog"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	_, _ = ctx.WriteString("The cake is a lie!")
}
