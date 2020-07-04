package src

import (
	"github.com/kataras/iris"
	"io"
	"os"
)

const FilePath = "./file/"

func PostUpload(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("IsLog"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	file, info, err := ctx.FormFile("file")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Application().Logger().Warnf("Error while uploading: %v", err.Error())
		return
	}
	defer file.Close()
	filename := info.Filename
	out, err := os.OpenFile(FilePath + filename, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Application().Logger().Warnf("Error while preparing the new file: %v", err.Error())
		return
	}
	defer out.Close()
	_, _ = io.Copy(out, file)
	RtData := flag {
		Success: "1",
	}
	ctx.ContentType("application/json")
	_, _ = ctx.JSON(RtData)
}