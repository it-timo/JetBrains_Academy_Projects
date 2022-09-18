package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type memory struct {
	command string
	notes   []string
}

const (
	Separator = " "
)

var (
	maxValues int
)

func main() {
	var mem memory

	fmt.Println("Enter the maximum number of notes:")
	_, err := fmt.Scanf("%d", &maxValues)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Enter a command and data: ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), Separator)
		mem.command = input[0]

		note := strings.Join(input[1:], Separator)

		fail := handleWrongInput(mem.command, note)

		if fail != "" {
			fmt.Printf(fail)
		} else {
			switch mem.command {
			case "create":
				createNote(&mem, note)
			case "list":
				listMemory(&mem)
			case "clear":
				clearMemory(&mem)
			case "update":
				updateMemory(&mem, note)
			case "delete":
				deleteEntry(&mem, note)
			case "exit":
				fmt.Println("[Info] Bye!")
				return
			default:
				fmt.Println("[Error] Unknown command")
				break
			}
		}

		fmt.Println("Enter a command and data: ")
	}

	if err = scanner.Err(); err != nil {
		log.Println(err)
	}
}

func handleWrongInput(command, note string) string {
	if len(note) < 1 || strings.TrimSpace(note) == "" {
		switch command {
		case "create":
			return "[Error] Missing note argument\n"
		case "update":
			return "[Error] Missing position argument\n"
		case "delete":
			return "[Error] Missing position argument\n"
		}
	}

	tmp := strings.Split(note, Separator)

	if len(tmp) <= 1 {
		switch command {
		case "update":
			return "[Error] Missing note argument\n"
		}
	} else if len(tmp) > 1 && strings.TrimSpace(tmp[1]) == "" {
		switch command {
		case "update":
			return "[Error] Missing note argument\n"
		}
	}

	index, err := strconv.Atoi(tmp[0])
	if err != nil {
		return ""
	}

	if index > maxValues {
		return fmt.Sprintf("[Error] Position %d is out of the boundary [1, %d]\n", index, maxValues)
	}

	return ""
}

func createNote(mem *memory, note string) {
	if len(mem.notes) >= maxValues {
		fmt.Println("[Error] Notepad is full")
	} else {
		fmt.Println("[OK] The note was successfully created")
		mem.notes = append(mem.notes, note)
	}
}

func listMemory(mem *memory) {
	if len(mem.notes) < 1 {
		fmt.Println("[Info] Notepad is empty")
	}
	for i, entry := range mem.notes {
		if entry != "" {
			fmt.Printf("[Info] %d: %s\n", i+1, entry)
		}
	}
}

func clearMemory(mem *memory) {
	fmt.Println("[OK] All notes were successfully deleted")
	mem.notes = []string{}
}

func updateMemory(mem *memory, note string) {
	tmp := strings.Split(note, Separator)
	index, err := strconv.Atoi(tmp[0])
	if err != nil {
		fmt.Printf("[Error] Invalid position: %s\n", note)
		return
	}

	if len(mem.notes) > index-1 {
		mem.notes[index-1] = strings.Join(tmp[1:], Separator)
		fmt.Printf("[OK] The note at position %d was successfully updated\n", index)
	} else {
		fmt.Println("[Error] There is nothing to update")
	}
}

func deleteEntry(mem *memory, note string) {
	index, err := strconv.Atoi(note)
	if err != nil {
		fmt.Printf("[Error] Invalid position: %s\n", note)
		return
	}

	if len(mem.notes) > index {
		mem.notes = append(mem.notes[:index-1], mem.notes[index:]...)
		fmt.Printf("[OK] The note at position %d was successfully deleted\n", index)
	} else {
		fmt.Println("[Error] There is nothing to delete")
	}
}
