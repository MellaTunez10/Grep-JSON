# Grep-JSON

A collection of lightweight, composable Go utilities for managing and filtering JSON data.

## Project Structure 📂
- `cmd/todo`: A task manager for creating and listing to-do items.
- `cmd/gfilter`: A "smart" generic filter that processes JSON arrays using map-based key/value matching.

## Installation 🚀
1. Clone the repo: `git clone https://github.com/mellatunez10/Grep-JSON`
2. Build the tools:
   ```bash
   go build -o todo ./cmd/todo
   go build -o gfilter ./cmd/gfilter

To-Do App (todo)
The Data Producer. It manages a tasks.json file and outputs either human-readable tables or raw JSON for machine processing.

Command,Arguments,Description
add, [title], Adds a new task with default priority (3.0) and status (false).

./todo list, none, Smart Output: Shows a table if in terminal; outputs JSON if piped.
update ,[id] [key] [value], Updates a field. Supports 1-based IDs and smart type detection.
delete, [id], Removes a task after a safety confirmation (y/n).

GFilter App (gfilter)
The Data Consumer. It reads JSON from standard input, filters it based on logic, and outputs the result as JSON.

Flag,Example,Description
--key,priority,The JSON key you want to inspect.
--op,>,"The operator: >, <, or ==."
--value,2,The value to compare against (numeric or string).
