package solution

import (
	"fmt"
	"testing"

	"github.com/kyokomi/emoji/v2"
)

func TestHello(t *testing.T) {
	world := emoji.Sprint(":world_map:!")
	got := fmt.Sprintln("Hello", world)
	if got != fmt.Sprintln("Hello", world) {
		t.Errorf("CheckLen return %v; want true", got)
	}
}
