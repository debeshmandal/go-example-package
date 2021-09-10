package illiteracy

import (
	"fmt"
	"testing"

	"github.com/debeshmandal/illiteracy/pkg/illiteracy"
)

func Test_HelloWorld(t *testing.T) {
	illiteracy.HelloWorld()
	fmt.Println("Testing complete")
}

func Test_MessUpString(t *testing.T) {
	s := illiteracy.MessUpString("Hello World!")
	fmt.Println(s)
}
