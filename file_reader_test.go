package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileReader(t *testing.T) {
	path := "./mock/hello.txt"
	var reader FileReader
	content, err := reader.Read(path)
	got := string(content)
	expected := "Hello, World!"
	if err != nil {
		log.Fatalln(err)
	}
	// if got != expected {
	// 	t.Errorf("expected %s got %s", expected, got)
	// }
	assert.Equal(t, got, expected)

}

func TestFileNotExist(t *testing.T) {
	path := "./mock/hello2.txt"
	var reader FileReader
	_, err := reader.Read(path)
	if err == nil {
		t.Error("File doesn't exist fails")
	}
}
