package main

import (
	"DIRService/Handle"
	"DIRService/Model"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

var (
	root   = ""           //设置根路径
	ostype = runtime.GOOS //获取系统类型，针对不同系统写路径时使用“\\”或者“/”
)
var strRet string

func init() {
	if ostype == "windows" {
		strRet = "\\"
	} else {
		strRet = "/"
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.RequestURI() == "/favicon.ico" { //屏蔽
		return
	}

	r.ParseForm()
	res := new(Model.Res)
	res.Path = r.Form.Get("path")
	//组合路径
	path := root + strRet + res.Path
	//循环
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		//fmt.Println(path, ":", info.Name(), ",isDir", info.IsDir(), "size", info.Size())
		dir := Model.Dir{Name: info.Name(), IsDir: info.IsDir(), Size: info.Size()}
		res.Dirs = append(res.Dirs, dir)
		return err
	})
	if err != nil {
		fmt.Fprintf(w, "filepath.Walk() 出错 %v\n", err)
	}
	result, _ := json.Marshal(res)

	//io.WriteString(w, string(result))
	//json.NewEncoder(w).Encode(res)

	w.Header().Set("content-type", "text/json")
	w.Write(result)
}

func main() {
	list := os.Args //监听输入参数
	if list != nil && len(list) > 1 {
		root = list[1]
	}
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(Handle.SerAddr, nil) //开始监听
}
