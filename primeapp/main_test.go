package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}
func Test_prompt(t *testing.T) {
	oldOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	prompt()

	_ = w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "->") {
		t.Errorf("incorrect prompt: expected -> but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	oldOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	intro()

	_ = w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("intro text not correct; got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	expected := []struct {
		n    string
		msg  string
		quit bool
	}{
		{"q", "", true},
		{"sdhfghs", "Please enter a whole number!", false},
		{"3", "3 is a prime number!", false},
		{"10", "10 is not a prime number because it is divisible by 2!", false},
		{"0", "0 is not prime, by definition!", false},
		{"1", "1 is not prime, by definition!", false},
		{"-23", "Negative numbers are not prime, by definition!", false},
	}

	for _, v := range expected {
		input := strings.NewReader(v.n)
		scanner := bufio.NewScanner(input)
		msg, quit := checkNumbers(scanner)
		if msg != v.msg || quit != v.quit {
			t.Errorf("checkNumbers(%q) = (%q, %t); expected (%q, %t)", v.n, msg, quit, v.msg, v.quit)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)

	stdin := strings.NewReader("1\nq\n")

	go readUserInput(stdin, doneChan)

	<-doneChan

	close(doneChan)
}
