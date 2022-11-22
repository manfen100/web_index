package lesson09

import (
	"errors"
	"fmt"
)

func showBookInfo(bookName, authorName string) (string, error) {
	if bookName == "" {
		return "", errors.New("图书名称为空")
	}
	if authorName == "" {
		return "", errors.New("作者名称为空")
	}

	return bookName + ",作者" + authorName, nil
}
func main() {
	bookInfo, err := showBookInfo("《和》", "欢喜")
	fmt.Println("bookInfo = %s,error = %v", bookInfo, err)
}

//
//func(parameter_list)(re){
//
//}
