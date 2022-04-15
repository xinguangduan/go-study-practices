package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

func main() {
	fileContents := []byte("{\"event\":{    \"header\":    {        \"namespace\": \"SpeechRecognizer\",        \"name\": \"TranslateText\",        \"messageId\": \"34471018-2c1e-4610-8124-2d142b2wwwww\",        \"dialogRequestId\": \"b13dba86-232b-4ecf-9e34-e432d142b2wwwww\"    },    \"payload\":    {        \"text\": \"新型冠状病毒引起了世界范围大流行\",        \"language\":        {            \"source\": \"auto\",            \"target\": \"en-US\"        }    }}\n}")

	var b bytes.Buffer
	w := multipart.NewWriter(&b) //返回一个设定了一个随机boundary的Writer w，并将数据写入&b
	{
		//part, err := w.CreateFormFile("myfile", "my-file.txt") //使用给出的属性名（对应name）和文件名（对应filename）创建一个新的form-data头，part为io.Writer类型
		//if err != nil {
		//	fmt.Printf("CreateFormFile: %v\n", err)
		//}
		w.WriteField("", string(fileContents))

		//part ,_:= w.CreatePart(textproto.MIMEHeader{
		//	"Content-Type": []string{"boundary=--"+w.Boundary()+";application/json;charset=UTF-8;"},
		//})
		//
		//part.Write([]byte("--"+w.Boundary()))
		//part.Write([]byte("\r\n"))
		//part.Write(fileContents)         //然后将文件的内容添加到form-data头中
		//part.Write([]byte("\r\n"))
		//part.Write([]byte("--"+w.Boundary()))
		s := b.String()
		if len(s) == 0 {
			fmt.Println("String: unexpected empty result")
		}
		fmt.Println(s)
	}
	//fmt.Println(w.Boundary()) //随机生成的boundary为284b0f2fc979a7e51d4e056a96b32ea8f8d94301287968d78723bd0113e9
	//r := multipart.NewReader(&b, w.Boundary())

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	token := "eyJlbmMiOiJBMTI4R0NNIiwiYWxnIjoiZGlyIn0..sWJzQqwtXjhU7Bpk.FGYS2EvuvWx887aFYjyHoK4weZlN96JBakaewWiGwHfMC2tTNbazSictvG-rvUH1mGSmjgWLUXGXMV7mYqrUKevdA5CcHQkltC2Vl44WzmfoYlNctsgDnz4pn7VuCs42tBITMaDOp4W2vvkgxC3llhKBGWS0Fn5wY09UlnDmzue3IbTJL-OKFb74_fYN-3_g6gvXMCITaGArqG3FdkWwHm-kUo8KUU0PL1AVKto56lZH3XIBetaJ8bEdYJrYjOJH2-byg5IDWMmMcBI5E1kXd5q0rcZcWzpj1DPjMJTYJsU24Auq8lYAgQ8HqZZAPuWy6rv3c9GQaQDFgiqYroBfdItMVl7RUWSOq-R1WXjrCAV91aybyhWbwKisJPDfqmznAB9K2SRDRDTw_Ny9x7A-1Ea_ykbPl-VOhTYYqbsyJX4s_ms4SWoXgeAjriJ6kX7wNPHfPl_45RAUO-KcV6y_gIhedbztrwluxNOHN79aEPt-4sSSLetS6-bjmRvHLvVVQBOfh-WdjevHEH6pbtV2wWqQjaSGjAQuTo4-cdcoZ5oAOMPIg3OhL1TxGmlyy2UZqfgmzJFFdhNTIVYHmdVd2j4p7R9rnN4sc-DFzL5E4r1csXsBrj0-Pg4qdYG72PRBDMqxE38HxHLx6lxsTOIEwoNXQ5sq9n9BxK1PhRbo_T7Ldqiqglnynkt8ALYxyLZ5KdM__ARnrZVz3iMQk_hKifzMhwQTcBQiQhVxQdHIsNdVvqh3oCmPyLgze3N9asypdZ-gfrWJmCN_KNnZtVJ-unUgZp0XxrUmjP1a2l_dot2PtWAPdED5EIoHaHKZse3IY3FaW5QMdlQZCxKLObQdrtM-EsttmfJsfVHvGtu5B8Nzn50GSWccbDzy3I2gbX0SjsOrLD8vNNONJ5FYkDkCMcunP46GrF4wplH80RMwfAwT_4zwEnrGZmz3fkAv2Ys7nV4r_5AyiODVTCbQnLSid5AMDYOulhmg-qbFrpzYKmt_vFXzRicFbmXx6xfRgD_fgbiTHQ.y3uztq8ahKlIZ2Ho1id1xA"
	// resp, err := http.Post(destURL, writer.FormDataContentType(), bufReader)

	req, _ := http.NewRequest("POST", "https://pvtcui.lenovo.com.cn/v20180430/events", &b)

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
