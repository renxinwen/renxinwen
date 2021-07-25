package main

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)
type Book struct {
	ID int64
}

func main() {
	err := getBookInfo()
	if err != nil {
		log.Printf("%T %v", errors.Cause(err), errors.Cause(err))
		log.Printf("堆栈信息:%+v", err)
	}
	return
}

func getBookInfo()(err error)  {
	book, err := connDB()
	if err != nil{
		return errors.WithMessage(err, "获取book信息失败")
	}
	fmt.Println(book.ID)
	return
}

func connDB() ( book Book, err error) {
	dsn := "XXX"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return
	}
	table := "books"
	err = db.Table(table).Select("id").Where("id = ?", 0).First(&book).Error
	return book, errors.Wrap(err, "没有获取到book信息")
}
