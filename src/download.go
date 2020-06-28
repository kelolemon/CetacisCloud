package src

import "github.com/kataras/iris"

func PostDownload(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("IsLog"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	filename := ctx.Params().Get("filename")
	FilePath := FilePath + filename
	_ = ctx.SendFile(FilePath, filename)
}