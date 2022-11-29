package main

import (
	"fmt"
	"github.com/taylormadore/cachito-gomod-local-parent-deps/pkg/baz-package"
	"github.com/taylormadore/cachito-gomod-local-parent-deps/bar-module/bar-package"
)


func main() {
	fmt.Println("Hello, local dependencies.")
	bar.Hi()
	baz.Hi()
}
