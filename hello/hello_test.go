package main

import "testing"

func TestHello(t *testing.T) {
	verifyCorrectMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	}

	t.Run("say hello Alessandra", func(t *testing.T) {
		result := Hello("Alessandra", "")
		expected := "Hello, Alessandra"
		verifyCorrectMessage(t, result, expected)
	})

	t.Run("spanish", func(t *testing.T) {
		result := Hello("Alessandra", "spanish")
		expected := "Hola, Alessandra"
		verifyCorrectMessage(t, result, expected)
	})

	t.Run("french", func(t *testing.T) {
		result := Hello("Alessandra", "french")
		expected := "Bonjour, Alessandra"
		verifyCorrectMessage(t, result, expected)
	})

	t.Run("say hello world when name is empty", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello, world"
		verifyCorrectMessage(t, result, expected)
	})
}
