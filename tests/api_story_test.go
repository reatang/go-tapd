package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/reatang/go-tapd/tapd/service"
)

func TestStory_GetStories(t *testing.T) {
	c := getClient()

	resp, err := c.Story.GetStories(context.Background(), &service.GetStoriesRequest{
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
		fmt.Println(datum.Story.Name)
	}
}

func TestStory_GetStoriesCount(t *testing.T) {
	c := getClient()

	resp, err := c.Story.GetStoriesCount(context.Background(), &service.GetStoriesRequest{})
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println(resp.Data.Count)
}
