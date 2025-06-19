package main

import (
	"fmt"

	"github.com/lgyong511/sffld-go/config/vp"
)

func main() {
	vp := vp.New()
	conf := vp.Get()
	fmt.Printf("conf: %v\n", conf)
	conf.App.AuthTimeout = 10
	// vp.MergeConfigMap()
	fmt.Printf("conf: %v\n", conf)
	// vp.Save()
	fmt.Printf("vp.Get(): %v\n", vp.Get())

	select {}

}
