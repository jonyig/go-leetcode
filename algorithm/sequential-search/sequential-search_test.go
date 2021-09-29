package algorithm

import (
	"testing"
)

func TestSequentialSearch(t *testing.T) {

	list := []string{"皮卡丘", "喵喵", "大岩蛇", "綠毛蟲"}

	target := "綠毛蟲"
	result := SequentialSearch(list, target)

	expect := 3

	if result != expect {
		t.Fatalf("expect is %s in %d, but faitl", target, expect)
	}
}
