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

func TestReleases_GetReleases(t *testing.T) {
	w, _ := strconv.Atoi(os.Getenv("TAPD_API_WorkspaceID"))
	c := tapd.NewTAPDClient(tapd.WithBasicAuth(os.Getenv("TAPD_API_USER"), os.Getenv("TAPD_API_PASSWORD")), tapd.WithWorkspaceID(w))

	resp, err := c.Releases.GetReleases(context.Background(), &service.GetReleasesRequest{
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
		fmt.Println(datum.Release.StartDate)
	}
}

func TestReleases_GetReleasesCount(t *testing.T) {
	w, _ := strconv.Atoi(os.Getenv("TAPD_API_WorkspaceID"))
	c := tapd.NewTAPDClient(tapd.WithBasicAuth(os.Getenv("TAPD_API_USER"), os.Getenv("TAPD_API_PASSWORD")), tapd.WithWorkspaceID(w))

	resp, err := c.Releases.GetReleasesCount(context.Background(), &service.GetReleasesRequest{})
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println(resp.Data.Count)
}
