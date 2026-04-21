// Package main is a stand-in leaf consumer for the automation playground.
package main

import (
	"fmt"

	"github.com/tomleb/frameworks-automation-steve/pkg/steve"
	"github.com/tomleb/frameworks-automation-wrangler/pkg/wrangler"
)

func main() {
	fmt.Println(wrangler.Hello())
	fmt.Println(steve.Hello())
}
