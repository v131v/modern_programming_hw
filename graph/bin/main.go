package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/v131v/modern_programming_hw/graph/src/graph"
)

func main() {
	var file string
	var source, drain int

	flag.StringVar(&file, "f", "", "File with graph in json format")
	flag.IntVar(&source, "s", 0, "Source vertex")
	flag.IntVar(&drain, "d", 0, "Drain vertex")
	flag.Parse()

	g := &graph.Graph{}

	graphConfig, err := os.Open(file)
	if err != nil {
		fmt.Printf("error opening file: %w", err)
		os.Exit(1)
	}
	defer graphConfig.Close()

	decoder := json.NewDecoder(graphConfig)

	err = decoder.Decode(&g)
	if err != nil {
		fmt.Printf("error decoding JSON: %w", err)
		os.Exit(1)
	}

	g.Init()

	path, err := g.FindPath(source, drain)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", path)
}
