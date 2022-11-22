package lesson10

//简化导入包语句
import . "fmt"

//匿名导入
import _ "fmt"

func main() {
	Println("hello")
	chinese := 88
	math := 20
	english := 95

	switch getResult(chinese, math, english) {
	case true:
		Println("pass")
	case false:
		Println("fail")

	}
	si()

}

func con(n string) (string, error) {
	if socre := 88; socre >= 60 {
		Println("")
	}
	return "2", nil

}

func getResult(args ...int) bool {
	for _, v := range args {
		if v < 60 {
			return false
		}
	}
	return true
}

func si() {
	s := "从0到Go语言微服务架构师"
	switch {
	case s == "从0到Go语言微服务架构师":
		Println("从0到Go语言微服务架构师")
		fallthrough
	case s == "Go语言微服务架构核心22讲":
		Println("Go语言微服务架构核心22讲")
	case s != "Go语言极简一本通":
		Println("Go语言极简一本通")
	}
}
