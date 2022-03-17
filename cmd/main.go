package main

import (
	"math/rand"
	"time"

	"github.com/DumDumGeniuss/ggol"
	"github.com/gin-gonic/gin"
)

var g ggol.Game
var count int
var width int = 120
var height int = 75
var size ggol.Size
var period time.Duration = 100

func initGame() ggol.Game {
	size = ggol.Size{Width: width, Height: height}
	generation := make(ggol.Generation, width)
	for x := 0; x < width; x++ {
		generation[x] = make([]ggol.Cell, height)
		for y := 0; y < height; y++ {
			generation[x][y] = rand.Intn(2) == 0
		}
	}
	seed := ggol.ConvertGenerationToSeed(generation)
	newG, _ := ggol.NewGame(
		&size,
		&seed,
	)
	// newG.SetShouldCellDie(func(liveNbrsCount int, c *ggol.Coordinate) bool {
	// 	return true
	// })
	return newG
}

func heartBeat() {
	for range time.Tick(time.Millisecond * period) {
		count++
		if count == 1000 {
			count = 0
			g = initGame()
		}
		g.Evolve()
	}
}

func main() {
	g = initGame()
	go heartBeat()

	route := gin.Default()
	route.GET("/api/generation", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"size":       g.GetSize(),
			"period":     period,
			"generation": *g.GetGeneration(),
		})
	})
	route.Static("/demo", "./cmd/public")
	route.Run(":8000")
}
