package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/reatang/go-tapd/tapd/service"
)

func TestTask_GetTasks(t *testing.T) {
	c := getClient()

	resp, err := c.Task.GetTasks(context.Background(), &service.GetTaskRequest{
		QueryFields: service.QueryFields{
			Order: "modified desc",
			Limit: 10,
		},
	})
	if err != nil {
		t.Fatal(err)
		return
	}

	for _, datum := range resp.Data {
		fmt.Println(datum.Task.Name)
	}
}

func TestTask_GetTasksCount(t *testing.T) {
	c := getClient()

	resp, err := c.Task.GetTasksCount(context.Background(), &service.GetTaskRequest{})
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println(resp.Data.Count)
}
