package src

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"io/ioutil"
)

func GetFileList(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("IsLog"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	m := make(map[string]interface{})
	bytes, _ := ioutil.ReadFile(ShareMapPath)
	_ = json.Unmarshal(bytes, &m)
	files, _ := ioutil.ReadDir(FilePath)
	DataList := FileList {
		List: make([]FileInfo, 0),
	}
	for _, f := range files {
		_, ok := m[f.Name()]
		DataFile := FileInfo {
			FileName: f.Name(),
			FileSize: f.Size(),
			FileType: f.IsDir(),
			IsShare: ok,
		}
		DataList.List = append(DataList.List, DataFile)
		fmt.Println(f.Name(), f.Size(), f.IsDir())
	}
	_, _ = ctx.JSON(DataList)
}
