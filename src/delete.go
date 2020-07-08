package src

import (
	"github.com/kataras/iris"
	"os"
)

func GetDelete(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("IsLog"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	filename := ctx.Params().Get("filename")
	err := os.Remove(FilePath + filename)
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
