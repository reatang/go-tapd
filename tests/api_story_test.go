package tests

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/reatang/go-tapd/tapd"
	"github.com/reatang/go-tapd/tapd/service"
)

func TestStory_GetStories(t *testing.T) {
	w, _ := strconv.Atoi(os.Getenv("TAPD_API_WorkspaceID"))
	c := tapd.NewTAPDClient(tapd.WithBasicAuth(os.Getenv("TAPD_API_USER"), os.Getenv("TAPD_API_PASSWORD")), tapd.WithWorkspaceID(w))

	resp, err := c.Story.GetStories(context.Background(), &service.GetStoriesRequest{
		QueryRequest: service.QueryRequest{
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
	w, _ := strconv.Atoi(os.Getenv("TAPD_API_WorkspaceID"))
	c := tapd.NewTAPDClient(tapd.WithBasicAuth(os.Getenv("TAPD_API_USER"), os.Getenv("TAPD_API_PASSWORD")), tapd.WithWorkspaceID(w))

	resp, err := c.Story.GetStoriesCount(context.Background(), &service.GetStoriesRequest{})
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println(resp.Data.Count)
}
