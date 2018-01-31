package main

import (
	"os"
	"strconv"
	"gophertile/gophertile"
	"encoding/json"
	"fmt"
	"math"
)

type TileBox struct {
	Tile *gophertile.Tile
	Box *gophertile.LngLatBbox
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 1 && IsInt(argsWithoutProg[0]) {
		z, err := strconv.Atoi(argsWithoutProg[0])
		maxdim := int(math.Pow(float64(2), float64(z)))
		var bounds []*TileBox
		for i := 0; i < maxdim; i++ {
			for j := 0; j < maxdim; j++ {
				tile := gophertile.Tile{i, j, z}
				bound := tile.Bounds()
				tb := new(TileBox)
				tb.Box = bound
				tb.Tile = &tile
				bounds = append(bounds, tb)
			}
		}

		b, err := json.Marshal(bounds)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))
	} else {
		fmt.Println("Getting bounding boxes for the zoom level requires zoom level parameters as an integer")
	}
}

func IsInt(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}
