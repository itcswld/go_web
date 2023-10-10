//record request
/*
階段管理：登入、購物車、遊戲得分、srv
實現個性化： 偏好、 主題或其他設定
追蹤： 記錄和分析使用者行為
*/
package main

import (
	"fmt"
	"net/http"
)

func cookieHandler(w http.ResponseWriter, r *http.Request) {
	//獲取Cookie
	c, err := r.Cookie("req_cookie")
	if err != nil {
		fmt.Printf("cookie: %#v, err: %v \n", c, err)
	}
	//返回請求中名為“req_cookie”的cookie
	cookie := &http.Cookie{
		Name:   "req_cookie",
		Value:  "123456",
		MaxAge: 3600,        //3600sec = 6min， MaxAge < 0  || MaxAge=0 :  delete cookie
		Domain: "localhost", //包含子網域 ； 沒設Domain	則不含子網域
		Path:   "/",         //作用域 “/”含子網域
	}
	//回應前設定Cookie
	http.SetCookie(w, cookie)
	w.Write([]byte("hello"))

}

func main() {
	http.HandleFunc("/setCookie", cookieHandler)
	http.ListenAndServe(":8080", nil)
}
