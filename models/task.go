package models

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type Node struct {
	ProcessorId int
	Parents     []*Node
	Children    []*Node
}

type Graph map[int]*Node

func (d Graph) AddNode(processorId int) {
	d[processorId] = &Node{ProcessorId: processorId}
}

func (d Graph) AddEdge(relation []int) {
	from := relation[0]
	to := relation[1]
	d[from].Children = append(d[from].Children, d[to])
	d[to].Parents = append(d[to].Parents, d[from])
}

type DagConfig struct {
	Processors []int   `json:"processors"`
	Relations  [][]int `json:"relations"`
}

// GenerateGraph generate Graph.
func GenerateGraph(dagConfig string) (Graph, error) {
	var err error
	config := new(DagConfig)
	if err = json.Unmarshal([]byte(dagConfig), &config); err != nil {
		return nil, err
	}
	log.Printf("dag config: %", config)
	dag := make(Graph)
	for _, processorId := range config.Processors {
		dag.AddNode(processorId)
	}
	for _, relations := range config.Relations {
		dag.AddEdge(relations)
	}
	return dag, nil
}
