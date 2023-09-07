package loop

import "testing"

func TestLoopA(t *testing.T) {
	loop := LoopA("a")
	expected := "aaaaaa"

	if loop != expected {
		t.Errorf("expected `%s`, returned `%s`", expected, loop)
	}
}

func BenchmarkLoopA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopA("A")
	}
}
