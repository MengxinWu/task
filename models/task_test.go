package models

import (
	"testing"
)

func TestGenerateDAG(t *testing.T) {
	dagConfig := "{\"processors\":[1,2,3,4],\"relations\":[[1,2],[1,3],[2,4],[3,4]]}"
	dag, _ := GenerateDAG(dagConfig)
	for _, node := range dag {
		t.Logf("node: %v\n", node)
	}
}
