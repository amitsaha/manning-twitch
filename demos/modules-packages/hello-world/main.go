package main

import (
	"fmt"
	"io"
	"os"

	"github.com/amitsaha/manning-twitch/modules-packages/hello-world/hello"
	"github.com/amitsaha/manning-twitch/modules-packages/hello-world/world"
)

func displayGreetings(w io.Writer) {
	fmt.Fprintln(w, hello.Greet())
	fmt.Fprintln(w, world.Greet())
}

func main() {
	displayGreetings(os.Stdout)
}
