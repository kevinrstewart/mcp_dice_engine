package main

import (
	"fmt"
	"kevinrstewart/mcp_dice_enginge/src/engine"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	results := make([]int, 50000)

	t := time.Now().UnixMilli()

	for i := 0; i < len(results); i++ {
		results[i] = engine.GenerateResult(5, nil, 3, nil)
	}

	analysis := engine.AnalyzeResults(results)

	t = time.Now().UnixMilli() - t

	fmt.Println(fmt.Sprintf("%d Iterations:\n%v\nTotal Time (millis): %v", len(results), analysis, t))
}
