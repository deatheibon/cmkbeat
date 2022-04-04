package main

import (
	"github.com/elastic/beats/libbeat/beat"
	"github.com/deatheibon/cmkbeat/beater"
)

func main() {
	beat.Run("cmkbeat", "", beater.New)
}
