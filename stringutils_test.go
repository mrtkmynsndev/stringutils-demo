package stringutils_test

import (
	"fmt"
	"testing"

	"github.com/mrtkmynsndev/stringutils-demo"
)

func TestReverse(t *testing.T) {
	inputs := map[string]struct {
		input string
		want  string
	}{
		"with Turkish Letter": {input: "Kimyonşen", want: "neşnoymiK"},
		"none Turkish Letter": {input: "Hello World", want: "dlroW olleH"},
	}

	for k, v := range inputs {
		t.Run(k, func(t *testing.T) {
			got := stringutils.Reverse(v.input)
			if got != v.want {
				t.Errorf("Failed. want -> %s, got -> %s", v.want, got)
			}
		})
	}
}

var gs string

func BenchmarkReverse(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = stringutils.Reverse("hello world!")
	}
	gs = s
}

func ExampleReverse() {
	fmt.Println(stringutils.Reverse("mert"))
	// Output: trem
}
