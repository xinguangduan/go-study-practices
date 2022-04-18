package vo

import "strconv"

type Student struct {
	Name   string
	age    int
	Phone  string
	Gender string
	Mail   string
}

// 获取数据
func getStudents() []Student {
	students := make([]Student, 0)
	for i := 0; i < 10; i++ {
		stu := Student{}
		stu.Name = "name" + strconv.Itoa(i+1)
		stu.Mail = stu.Name + "@chairis.cn"
		stu.Phone = "1380013800" + strconv.Itoa(i)
		stu.age = 20
		stu.Gender = "男"
		students = append(students, stu)
	}
	return students
}
