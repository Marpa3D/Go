// Два способа тестирования
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Mark")
	want := "Hello Mark"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	t.Run("Вместо Hello Mark передана пустая строка", func(t *testing.T) {
		got := Hello("")
		want := "Hello Mark"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
