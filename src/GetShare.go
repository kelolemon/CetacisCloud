package src

import (
	"encoding/base64"
	"encoding/json"
	"github.com/kataras/iris"
	"io/ioutil"
	"os"
)

const ShareMapPath = "./config/share_map.data"

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func GetShareLink(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("IsLog"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	filename := ctx.Params().Get("filename")
	m := make(map[string]interface{})
	if !IsExist(ShareMapPath) {
		data, _ := json.Marshal(m)
		_ = ioutil.WriteFile(ShareMapPath, data, 0666)
	} else {
		bytes, _ := ioutil.ReadFile(ShareMapPath)
		_ = json.Unmarshal(bytes, &m)
	}
	m[filename] = true
	data, _ := json.Marshal(m)
	_ = ioutil.WriteFile(ShareMapPath, data, 0666)
	ShareName := base64.StdEncoding.EncodeToString([]byte(filename))
	ShareLink := "/api/share/file/" + ShareName
	RtData := LinkInfo {
		Link: ShareLink,
	}
	_, _ = ctx.JSON(RtData)
}

func DeleteShareLink(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("IsLog"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	filename := ctx.Params().Get("filename")
	bytes, _ := ioutil.ReadFile(ShareMapPath)
	m := make(map[string]interface{})
	_ = json.Unmarshal(bytes, &m)
	if _, ok := m[filename]; !ok {
		RtData := flag {
			Success: "0",
		}
		_, _ = ctx.JSON(RtData)
		return
	}
	delete(m, filename)
	data, _ := json.Marshal(m)
	_ = ioutil.WriteFile(ShareMapPath, data, 0666)
	RtData := flag {
		Success: "1",
	}
	_, _ = ctx.JSON(RtData)
}

func GetShareFile(ctx iris.Context) {
	filename := ctx.Params().Get("filename")
	bytes, _ := base64.StdEncoding.DecodeString(filename)
	filename = string(bytes)
	bytes, _ = ioutil.ReadFile(ShareMapPath)
	m := make(map[string]interface{})
	_ = json.Unmarshal(bytes, &m)
	if _, ok := m[filename]; ok {
		FilePath := FilePath + filename
		_ = ctx.SendFile(FilePath, filename)
	} else {
		RtData := flag {
			Success: "0",
		}
		_, _ = ctx.JSON(RtData)
	}
}
