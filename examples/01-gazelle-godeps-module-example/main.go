package main

import (
	"flag"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	fmt.Println(strfmt.IsUUID("a.b.c"))
	glog.Error("Hello")
}
