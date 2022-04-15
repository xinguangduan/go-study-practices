package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

const (
	destURL = "https://pvtcui.lenovo.com.cn/v20180430/events"
)

func main() {
	var bufReader bytes.Buffer
	mpWriter := multipart.NewWriter(&bufReader)
	fw, err := mpWriter.CreateFormFile("upload_file", "/Users/charles/Documents/ownprojects/goland-study-practices/multipart/a.json")
	if err != nil {
		fmt.Println("Create form file error: ", err)
		return
	}

	f, _ := os.Open("/Users/charles/Documents/ownprojects/goland-study-practices/multipart/a.json")
	_, err = io.Copy(fw, f)
	if err != nil {
		return
	}

	defer mpWriter.Close()

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	token := "eyJlbmMiOiJBMTI4R0NNIiwiYWxnIjoiZGlyIn0..sWJzQqwtXjhU7Bpk.FGYS2EvuvWx887aFYjyHoK4weZlN96JBakaewWiGwHfMC2tTNbazSictvG-rvUH1mGSmjgWLUXGXMV7mYqrUKevdA5CcHQkltC2Vl44WzmfoYlNctsgDnz4pn7VuCs42tBITMaDOp4W2vvkgxC3llhKBGWS0Fn5wY09UlnDmzue3IbTJL-OKFb74_fYN-3_g6gvXMCITaGArqG3FdkWwHm-kUo8KUU0PL1AVKto56lZH3XIBetaJ8bEdYJrYjOJH2-byg5IDWMmMcBI5E1kXd5q0rcZcWzpj1DPjMJTYJsU24Auq8lYAgQ8HqZZAPuWy6rv3c9GQaQDFgiqYroBfdItMVl7RUWSOq-R1WXjrCAV91aybyhWbwKisJPDfqmznAB9K2SRDRDTw_Ny9x7A-1Ea_ykbPl-VOhTYYqbsyJX4s_ms4SWoXgeAjriJ6kX7wNPHfPl_45RAUO-KcV6y_gIhedbztrwluxNOHN79aEPt-4sSSLetS6-bjmRvHLvVVQBOfh-WdjevHEH6pbtV2wWqQjaSGjAQuTo4-cdcoZ5oAOMPIg3OhL1TxGmlyy2UZqfgmzJFFdhNTIVYHmdVd2j4p7R9rnN4sc-DFzL5E4r1csXsBrj0-Pg4qdYG72PRBDMqxE38HxHLx6lxsTOIEwoNXQ5sq9n9BxK1PhRbo_T7Ldqiqglnynkt8ALYxyLZ5KdM__ARnrZVz3iMQk_hKifzMhwQTcBQiQhVxQdHIsNdVvqh3oCmPyLgze3N9asypdZ-gfrWJmCN_KNnZtVJ-unUgZp0XxrUmjP1a2l_dot2PtWAPdED5EIoHaHKZse3IY3FaW5QMdlQZCxKLObQdrtM-EsttmfJsfVHvGtu5B8Nzn50GSWccbDzy3I2gbX0SjsOrLD8vNNONJ5FYkDkCMcunP46GrF4wplH80RMwfAwT_4zwEnrGZmz3fkAv2Ys7nV4r_5AyiODVTCbQnLSid5AMDYOulhmg-qbFrpzYKmt_vFXzRicFbmXx6xfRgD_fgbiTHQ.y3uztq8ahKlIZ2Ho1id1xA"
	// resp, err := http.Post(destURL, writer.FormDataContentType(), bufReader)

	req, _ := http.NewRequest("POST", destURL, &bufReader)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", token)
	req.Header.Set("Accept", "*/*")
	response, _ := client.Do(req)
	fmt.Println(response.StatusCode)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
	}
}
