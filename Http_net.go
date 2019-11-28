package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

//type Data struct {
//	Name     string
//	Password int
//}

type Ret struct {
	Code  int
	Param string
	Msg   string
}

func main() {
	http.HandleFunc("/", HelloServer)

	// 设置监听的端口
	err := http.ListenAndServe(":9092", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	password := r.FormValue("password")
	fmt.Println("name",name)
	fmt.Println("password",password)
	ret := new(Ret)
	if name != ""  {
		ret.Code = 0
		ret.Msg = "success"
		ret_json,_ := json.Marshal(ret)
		io.WriteString(w, string(ret_json))
	}else {
		ret.Code = 1
		ret.Msg = "error"
		ret_json,_ := json.Marshal(ret)
		io.WriteString(w, string(ret_json))
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	// 解析参数, 默认是不会解析的
	r.ParseForm()
	// 这些信息是输出到服务器端的打印信息
	fmt.Println("request map:", r.Form)
	fmt.Println("request map:", r.Method)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ";"))
	}
	// 这个写入到w的信息是输出到客户端的
	fmt.Fprintf(w, "Hello gerryyang!\n")
}
