// This is a placeholder file to appease the linter and will be overwritten during the build process by the generated files
package main

import (
	"net/http"
)

type Placeholder struct {
}

func (p *Placeholder) Open(name string) (http.File, error) {
	return nil, nil
}

var FS = &Placeholder{}
