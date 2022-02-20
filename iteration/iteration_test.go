package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := UpgradeRepeat("a", 5)
	expected := "5...4...3...2...1...a"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleUpgradeRepeat() {
	sum := UpgradeRepeat("Baam!", 5)
	fmt.Println(sum)
	// Output: 5...4...3...2...1...Baam!
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UpgradeRepeat("a", 100)
	}
}
