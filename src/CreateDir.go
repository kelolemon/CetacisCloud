package src

import (
	"github.com/kataras/iris"
	"os"
)

func CreateDir(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("IsLog"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	dirname := ctx.Params().Get("dirname")
	err := os.MkdirAll(FilePath + dirname, os.ModePerm)
	if err != nil {
		RtData := flag {
			Success: "0",
		}
		ctx.ContentType("application/json")
		_, _ = ctx.JSON(RtData)
		return
	}
	RtData := flag {
		Success: "1",
	}
	ctx.ContentType("application/json")
	_, _ = ctx.JSON(RtData)
}
