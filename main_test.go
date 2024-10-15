package main

import (
	//	"bytes"
	//	"fmt"
	"os"
	"strings"
	"testing"
)

func TestGetTodosFilePath(t *testing.T) {
	path, err := getTodosFilePath()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !strings.HasSuffix(path, "todos.txt") {
		t.Errorf("Expected file path to end with 'todos.txt', got %s", path)
	}
}

func TestAddTodo(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test_todos.txt")
	if err != nil {
		t.Fatalf("could not create temp file: %v", err)
	}

	defer os.Remove(tmpFile.Name())

	add_todo("Test todo item", tmpFile.Name())

	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("could not read temp file: %v", err)
	}

	if !strings.Contains(string(content), "Test todo item") {
		t.Errorf("expected 'Test todo item' in file content, got %s", content)
	}
	//fmt.Print("TEST PASSED")
}

func TestDeleteTodo(t *testing.T) {
	content := "First\nSecond\nThird"
	tmpFile, err := os.CreateTemp("", "test_delete.txt")
	if err != nil {
		t.Fatalf("Could not create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	_, _ = tmpFile.WriteString(content)
	tmpFile.Close()

	delete_todo(1, tmpFile.Name())

	updatedContent, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Could not read temp file: %v", err)
	}
	if strings.Contains(string(updatedContent), "Second") {
		t.Errorf("Expected 'Second' to be removed, got %s", updatedContent)
	}
}

// func TestListTodos(t *testing.T) {
// 	// Sample content for the todos file
// 	content := "Todo 1\nTodo 2\nTodo 3"
//
// 	// Create a temporary file for testing
// 	tmpFile, err := os.CreateTemp("", "test_list_todos.txt")
// 	if err != nil {
// 		t.Fatalf("Could not create temp file: %v", err)
// 	}
// 	defer os.Remove(tmpFile.Name()) // Clean up the file after the test
//
// 	// Write sample content to the temporary file
// 	_, err = tmpFile.WriteString(content)
// 	if err != nil {
// 		t.Fatalf("Could not write to temp file: %v", err)
// 	}
// 	tmpFile.Close()
//
// 	// Capture the output of list_todos by redirecting os.Stdout
// 	var output bytes.Buffer
// 	originalStdout := os.Stdout
// 	defer func() { os.Stdout = originalStdout }() // Restore original os.Stdout after the test
//
// 	r, w, _ := os.Pipe()
// 	os.Stdout = w
//
// 	// Run list_todos with the temporary file path
// 	list_todos(tmpFile.Name())
//
// 	// Close the writer and copy the output
// 	w.Close()
// 	output.ReadFrom(r)
//
// 	// Expected output format for each line in the todos file
// 	expectedOutput := "(1) Todo 1\n(2) Todo 2\n(3) Todo 3\n"
// 	if strings.TrimSpace(output.String()) != strings.TrimSpace(expectedOutput) {
// 		t.Errorf("Expected:\n%s\nGot:\n%s", expectedOutput, output.String())
// 	}
// }
