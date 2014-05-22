package main

import (
	"fmt"
	"m"
	"net/http"
)

func router(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var (
		token     string = "DataE4st"
		signature string = r.Form.Get("signature")
		timestamp string = r.Form.Get("timestamp")
		nonce     string = r.Form.Get("nonce")
		echostr   string = r.Form.Get("echostr")
	)
	if echostr == "" {
		if m.CheckSignature(token, signature, timestamp, nonce) { //验证成功,开始处理消息
			restr, e := m.ProMsg(r.Body)
			if e == nil {
				fmt.Fprintf(w, restr) //返回给微信服务器
			} else {
				fmt.Fprint(w, e.Error()) //返回出错信息
			}
		} else {
			fmt.Fprintf(w, "404 page not found")
		}
	} else {
		fmt.Fprintf(w, echostr)
	}
}
func main() {
	http.HandleFunc("/WeChat", router) //设置访问路由
	http.ListenAndServe(":80", nil)    //设置监听端口
}

// http://localhost/WeChat?signature=652abf78a6077402ff735a5ad7666d91d258fb19
