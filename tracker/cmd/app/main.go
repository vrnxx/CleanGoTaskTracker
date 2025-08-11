package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/common/events"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/aggregate"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo/priority"
	"log"
)

var (
	authorID = uuid.New()
)

func main() {
	title, err := vo.NewTitle("TestTitle")
	if err != nil {
		log.Fatal(err.Error())
	}
	desc, err := vo.NewDescription("Test Desctiption")
	if err != nil {
		log.Fatal(err.Error())
	}
	t := aggregate.NewTask(
		vo.TaskID{Value: uuid.New()},
		title,
		authorID,
		priority.Medium,
		desc,
		1,
	)
	fmt.Println(t)
	eventTest := events.NewBaseEvent("test")
	b, _ := eventTest.Bytes()
	fmt.Println(eventTest, b)
}
