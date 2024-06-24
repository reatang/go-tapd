package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/reatang/go-tapd/tapd/service"
)

func TestReleases_GetReleases(t *testing.T) {
	c := getClient()

	req := &service.GetReleasesRequest{
		QueryFields: service.QueryFields{
			Limit: 10,
		},
	}
	req.OrderByModifiedDesc()
	req.WhereModifiedGt(time.Date(2024, 06, 13, 0, 0, 0, 0, time.Local))
	req.Select("workspace_id", "name", "modified")

	resp, err := c.Releases.GetReleases(context.Background(), req)
	if err != nil {
		t.Fatal(err)
		return
	}

	for _, datum := range resp.Data {
		fmt.Println(datum.Release.Name, datum.Release.Modified)
	}
}

func TestReleases_GetReleasesCount(t *testing.T) {
	c := getClient()

	resp, err := c.Releases.GetReleasesCount(context.Background(), &service.GetReleasesRequest{})
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println(resp.Data.Count)
}
