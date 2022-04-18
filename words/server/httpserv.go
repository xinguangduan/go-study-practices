package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang-study-practices/words/db"
	"golang-study-practices/words/db/vo"
	"net/http"
	"strconv"
)

func InitServer() {
	//1.调用 httprouter 生成句柄
	router := httprouter.New()

	//2.通过句柄进行路由解析
	router.GET("/words/getuser/:uuid", GetUser)
	//2.通过句柄进行路由解析
	router.GET("/words/getword/:word", GetWord)

	//3.将句柄放入http.handle
	http.Handle("/", router)

	//4.建立网络,等待网络连接
	http.ListenAndServe(":10010", nil)
}

//解析请求中带的参数
func GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	uid := params.ByName("uuid")
	fmt.Fprintf(w, "获取到用户ID:%s\n", uid)
	words2 := vo.EnglishWords{
		Id:         2,
		WordName:   "user-friendly",
		SoundMark:  "[jus-frendly]",
		Paraphrase: "为用户着想",
		Frequency:  12000,
		Memo:       "为用户着想的意思",
	}
	fmt.Fprintf(w, "获取到word object:%#v\n", words2)

	//获取请求方式
	mo := r.Method
	fmt.Fprintf(w, "获取到请求方式为:%s", mo)
}

func GetWord(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	wordName := params.ByName("word")
	freq, _ := strconv.Atoi(params.ByName("freq"))

	word := db.StructQuerySingle(wordName, freq)

	fmt.Printf("%#v\n", word)

	if word == nil {
		fmt.Printf("")
		fmt.Fprintf(w, "not found any word for :%s\n", word)
		return
	}

	jData, _ := json.Marshal(&word)

	var formattedJSON bytes.Buffer
	_ = json.Indent(&formattedJSON, jData, "", "    ")
	w.Write([]byte(formattedJSON.String()))

}

func outputStructContent() {
	var words vo.EnglishWords = vo.EnglishWords{
		Id:         1,
		WordName:   "dictator",
		SoundMark:  "sss",
		Paraphrase: "独裁者",
		Frequency:  12200,
		Memo:       "none",
	}

	fmt.Printf("struct init style1:%#v\n", words)

	words2 := vo.EnglishWords{
		Id:         2,
		WordName:   "user-friendly",
		SoundMark:  "[jus-frendly]",
		Paraphrase: "为用户着想",
		Frequency:  12000,
		Memo:       "为用户着想的意思",
	}

	fmt.Printf("struct init style2:%#v\n", words2)

	var words3 *vo.EnglishWords = &vo.EnglishWords{}
	(*words3).Id = 3
	(*words3).Frequency = 2000

	fmt.Printf("struct init style3:%#v\n", words3)
}
