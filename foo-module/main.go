package main

import (
	"fmt"
	"github.com/taylormadore/cachito-gomod-local-parent-deps/pkg/baz-package"
	"github.com/taylormadore/cachito-gomod-local-parent-deps/foo-module/foo-package"
)


func main() {
	fmt.Println("Hello, local dependencies.")
	foo.Hi()
	baz.Hi()
}
