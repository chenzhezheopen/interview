package main

import (
	"Db"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"code.google.com/p/go.net/websocket"
	// "code.google.com/p/go.net/websocket"
	//"zhe\毕业设计应用程序\bishe"
	// "github.com/golang/net/websocket"
	// "golang.org/x/net/websocket"
	// "github.com/golang/net/websocket"
	// "golang.org/x/net/websocket"
)

type Config struct {
	MimeType map[string]string `json:"mimetype"`
}

var Connty = make(map[*websocket.Conn]string)

type jur struct {
	doclogger *log.Logger
	drulogger *log.Logger
	uselogger *log.Logger
}

var mdoctor []*websocket.Conn
var doc Db.Doctor
var cfg *Config
var logger jur
var usname string

// var user Db.User

func init() {
	logfile, _ := os.OpenFile("logger.txt", os.O_RDWR|os.O_APPEND, 0)
	logger.doclogger = log.New(logfile, "\r\n", log.Ldate|log.Lshortfile|log.Ltime) //医生登录日志初始化

	logfiletwo, _ := os.OpenFile("drugs.txt", os.O_RDWR|os.O_APPEND, 0) //药品库存出入日志
	logger.drulogger = log.New(logfiletwo, "\r\n", log.Ldate|log.Lshortfile|log.Ltime)

	path := "./config.toml"
	fp, _ := filepath.Abs(path)
	fmt.Println(fp)
	Db.Peiz(fp)
	// config := new(BlogConfig)
	// toml.DecodeFile(path, config)
	// tol = config
	// fmt.Println(cfg.Content.Css)

	if cfg == nil {
		t := &Config{}
		buf, err := ioutil.ReadFile("config.json")
		if err != nil {
			fmt.Printf("read config.json error:%s\n", err)
			panic(err)
		}
		err = json.Unmarshal(buf, t)
		if err != nil {
			fmt.Printf("config.json file error:%s\n", err)
			panic(err)
		}
		cfg = t
	}
}
func main() {
	http.HandleFunc("/page/", SayHello)         //页面返回
	http.HandleFunc("/page/hospital", Hospital) //判断医生账号密码是否正确
	server := http.Server{
		Addr: ":5656",
	}
	http.Handle("/add", websocket.Handler(add))                 //添加websocket.Conn
	http.Handle("/communicate", websocket.Handler(communicate)) //通话服务
	http.Handle("/user_obt", websocket.Handler(obt))
	http.Handle("/pay", websocket.Handler(pay))   //用户付款服务
	http.Handle("/land", websocket.Handler(land)) //用户咨询服务
	server.ListenAndServe()
}
func Hospital(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	s := make([][]byte, 2)
	doc = Db.Chaxun(req.Form["user"][0], req.Form["pwd"][0])
	// fmt.Println(Db.Chaxun(req.Form["user"][0], req.Form["pwd"][0]))
	// fmt.Println(doc)
	// fmt.Println(doc.Id)
	if doc.Id > 0 { //判断账号密码是否正确

		logger.doclogger.Println(doc.Name + "登录")
		buf, _ := ioutil.ReadFile("page/hoshou.html")
		// sep := []byte(doc.User)
		s[0] = buf
		s[1] = []byte("<script>$(function(){$('#ii').text('" + doc.Name + "');})</script>")
		sep := []byte("")
		w.Write(bytes.Join(s, sep))
		return
	}

	buf, _ := ioutil.ReadFile("page/index.html")
	s[0] = buf
	s[1] = []byte("<script>alert('账号或密码错误')</script>")
	sep := []byte("")
	w.Write(bytes.Join(s, sep))
}
func SayHello(w http.ResponseWriter, req *http.Request) {
	// fmt.Println(req.RequestURI)
	if req.RequestURI == "/page/hoshou.html" { //判断用户是否会直接访问医院后台
		s := make([][]byte, 2)
		buf, _ := ioutil.ReadFile("page/index.html") //如果直接访问则通知他通过账号密码登录并返回主页面
		s[0] = buf
		s[1] = []byte("<script>alert('请通过正常途径登录')</script>")
		sep := []byte("")
		w.Write(bytes.Join(s, sep))
		fmt.Println("恶意访问")
		return
	}
	pos := strings.LastIndexByte(req.RequestURI, '.')
	uri := req.RequestURI

	ext := uri[pos+1:]
	for v, k := range cfg.MimeType {
		if ext == v {
			w.Header().Set("Content-Type", k)
		}
	}
	buf, err := ioutil.ReadFile("." + req.RequestURI)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Write(buf)
}

func obt(w *websocket.Conn) { //获得当前在线医生信息

	fmt.Println("sssss")
	var error error
	for {
		for _, k := range Connty {
			// fmt.Println(k)
			usname = Db.Pass(k)
			if error = websocket.Message.Send(w, usname); error != nil {
				fmt.Println("不能够发送消息 ")
			}
		}
		break
	}
}

func land(w *websocket.Conn) {
	fmt.Println("pppppppp")
	var error error
	for {
		var reply string
		if error = websocket.Message.Receive(w, &reply); error != nil {
			fmt.Println("不能够接受消息 error==", error)
			delete(Connty, w)
			break
		}
		// msg := "我已经收到消息 Received:" + reply
		Connty[w] = reply
		mg := Db.Pass(reply)
		// fmt.Println(mg)
		if error = websocket.Message.Send(w, mg); error != nil {
			fmt.Println("不能够发送消息 ")
			break
		}
	}
}

func pay(w *websocket.Conn) {
	var error error
	for {
		var reply string
		if error = websocket.Message.Receive(w, &reply); error != nil {
			fmt.Println("不能够接受消息 error==", error)
			break
		}
		msg := "我已经收到消息 Received:" + reply
		if error = websocket.Message.Send(w, msg); error != nil {
			fmt.Println("不能够发送消息 悲催哦")
			break
		}
	}
}
func communicate(w *websocket.Conn) {

	var error error
	for {
		var reply string
		if error = websocket.Message.Receive(w, &reply); error != nil {
			fmt.Println("不能够接受消息 error==", error)
			break
		}
		msg := "我已经收到消息 Received:" + reply
		if error = websocket.Message.Send(w, msg); error != nil {
			fmt.Println("不能够发送消息 悲催哦")
			break
		}
	}
}
func add(w *websocket.Conn) {
	fmt.Println("不能够发")
	var error error
	for {
		var reply string
		if error = websocket.Message.Receive(w, &reply); error != nil {
			fmt.Println("不能够接受消息 error==", error)
			break
		}
		fmt.Println(reply)

		msg := "我已经收到消息 Received:" + reply
		if error = websocket.Message.Send(w, msg); error != nil {
			fmt.Println("不能够发送消息 悲催哦55")
			break
		}
	}
}
