package main

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"strings"
	"time"
)

func init() {
	// 获取日志文件句柄
	// 已 只写入文件|没有时创建|文件尾部追加 的形式打开这个文件
	logFile, err := os.OpenFile(`./ws_touch.log`, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	// 设置存储位置
	log.SetOutput(logFile)
}

type websocketClientManager struct {
	conn        *websocket.Conn
	protocol    *string
	addr        *string
	path        string
	sendMsgChan chan string
	recvMsgChan chan string
	isAlive     bool
	timeout     int
}

type ConnectionInfo struct {
	RemoteAddress string `json:"remoteAddress,omitempty"`
	Interval      int    `json:"interval,omitempty"`
	Mode          string `json:"mode,omitempty"`
	Num           int    `json:"num,omitempty"`
}

// NewWsClientManager 构造函数
func NewWsClientManager(protocol, addr, path string, timeout int) *websocketClientManager {
	if !strings.Contains("wss", strings.ToLower(protocol)) {
		log.Fatal("incorrect protocol type")
		os.Exit(1)
	}
	var sendChan = make(chan string, 10)
	var recvChan = make(chan string, 10)
	var conn *websocket.Conn
	return &websocketClientManager{
		addr:        &addr,
		protocol:    &protocol,
		path:        path,
		conn:        conn,
		sendMsgChan: sendChan,
		recvMsgChan: recvChan,
		isAlive:     false,
		timeout:     timeout,
	}
}

// 链接服务端
func (wsc *websocketClientManager) dail() {
	beginTime := time.Now().UnixMicro()
	var err error
	u := url.URL{Scheme: *wsc.protocol, Host: *wsc.addr, Path: wsc.path}
	//log.Printf("connecting to %s", u.String())
	wsc.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println(err)
		return
	}
	wsc.isAlive = true
	log.Printf("connecting to %s, success, %d(us)", u.String(), time.Now().UnixMicro()-beginTime)
}

// 发送消息
func (wsc *websocketClientManager) sendMsgThread() {
	go func() {
		msg := <-wsc.sendMsgChan
		err := wsc.conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Println("write:", err)
			return
		}
		log.Printf("send data to %s 成功", wsc.path)
	}()
}

// 读取消息
func (wsc *websocketClientManager) readMsgThread() {
	go func() {
		for {
			if wsc.conn != nil {
				_, message, err := wsc.conn.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					wsc.isAlive = false
					// 出现错误，退出读取，尝试重连
					break
				}
				log.Printf("recv: %s", message)
				// 需要读取数据，不然会阻塞
				wsc.recvMsgChan <- string(message)
			}

		}
	}()
}

// 开启服务并重连
func (wsc *websocketClientManager) start() {
	for {
		if wsc.isAlive == false {
			wsc.dail()
			wsc.sendMsgThread()
			wsc.readMsgThread()
		} else {
			//log.Println("connection success ", wsc.path)
			return
		}
		//	time.Sleep(time.Second * time.Duration(wsc.timeout))
	}
}

func batchTouchAndCheckConnection(conn ConnectionInfo) {
	for i := 0; i < conn.Num; i++ {
		go touchAndCheckConnection(conn)
	}
	log.Println("started all goroutines.")
	for {
		time.Sleep(1 * time.Second)
	}
}
func touchAndCheckConnection(conn ConnectionInfo) {
	protocol, address, path := findAddressAndPath(conn.RemoteAddress)
	if conn.Mode == "s" {
		for {
			wsc := NewWsClientManager(protocol, address, path, 1)
			wsc.start()
			//if wsc.isAlive == false {
			//	wsc.dail()
			//	wsc.sendMsgThread()
			//	wsc.readMsgThread()
			//}
			if wsc.isAlive {
				wsc.conn.Close()
			}
			time.Sleep(time.Duration(conn.Interval) * time.Second)
		}
	} else if conn.Mode == "m" {
		wsc := NewWsClientManager(protocol, address, path, 1)
		wsc.start()
		//if wsc.isAlive == false {
		//	wsc.dail()
		//	wsc.sendMsgThread()
		//	wsc.readMsgThread()
		//}
		if wsc.isAlive {
			wsc.conn.Close()
		}
	} else {
		log.Println("not support mode ")
	}
}
func findAddressAndPath(str string) (string, string, string) {
	strs := strings.Split(str, ":")

	if len(strs) < 3 {
		panic("incorrect remote address.")
		os.Exit(1)
	}
	protocol := strs[0]
	path := strs[2]
	i := strings.IndexAny(path, "/")
	i2 := strings.IndexAny(str, "//")

	path = string([]byte(path)[i:])
	address := str[i2+2 : len(str)-len(path)]
	return protocol, address, path
}

func main() {
	// 定义变量接收控制台参数
	// 完整请求路径地址 ws://127.0.0.1:8899/ws/v1/
	var remoteAddress string
	// 模式
	var mode string
	// 请求频率
	var interval int
	// 批量数
	var num int

	// StringVar用指定的名称、控制台参数项目、默认值、使用信息注册一个string类型flag，并将flag的值保存到p指向的变量
	flag.StringVar(&remoteAddress, "p", "ws://127.0.0.1:8899/ws/v1", "WebSocket服务器完整地址,默认为ws://127.0.0.1:8899/ws/v1/")
	flag.StringVar(&mode, "m", "s", "请求模式,s为单线程循环请求,m为多线程请求一次,默认为单线程")
	flag.IntVar(&interval, "i", 3, "请求间隔频率,默认为3秒请求一次")
	flag.IntVar(&num, "n", 10, "批量请求数,默认为10")

	// 从arguments中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()

	connInfo := ConnectionInfo{
		RemoteAddress: remoteAddress,
		Interval:      interval,
		Mode:          mode,
		Num:           num,
	}

	log.Println("start to testing...")
	bytes, _ := json.Marshal(connInfo)
	log.Println(string(bytes))
	switch connInfo.Mode {
	case "s":
		touchAndCheckConnection(connInfo)
	case "m":
		batchTouchAndCheckConnection(connInfo)
	default:
		log.Println("not support mode value ", connInfo.Mode)
		os.Exit(1)
	}
	log.Println("test finished.")
}
