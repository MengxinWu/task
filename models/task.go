package models

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	ProcessorId int
	Parents     []*Node
	Children    []*Node
}

type DAG map[int]*Node

func (d DAG) AddNode(processorId int) {
	d[processorId] = &Node{ProcessorId: processorId}
}

func (d DAG) AddEdge(relation []int) {
	from := relation[0]
	to := relation[1]
	d[from].Children = append(d[from].Children, d[to])
	d[to].Parents = append(d[to].Parents, d[from])
}

type DagConfig struct {
	Processors []int   `json:"processors"`
	Relations  [][]int `json:"relations"`
}

// GenerateDAG generate DAG.
func GenerateDAG(dagConfig string) (DAG, error) {
	var err error
	config := new(DagConfig)
	if err = json.Unmarshal([]byte(dagConfig), &config); err != nil {
		return nil, err
	}
	fmt.Printf("dag config: %v\n", config)
	dag := make(DAG)
	for _, processorId := range config.Processors {
		dag.AddNode(processorId)
	}
	for _, relations := range config.Relations {
		dag.AddEdge(relations)
	}
	return dag, nil
}
