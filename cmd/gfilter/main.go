package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	// 1. Define Flags
	key := flag.String("key", "", "The JSON key to filter by")
	op := flag.String("op", "==", "Comparison operator (>, <, ==)")
	value := flag.String("value", "", "The value to compare against")
	flag.Parse()

	if *key == "" || *value == "" {
		fmt.Fprintln(os.Stderr, "Usage: ./gfilter --key=priority --op='<' --value=3")
		return
	}

	// 2. Read from Stdin (The Pipe)
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		return
	}

	var data []map[string]any
	if err := json.Unmarshal(input, &data); err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing JSON:", err)
		return
	}

	// 3. Filter the Data
	var filtered []map[string]any
	for _, item := range data {
		val, exists := item[*key]
		if !exists {
			fmt.Fprintf(os.Stderr, "Warning: Key '%s' missing in an item\n", *key)
			continue
		}

		if compare(val, *op, *value) {
			filtered = append(filtered, item)
		}
	}

	// 4. Output the Result
	output, _ := json.MarshalIndent(filtered, "", "  ")
	fmt.Println(string(output))
}

func compare(jsonVal any, op string, userValStr string) bool {
	// Try numeric comparison first
	jsonNum, err1 := strconv.ParseFloat(fmt.Sprint(jsonVal), 64)
	userNum, err2 := strconv.ParseFloat(userValStr, 64)

	if err1 == nil && err2 == nil {
		switch op {
		case ">":
			return jsonNum > userNum
		case "<":
			return jsonNum < userNum
		case "==":
			return jsonNum == userNum
		}
	}

	// Fallback to string comparison
	jsonStr := fmt.Sprint(jsonVal)
	switch op {
	case "==":
		return jsonStr == userValStr
	default:
		return false
	}
}
