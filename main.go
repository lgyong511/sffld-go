package main

import (
	"fmt"

	"github.com/lgyong511/sffld-go/config/vp"
)

func main() {
	vp := vp.New().Set().Relod()
	conf := vp.Get()
	fmt.Printf("conf: %v\n", conf)
	select {}
}
