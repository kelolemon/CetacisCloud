package src

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"io/ioutil"
)

func PostLogin(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("IsLog"); auth {
		RtData := flag {
			Success: "-1",
		}
		ctx.ContentType("application/json")
		_, _ = ctx.JSON(RtData)
		return
	}
	var Data account
	bytes, _ := ioutil.ReadFile("./config/account.json")
	_ = json.Unmarshal(bytes, &Data)
	user := ctx.FormValue("user")
	pass := ctx.FormValue("pass")
	if user == Data.User && pass == Data.Pass {
		fmt.Println("Login success")
		session := sess.Start(ctx)
		session.Set("IsLog", true)
		RtData := flag {
			Success: "1",
		}
		ctx.ContentType("application/json")
		_, _ = ctx.JSON(RtData)
		return
	} else {
		RtData := flag {
			Success: "0",
		}
		ctx.ContentType("application/json")
		_, _ = ctx.JSON(RtData)
	}
}
