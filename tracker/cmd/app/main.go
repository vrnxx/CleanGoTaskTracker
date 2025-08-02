package main

import (
	"fmt"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo"
)

func main() {
	fmt.Printf("application start successfull\n\n")
	t, err := vo.Title{}.Create("Hello")
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}
	fmt.Println(t, t.AsGenericType())
}
