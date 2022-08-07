package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

var headerContentTypeJson = []byte("application/json")

var client *fasthttp.Client

type Entity struct {
	Id   int
	Name string
}

func main() {
	// You may read the timeouts from some config
	readTimeout, _ := time.ParseDuration("500ms")
	writeTimeout, _ := time.ParseDuration("500ms")
	maxIdleConnDuration, _ := time.ParseDuration("1h")
	client = &fasthttp.Client{
		ReadTimeout:                   readTimeout,
		WriteTimeout:                  writeTimeout,
		MaxIdleConnDuration:           maxIdleConnDuration,
		NoDefaultUserAgentHeader:      true, // Don't send: User-Agent: fasthttp
		DisableHeaderNamesNormalizing: true, // If you set the case on your headers correctly you can enable this
		DisablePathNormalizing:        true,
		// increase DNS cache time to an hour instead of default minute
		Dial: (&fasthttp.TCPDialer{
			Concurrency:      4096,
			DNSCacheDuration: time.Hour,
		}).Dial,
	}
	//sendGetRequest()
	for i := 0; i < 1; i++ {
		sendPostRequest()
	}

}

func sendGetRequest() {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://localhost:8080/")
	req.Header.SetMethod(fasthttp.MethodGet)
	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	if err == nil {
		fmt.Printf("DEBUG Response: %s\n", resp.Body())
	} else {
		fmt.Fprintf(os.Stderr, "ERR Connection error: %v\n", err)
	}
	fasthttp.ReleaseResponse(resp)
}

type Monitor struct {
	Event Event `json:"event"`
}

type Header struct {
	Namespace       string `json:"namespace"`
	Name            string `json:"name"`
	DialogRequestID string `json:"dialogRequestId"`
	MessageId       string `json:"messageId"`
}

type Payload struct {
}

type Event struct {
	Header  Header  `json:"header"`
	Payload Payload `json:"payload"`
}

func sendPostRequest() {
	// per-request timeout
	reqTimeout := time.Duration(2000) * time.Millisecond

	//reqEntity := &Entity{
	//	Name: "New entity",
	//}
	//reqEntityBytes, _ := json.Marshal(reqEntity)
	//
	reqBody := `
{
   "event":{
       "header":{
           "namespace": "FDNController",
           "name": "GetActiveMonitors",
           "dialogRequestId": "2316192e-1ec6-4e10-9b22-3581516ce0123",
"messageId":"ddss"
       },
       "payload":{

       }
   }
}
`

	reqEntity := &Monitor{
		Event: Event{
			Header: Header{
				Namespace:       "FDNController",
				Name:            "GetActiveMonitors",
				DialogRequestID: GenUUID(),
			},
			Payload: Payload{},
		},
	}
	reqEntityBytes, _ := json.Marshal(reqEntity)
	m := &Monitor{}
	json.Unmarshal([]byte(reqBody), m)
	fmt.Println(m)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI("https://pvtcuiapi.lenovo.com.cn/fdnservice/v1.0/monitor")
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentTypeBytes(headerContentTypeJson)
	req.Header.Set("authorization", "Bearer eyJlbmMiOiJBMTI4R0NNIiwiYWxnIjoiZGlyIn0..PGGlmZRwjqZVeFf8.hPzWQ4jOT6VTY4d_NE7V-TQ9YdZCpvAazBEmiSUVdJLr8vfnluS0lVhs1CbuWDXsj6VcvdJue4JaPHQZD6NqVmbIjgBC_0MQGgxfeeangHOOJpcxANZTvvKUEYbbeSl6jfflNMjXxrXMkCfjFnU8OBdNifbdxCuw-spz996KSradxh75s6GnAEmNt3mKcC3aSsSB5RkqEj7_wc5bzhG8b2iFMPsQZsrARhP9iE1MJ3GDz1LeJWWxtRWgWGgdnm2MPiINYDs5rT4_wpXvzyk-xMzyqFIqswPrZupRyeSDM4fcNdHHckNUdVpWVobuXihrr19_W8szV_h2UPM_dh5pSmrNqQZYR6mbL90AZvif33lLP7o6iKXPhpYcG4xO_Ncst7rS-_otWuyHfkxsYQbQEj5K_632eFTF_IVDXm0R1pV7O9fB5mh0hHsAy13Tt0-Jx1eKBu3IUWYDjJ1GV2SYID0tk-hRpzBWIsHXRhdovnVHVsnABY0a7U2Rn8RjcuCm9hDX-hp7md6xvrzjxuFWU0boqnpxI0Yg0_vUWoh8KB3ZyrXKIJoSCL4cfErbu7JupWPAL9kgz4eObrSSu8g9JnHIYg2H4kypaVpIOjVsKxAFA1iBYIijwyu8FbYzzdWdVvSfNw4rAaGwELpUlNWCbHbwG3YTL0cqc1yB2tRKP_4Ddx_K2zpMXG6bnbhagc8iUmn8jDJy8HO1XKI8eIKfKK6GIKIsswBmFlgwykTgrlqV4jTd0Js5.szO-SFo76nFdbaLrMHdlJw")
	//req.SetBodyString(reqBody)
	req.SetBody(reqEntityBytes)
	resp := fasthttp.AcquireResponse()
	err := client.DoTimeout(req, resp, reqTimeout)
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	if err == nil {
		statusCode := resp.StatusCode()
		respBody := resp.Body()
		fmt.Printf("DEBUG Response: %s\n", respBody)
		if statusCode == http.StatusOK {
			//fmt.Println(respBody)
			//respEntity := &Entity{}
			//err = json.Unmarshal(respBody, respEntity)
			//if err == io.EOF || err == nil {
			//	fmt.Printf("DEBUG Parsed Response: %v\n", respEntity)
			//} else {
			//	fmt.Fprintf(os.Stderr, "ERR failed to parse reponse: %v\n", err)
			//}
		} else {
			fmt.Fprintf(os.Stderr, "ERR invalid HTTP response code: %d\n", statusCode)
		}
	} else {
		errName, known := httpConnError(err)
		if known {
			fmt.Fprintf(os.Stderr, "WARN conn error: %v\n", errName)
		} else {
			fmt.Fprintf(os.Stderr, "ERR conn failure: %v %v\n", errName, err)
		}
	}

}

func GenUUID() string {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
	//fmt.Printf("%s", out)
	return strings.ReplaceAll(string(out), "\n", "")
}

func httpConnError(err error) (string, bool) {
	errName := ""
	known := false
	if err == fasthttp.ErrTimeout {
		errName = "timeout"
		known = true
	} else if err == fasthttp.ErrNoFreeConns {
		errName = "conn_limit"
		known = true
	} else if err == fasthttp.ErrConnectionClosed {
		errName = "conn_close"
		known = true
	} else {
		errName = reflect.TypeOf(err).String()
		if errName == "*net.OpError" {
			// Write and Read errors are not so often and in fact they just mean timeout problems
			errName = "timeout"
			known = true
		}
	}
	return errName, known
}
