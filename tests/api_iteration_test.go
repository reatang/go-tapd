package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/reatang/go-tapd/tapd/service"
)

func TestIteration_GetIteration(t *testing.T) {
	c := getClient()
	resp, err := c.Iteration.GetIterations(context.Background(), &service.GetIterationsRequest{
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
		fmt.Println(datum.Iteration.Name)
	}
}

func TestIteration_GetIterationCount(t *testing.T) {
	c := getClient()
	resp, err := c.Iteration.GetIterationsCount(context.Background(), &service.GetIterationsRequest{})
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println(resp.Data.Count)
}
