package main

import (
	"fmt"
	"github.com/taylormadore/cachito-gomod-local-parent-deps/baz-package"
)


func main() {
	fmt.Println("Hello, local dependencies.")
	baz.Hi()
}
