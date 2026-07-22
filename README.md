# TODO-APP

A simple command-line task manager written in Go that lets you add, view, complete, and delete tasks during a session.

## Authors

- Yusuf Zainab

## Usage

### How to run

```bash
go run . todo-app
```

### Available options

Once the app starts, you'll see a menu with 5 options:

1. **View tasks** — displays all your tasks in a table with their ID, description, and status
2. **Add a task** — prompts you to type a new task and saves it to your list
3. **Mark as completed** — enter a task ID to mark it as COMPLETED
4. **Delete a task** — enter a task ID to remove it from the list
5. **Exit** — closes the app

### Example session

```
Welcome To TODO-APP!
Do you want to?
1. View your recorded tasks and their status.
2. Add a new task
3. Mark task as completed
4. Delete a task
5. Exit

2
Note down a task
Buy groceries
Your task has been added!

1
ID   TASK             STATUS
--   ----             ------
1    Buy groceries    PENDING
```

## Implementation details

The app runs a loop that displays the menu and waits for user input. Tasks are stored in a slice of `Todo` structs in memory — data is not saved between sessions.

Each `Todo` has:
- `ID` — unique identifier assigned when task is created
- `Task` — the task description
- `Status` — either `PENDING` or `COMPLETED`

A `tabwriter` is used to format the task list into aligned columns in the terminal.