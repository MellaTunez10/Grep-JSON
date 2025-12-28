package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const FILENAME = "tasks.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo [add|list|update|delete]")
		return
	}

	tasks := loadTasks()

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo add [title]")
			return
		}
		tasks = append(tasks, map[string]any{
			"title":    os.Args[2],
			"priority": 3.0,
			"is_done":  false,
		})
		saveTasks(tasks)
		fmt.Println("Task added")

	case "list":
		fileInfo, _ := os.Stdout.Stat()

		// Check if the output is NOT a terminal (meaning it's a pipe or file)
		if (fileInfo.Mode() & os.ModeCharDevice) == 0 {
			// Output raw JSON for gfilter
			data, _ := os.ReadFile(FILENAME)
			fmt.Print(string(data))
		} else {
			// Output the pretty table for the human
			printHumanList(tasks)
		}

	case "update":
		if len(os.Args) < 5 {
			fmt.Println("Usage: todo update [id] [key] [value]")
			return
		}

		// Convert User Input (e.g., "1") to Internal Index (0)
		userID, _ := strconv.Atoi(os.Args[2])
		id := userID - 1

		// Safety check (Bounds)
		if id < 0 || id >= len(tasks) {
			fmt.Println("Error: Task ID not found")
			return
		}
		key := os.Args[3]
		valStr := os.Args[4]

		// Smart Type Detection
		if b, err := strconv.ParseBool(valStr); err == nil {
			tasks[id][key] = b
		} else if f, err := strconv.ParseFloat(valStr, 64); err == nil {
			tasks[id][key] = f
		} else {
			tasks[id][key] = valStr
		}

		saveTasks(tasks)
		fmt.Println("Task updated")

	case "delete":
		// Convert string ID to int
		userID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: ID must be a number")
			return
		}
		id := userID - 1

		// Safety check: is the ID valid?
		if id < 0 || id >= len(tasks) {
			fmt.Println("Error: No task found with that ID")
			return
		}

		// Confirmation step (Option B)
		fmt.Printf("delete '%s'? (y/n): ", tasks[id]["title"])
		var res string
		fmt.Scanln(&res)

		if res == "y" || res == "Y" {
			// The 'delete' trick: tasks[:id] + tasks[id+1:]
			tasks = append(tasks[:id], tasks[id+1:]...)

			// Save the updated list
			saveTasks(tasks)
			fmt.Println("Task deleted!")
		} else {
			fmt.Println("Deletion cancelled.")
		}
	}
}

func loadTasks() []map[string]any {
	data, err := os.ReadFile(FILENAME)
	if err != nil {
		return []map[string]any{}
	}

	var tasks []map[string]any
	json.Unmarshal(data, &tasks)
	return tasks
}

func saveTasks(t []map[string]any) {
	data, _ := json.MarshalIndent(t, "", " ")
	os.WriteFile(FILENAME, data, 0o644)
}

func printHumanList(tasks []map[string]any) {
	fmt.Println("ID | Done | Priority | title")
	for i, t := range tasks {
		fmt.Printf("%-2d | %-4v | %-4v | %s\n", i+1, t["is_done"], t["priority"], t["title"])
	}
}
