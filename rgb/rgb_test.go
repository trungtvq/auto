package rgb

import (
	"fmt"
	"testing"
)

func TestRGB(t *testing.T){
	//34.178 de4539 rgb(222,69,57)
	//35.217 1e2929	rgb(30,41,41)
	//35.254 22292b rgb(34,41,43)
	//36.293 d83331 rgb(216,51,49)
	var h = Hex("de4539")
	fmt.Print(h.toRGB())
}
