package hello_world

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Cuco", "")
		want := "Oh! Qué pasó Cuco?!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("defaults to 'mi niño'", func(t *testing.T) {
		got := Hello("", "")
		want := "Oh! Qué pasó mi niño?!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in English", func(t *testing.T) {
		got := Hello("Elodie", "English")
		want := "Hello, Elodie!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Antonio", "French")
		want := "Bonjour, Antonio!"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
