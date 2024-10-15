package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func clearConsole() {
	// ANSI escape code to clear the screen and move the cursor to the top left
	fmt.Print("\033[H\033[2J")
}

func getTodosFilePath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "todos.txt"), nil
}

func add_todo(todo string, filepath string) {
	//clearConsole()
	if todo == "" {
		fmt.Println("Usage: todo  add [todo item]")
	}
	// Attempt to open the file with write permissions
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Ensure the file is closed properly even if an error occurs
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	// Print the todo item and attempt to write it
	_, err = file.WriteString("\n" + todo)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	//fmt.Println("")
	//list_todos(filepath)
}

func delete_todo(index int, filepath string) {
	//clearConsole()
	rfile, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	lines := strings.Split(string(rfile), "\n")
	//prevLines := lines

	if index < 0 || index >= len(lines) {
		fmt.Printf("line number %d out of range", index)
		return
	}

	// remove line at index from lines array
	if index == 0 && len(lines) == 1 {
		lines = nil
	} else if index == 0 {
		var arrLen = len(lines)
		lines = lines[1:arrLen]
	} else if index == len(lines) {
		var arrLen = len(lines)
		lines = lines[0 : arrLen-1]
	} else {
		var arrLen = len(lines)
		lines = append(lines[0:index], lines[index+1:arrLen]...)
	}

	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("could not recreate file:", err)
		return
	}

	// Ensure the file is closed properly even if an error occurs
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if i == len(lines)-1 {
			_, err := file.WriteString(line)
			if err != nil {
				fmt.Println("could not write file:", err)
				return
			}
		} else {
			_, err := file.WriteString(line + "\n")
			if err != nil {
				fmt.Println("could not write file:", err)
				return
			}
		}
	}

	//for i := 0; i < len(prevLines); i++ {
	//	if i == index {
	//		fmt.Println("(x)", prevLines[i])
	//	} else {
	//		fmt.Println("(", i+1, ")", prevLines[i])
	//	}

	//}
}

func list_todos(filepath string) {
	//clearConsole()
	rfile, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	lines := strings.Split(string(rfile), "\n")

	for i := 0; i < len(lines); i++ {
		fmt.Println("(", i+1, ")", lines[i])
	}
}

func main() {
	for {
		fmt.Print(">")
		reader := bufio.NewReader(os.Stdin)

		request, _ := reader.ReadString('\n')

		split := strings.Split(request, " ")

		if len(split) == 0 {
			fmt.Println("Usage: todo <command> [arguments]")
		}

		split[len(split)-1] = strings.TrimSuffix(split[len(split)-1], "\n")

		command := split[0]

		filepath, err := getTodosFilePath()
		if err != nil {
			fmt.Println("cannot find todo file path")
		}

		switch command {
		case "add":
			add_todo(strings.Join(split[1:], " "), filepath)
			break
		case "-a":
			add_todo(strings.Join(split[1:], " "), filepath)
			break
		case "delete":
			i, err := strconv.Atoi(split[1])
			if err != nil {
				fmt.Println("Usage: todo <command> [arguments]")
				return
			}
			delete_todo(i-1, filepath)
			break
		case "-d":
			i, err := strconv.Atoi(split[1])
			if err != nil {
				fmt.Println("Usage: todo <command> [arguments]")
				return
			}
			delete_todo(i-1, filepath)
			break
		case "list":
			list_todos(filepath)
			break
		case "ls":
			list_todos(filepath)
			break
		case "exit":
			return
		}
	}

}
