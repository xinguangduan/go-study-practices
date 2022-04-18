package vo

import "fmt"

type EnglishWords struct {
	Id         int
	WordName   string
	SoundMark  string
	Paraphrase string
	Frequency  int
	Memo       string
}

func main() {
	var words EnglishWords
	words.Id = 1
	words.WordName = "dictator"
	words.SoundMark = "[dicteite]"
	words.Frequency = 1100
	words.Memo = "none"
	words.Paraphrase = "独裁者"

	var words2 EnglishWords = EnglishWords{
		Id:         1,
		WordName:   "dictator",
		SoundMark:  "sss",
		Frequency:  12200,
		Paraphrase: "独裁者",
		Memo:       "none",
	}
	fmt.Printf(words2.SoundMark)
}

func outputStructContent() {
	var words EnglishWords = EnglishWords{
		Id:         1,
		WordName:   "dictator",
		SoundMark:  "sss",
		Paraphrase: "独裁者",
		Frequency:  12200,
		Memo:       "none",
	}

	fmt.Printf("struct init style1:%#v\n", words)

	words2 := EnglishWords{
		Id:         2,
		WordName:   "user-friendly",
		SoundMark:  "[jus-frendly]",
		Paraphrase: "为用户着想",
		Frequency:  12000,
		Memo:       "为用户着想的意思",
	}

	fmt.Printf("struct init style2:%#v\n", words2)

	var words3 *EnglishWords = &EnglishWords{}
	(*words3).Id = 3
	(*words3).Frequency = 2000

	fmt.Printf("struct init style3:%#v\n", words3)
}
