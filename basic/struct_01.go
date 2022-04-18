package main

import "fmt"

type Book struct {
	Name    string
	Price   float64
	Address string
}

func (b Book) GetName() string {
	return b.Name
}
func (b Book) SetName(newValue string) {
	b.Name = newValue
}
func (b *Book) SetName1(newValue string) {
	b.Name = newValue
}
func main() {

	book1 := new(Book)
	book1.Name = "钢铁是怎么练成的"
	book1.Price = 100.34
	book1.Address = "beijing"
	fmt.Println(book1)
	fmt.Println(*book1)

	var book2 Book
	book2.Name = "大清盐商"
	book2.Price = 132.43
	book2.Address = "The qing dynasty of China"
	fmt.Println(book2)

	changeBookValue(&book2)

	fmt.Println(book2)

	// 简洁的创建对象方式
	book3 := Book{Name: "sss", Price: 982.333, Address: "China's song dynasty"}
	book3.Name = "清明上河图"
	fmt.Println(book3)

	fmt.Println(book3.GetName())
	book3.SetName("富春山居图")
	fmt.Println(book3)

}
func changeBookValue(book *Book) {
	book.Name = "li bai"
	book.Price = 10203.2
	book.Address = "China's tang dynasty"
}
