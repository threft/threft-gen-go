package gog

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/threft/threft/tidm"
)

func GenerateGo(t *tidm.TIDM) {
	cs := spew.NewDefaultConfig()
	cs.Indent = "    "
	fmt.Println("Dumping TIDM")
	cs.Dump(t)
}
