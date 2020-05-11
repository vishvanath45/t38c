// +build ignore

package main

import (
	"log"

	t38c "github.com/powercake/tile38-client"
)

func main() {
	tile38, err := t38c.New("localhost:9851", t38c.Debug())
	if err != nil {
		log.Fatal(err)
	}

	// To set a field when setting an object:
	if err := tile38.Set("fleet", "truck1", t38c.SetPoint(33.5123, -112.2693)).
		Field("speed", 90).
		Field("age", 21).
		Do(); err != nil {
		log.Fatal(err)
	}

	// To set a field when an object already exists:
	if err := tile38.FSet("fleet", "truck1").Field("speed", 90).Do(); err != nil {
		log.Fatal(err)
	}
}
