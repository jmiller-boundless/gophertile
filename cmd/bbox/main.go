package main

import (
	"encoding/json"
	"fmt"
	"gophertile/gophertile"
	"os"
	"strconv"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 3 && IsInt(argsWithoutProg[0]) && IsInt(argsWithoutProg[1]) && IsInt(argsWithoutProg[2]) {
		x, err := strconv.Atoi(argsWithoutProg[1])
		y, err := strconv.Atoi(argsWithoutProg[1])
		z, err := strconv.Atoi(argsWithoutProg[2])
		tile := gophertile.Tile{x, y, z}
		bounds := tile.Bounds()
		b, err := json.Marshal(bounds)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))
	} else {
		fmt.Println("Getting a bounding box requires x,y,z tile address parameters as integers")
	}
}

func IsInt(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}
