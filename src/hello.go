package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  //オプションを解析します。デフォルトでは解析しません。
    fmt.Println(r.Form)  //このデータはサーバのプリント情報に出力されます。
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!") //ここでwに入るものがクライアントに出力されます。
}

func saySawai(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    fmt.Fprintf(w, "You Are Sawai!!!!?!?!?!?!!??")
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //リクエストを取得するメソッド
    if r.Method == "GET" {
        t, _ := template.ParseFiles("src/templates/login.html")
        t.Execute(w, nil)
    } else {
        //ログインデータがリクエストされ、ログインのロジック判断が実行されます。
        r.ParseForm()
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}

func main() {
    fmt.Println("go server has waked up!")
    http.HandleFunc("/", sayhelloName) //アクセスのルーティングを設定します。
	http.HandleFunc("/sawai", saySawai)
    http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) //監視するポートを設定します。
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}