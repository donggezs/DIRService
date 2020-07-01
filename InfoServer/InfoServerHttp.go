package main

import (
	"DIRService/Handle"
	"DIRService/Model"
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

func GoDoIt(path string) ([]byte, error) {
	req, err := Handle.HttpGet("http://" + Handle.SerAddr + "/?path=" + path)
	//req, err := Handle.HttpGet(Handle.SerAddr + "/?path=" + path)
	if err != nil {
		return nil, err
	}
	res := Model.Res{}
	json.Unmarshal(req, &res)

	info := new(Model.Info)
	info.Path = path
	for _, item := range res.Dirs {
		if item.IsDir == true {
			info.DirCount += 1
		} else {
			info.FileCount += 1
			info.TotalSize += item.Size
		}
	}
	result, _ := json.Marshal(info)
	return result, nil
}

func InfodexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.RequestURI() == "/favicon.ico" { //屏蔽
		return
	}
	path := r.FormValue("path")

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res,err:= GoDoIt(path)
			if err != nil {
				log.Panic(err)
				return
			}
			w.Header().Set("content-type", "text/json")
			w.Write(res)
		}()
	}
	wg.Wait()

	/*res,err:= GoDoIt(path)
	if err != nil {
		fmt.Fprintf(w, "error %v\n", err)
	}
	w.Header().Set("content-type", "text/json")
	w.Write(res)*/


	/*req, err := Handle.HttpGet("http://" + Handle.SerAddr + "/?path=" + path)
	//req, err := Handle.HttpGet(Handle.SerAddr + "/?path=" + path)
	if err != nil {
		fmt.Fprintf(w, "请求目录信息出错 %v\n", err)
	}
	res := Model.Res{}
	json.Unmarshal(req, &res)

	info := new(Model.Info)
	info.Path = path
	for _, item := range res.Dirs {
		if item.IsDir == true {
			info.DirCount += 1
		} else {
			info.FileCount += 1
			info.TotalSize += item.Size
		}
	}

	result, _ := json.Marshal(info)
	w.Header().Set("content-type", "text/json")
	w.Write(result)*/
}

func main() {
	http.HandleFunc("/", InfodexHandler)
	http.ListenAndServe(Handle.InfoAddr, nil) //开始监听
}
