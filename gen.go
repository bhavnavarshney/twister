//+build generate

package main

import (
	"log"
	"os/exec"

	"github.com/zserge/lorca"
)

func main() {
	// You can also run "npm build" or webpack here, or compress assets, or
	// generate manifests, or do other preparations for your assets.
	cmd := exec.Command("npm", "install")
	cmd.Dir = "frontend"
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	cmd = exec.Command("npm", "run", "build")
	cmd.Dir = "frontend"
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	lorca.Embed("main", "assets.go", "frontend/build")
}
