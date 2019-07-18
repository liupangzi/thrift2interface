// +build linux darwin
// +build go1.12

package main

import (
	"github.com/liupangzi/thrift2interface/cmd"
)

func main() {
	_ = cmd.RootCmd.Execute()
}
