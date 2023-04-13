package main

import (
	"bufio"
	"io"
	"os"
	"reflect"
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

	t.Log("Test IsPrime: function should identify if number is prime or not.")
	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("Test %s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("Test %s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("Test %s: expected %s but got %s", e.name, e.msg, msg)
		}
	}

	t.Log("Test IsPrime: function identifies properly.")
}

func TestIntro(t *testing.T) {
	t.Log("Test Intro: introducing should be just some expressions")

	firstout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	intro()

	w.Close()
	os.Stdout = firstout

	in, _ := io.ReadAll(r)
	got := string(in)

	expected := "Is it Prime?\n------------\nEnter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.\n-> "

	if got != expected {
		t.Errorf("Test Intro: expected \"%s\" but got %s", expected, got)
	}

	t.Log("Test Intro: got as expected.")
}

func TestPrompt(t *testing.T) {
	t.Log("Test Prompt: writes \"-> \" with the space to prompt")
	exp := "-> "

	firstOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	prompt()

	w.Close()
	os.Stdout = firstOut
	in, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("error reading from pipe: %s", err)
	}
	got := string(in)

	if got != exp {
		t.Errorf("TestPrompt: expected %q, but got %q", exp, got)
	}

	t.Log("Test Prompt: got what expected.")
}

func TestCheckNumbers(t *testing.T) {
	tests := []struct {
		input string
		exp   string
	}{
		{"abc", "Please enter a whole number!"},
		{"0", "0 is not prime, by definition!"},
		{"1", "1 is not prime, by definition!"},
		{"3", "3 is a prime number!"},
		{"4", "4 is not a prime number because it is divisible by 2!"},
		{"q", ""},
	}

	t.Log("Test CheckNumbers: function should identify if user inputs number or not.")
	for _, test := range tests {
		scanner := bufio.NewScanner(strings.NewReader(test.input))
		got, res := checkNumbers(scanner)
		if res && test.exp != "" {
			t.Errorf("Test CheckNumbers(%v) returned true, but expected false", test.input)
		}
		if !res && got != test.exp {
			t.Errorf("Test CheckNumbers(%v) = %v, exp %v", test.input, got, test.exp)
		}
	}
}

func TestReadUserInput(t *testing.T) {
	testNumbers := "7\nq\n"
	scanner := bufio.NewReader(strings.NewReader(testNumbers))
	doneChan := make(chan bool)
	t.Log("Test ReadUserInput: should read what user inputs.")

	go readUserInput(scanner, doneChan)

	var output []string

	for res := range doneChan {
		if res {
			return
		}
		line, _, err := scanner.ReadLine()
		if err != io.EOF {
			return
		}
		output = append(output, string(line))
	}

	expected := []string{"7 is a prime number", "Goodbye."}
	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected %v, but got %v", expected, output)
	}
}
