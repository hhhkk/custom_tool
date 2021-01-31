package log

import "fmt"

func E(error ...error)  {
	fmt.Println(error)
}


func Fatal(error ...error)  {
	panic(error)
}
