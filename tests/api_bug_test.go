package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/reatang/go-tapd/tapd/service"
)

func TestBug_GetBugs(t *testing.T) {
	c := getClient()
	resp, err := c.Bug.GetBugs(context.Background(), &service.GetBugsRequest{
		QueryFields: service.QueryFields{
			Order: "modified desc",
			Limit: 10,
		},
	})
	if err != nil {
		t.Fatal(err)
		return
	}

	for _, bug := range resp.Data {
		fmt.Println(bug.Bug.Title)
	}
}

func TestBug_GetBugCount(t *testing.T) {
	c := getClient()
	resp, err := c.Bug.GetBugsCount(context.Background(), &service.GetBugsRequest{})
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println(resp.Data.Count)
}
