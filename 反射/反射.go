package main

import (
	"fmt"
	"net/http"
)

//定义index方法
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello world</h1>")
}

func main() {

	http.HandleFunc("/", index)       //通过index访问index方法内的数据
	fmt.Println("strat Server...")    //提示信息可要可不要
	http.ListenAndServe(":3307", nil) //定义端口为：3000 返回空（nil)
}
