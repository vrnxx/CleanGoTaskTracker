package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/common/events"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/aggregate"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo"
	"github.com/vrnxx/CleanGoTaskTracker/tracker/internal/domain/task/vo/task_priority"
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
		new(uuid.UUID),
		task_priority.Medium,
		desc,
		1,
	)
	fmt.Printf(
		"[Task Created]\n\tID: %s\n\tAssigneID: %s\n\tStatus: %s\n\tDescription: %s\n\tPriority: %s\n\n",
		t.ID.AsGenericType(),
		t.AssigneeID,
		t.Status,
		t.Description,
		t.Priority,
	)
	eventTest := events.NewBaseEvent("test")
	b, _ := eventTest.Bytes()
	fmt.Println(eventTest, b)
	err = t.DeleteTask()
	//err = t.DeleteTask()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(t)
}
