// a package is a collection of common source code files. Every single file must declare the package that it belongs to
// there are two types of packages. Executable (something we can run) and reusable (something used as helpers/libraries)
// we named this package main because any package named main turns into an executable. a package with any other name will instead bo considered a reusable package, aka a helper library/dependency
package main

// fmt is a standard package included with golang by default. fmt is short for "format"
import "fmt"

func main() {
	fmt.Println("Hello World")
}
