package db

import (
	"fmt"
	"golang-study-practices/words/db/vo"
)

// 查询数据，指定字段名
func StructQuerySingle(wordName string, freq int) *vo.EnglishWords {
	word := new(vo.EnglishWords)
	row := MySQLDB.QueryRow("select id, word_name, sound_mark,paraphrase,frequency,memo from english_words where word_name=? or frequency=?", wordName, freq)
	if err := row.Scan(&word.Id, &word.WordName, &word.SoundMark, &word.Paraphrase, &word.Frequency, &word.Memo); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return nil
	}
	fmt.Printf("%#v row:", word)
	return word
}

// 查询数据，取所有字段
func StructQueryMultiple() []vo.EnglishWords {
	// 通过切片存储
	englishWords := make([]vo.EnglishWords, 0)
	rows, _ := MySQLDB.Query("SELECT * FROM `english_words` ratelimit ?", 100)
	// 遍历
	var word vo.EnglishWords
	for rows.Next() {
		rows.Scan(&word.Id, &word.WordName, &word.SoundMark, &word.Paraphrase, &word.Frequency, &word.Memo)
		englishWords = append(englishWords, word)
	}
	fmt.Println(englishWords)
	return englishWords
}

// 插入数据
func StructInsertT() {
	ret, _ := MySQLDB.Exec("insert INTO english_words(id, word_name, sound_mark,paraphrase,frequency,memo) values(?,?,?,?,?,?)", "小红", 23)

	//插入数据的主键id
	lastInsertID, _ := ret.LastInsertId()
	fmt.Println("LastInsertID:", lastInsertID)

	//影响行数
	rowsAffected, _ := ret.RowsAffected()
	fmt.Println("RowsAffected:", rowsAffected)

}

// 插入数据
func StructInsert(word vo.EnglishWords) {
	ret, _ := MySQLDB.Exec("insert INTO english_words(id, word_name, sound_mark,paraphrase,frequency,memo) values(?,?,?,?,?,?)", word.Id, word.WordName, word.SoundMark, word.Paraphrase, word.Frequency, word.Memo)
	//插入数据的主键id
	lastInsertID, err := ret.LastInsertId()
	fmt.Println("LastInsertID:", lastInsertID)

	if err != nil {
		fmt.Printf("err %s", err)
	}

	//影响行数
	rowsAffected, _ := ret.RowsAffected()
	fmt.Println("RowsAffected:", rowsAffected)
}
func StructBatchInsert(word vo.EnglishWords) {
	stmt, _ := MySQLDB.Prepare(`insert INTO english_words(id, word_name, sound_mark,paraphrase,frequency,memo) values(?,?,?,?,?,?)`)
	defer stmt.Close()

	ret, err := stmt.Exec(word.Id, word.WordName, word.SoundMark, word.Paraphrase, word.Frequency, word.Memo)
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	if LastInsertId, err := ret.LastInsertId(); nil == err {
		fmt.Println("LastInsertId:", LastInsertId)
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		fmt.Println("RowsAffected:", RowsAffected)
	}
}

// 更新数据
func StructUpdate() {

	ret, _ := MySQLDB.Exec("UPDATE english_words set sound_mark=? where id=?", "[ssss]", 1)
	updNums, _ := ret.RowsAffected()

	fmt.Println("RowsAffected:", updNums)
}

// 删除数据
func StructDel() {
	ret, _ := MySQLDB.Exec("delete from english_words where id=?", 1)
	delNums, _ := ret.RowsAffected()

	fmt.Println("RowsAffected:", delNums)
}

// 事务处理,结合预处理
func StructTx() {
	//事务处理
	tx, _ := MySQLDB.Begin()

	// 新增
	userAddPre, _ := MySQLDB.Prepare("insert into english_words(id,word_name, sound_mark) values(?, ?, ?)")
	addRet, _ := userAddPre.Exec("zhaoliu", 15)
	insNums, _ := addRet.RowsAffected()

	// 更新
	userUpdatePre1, _ := tx.Exec("update english_words set word_name = 'zhansan'  where word_name=?", "张三")
	updNums1, _ := userUpdatePre1.RowsAffected()
	userUpdatePre2, _ := tx.Exec("update english_words set word_name = 'lisi'  where word_name=?", "李四")
	updNums2, _ := userUpdatePre2.RowsAffected()
	fmt.Println(insNums)
	fmt.Println(updNums1)
	fmt.Println(updNums2)
	if insNums > 0 && updNums1 > 0 && updNums2 > 0 {
		tx.Commit()
	} else {
		tx.Rollback()
	}

}

// 查询数据，指定字段名,不采用结构体
func RawQueryField() {

	rows, _ := MySQLDB.Query("select id,id, word_name, soundmark, paraphrase, frequency from english_words")
	if rows == nil {
		return
	}
	id := 0
	name := ""
	fmt.Println(rows)
	fmt.Println(rows)
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
}

// 查询数据,取所有字段,不采用结构体
func RawQueryAllField() {

	//查询数据，取所有字段
	rows2, _ := MySQLDB.Query("select * from english_words")

	//返回所有列
	cols, _ := rows2.Columns()

	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))

	//这里表示一行填充数据
	scans := make([]interface{}, len(cols))
	//这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k]
	}

	i := 0
	result := make(map[int]map[string]string)
	for rows2.Next() {
		//填充数据
		rows2.Scan(scans...)
		//每行数据
		row := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		//放入结果集
		result[i] = row
		i++
	}
	fmt.Println(result)
}
