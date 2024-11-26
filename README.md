# CLI Todo List App

A simple command-line todo list application built using Go.

## Getting Started
1. **Clone the repository:**
   ```bash
   git clone [https://github.com/splunkdevabhi/todolist-go.git](https://github.com/splunkdevabhi/todolist-go.git)

**Step 2: Install Dependencies**

To install the required dependencies for the project, run the following command in your terminal:

```bash
go mod todolist
```


**Step 3: Run the Application**

To run the application, execute the following command in your terminal:

```bash
go run .
```
**Step 4: Use the Application**

Once the application is running, you can interact with it using the following commands:

* **`list`:** Lists all todos.
* **`add <title>`:** Adds a new todo with the given title.
* **`edit <id> <new_title>`:** Edits the title of a todo with the given ID.
* **`toggle <id>`:** Toggles the completion status of a todo with the given ID.
* **`delete <id>`:** Deletes a todo with the given ID.

**Example Usage:**

```bash
./todolist-go add "Buy groceries"
./todolist-go add "Learn Go"
./todolist-go list
./todolist-go toggle 1
./todolist-go edit 2 "Read a book"
./todolist-go list
./todolist-go delete 2
./todolist-go list
```
## Usage

**Available Commands:**

* **`list`:** Lists all todos.
* **`add <title>`:** Adds a new todo with the given title.
* **`edit <id> <new_title>`:** Edits the title of a todo with the given ID.
* **`toggle <id>`:** Toggles the completion status of a todo with the given ID.
* **`delete <id>`:** Deletes a todo with the given ID.

**Example Usage:**

```bash
./todolist-go add "Buy groceries"
./todolist-go add "Learn Go"
./todolist-go list
./todolist-go toggle 1
./todolist-go edit 2 "Read a book"
./todolist-go list
./todolist-go delete 2
./todolist-go list
```
## Features

* **Persistence:** Todos are stored in a JSON file.
* **User-friendly interface:** Simple and intuitive command-line interface.
* **Basic task management:** Add, edit, delete, and toggle completion status of todos.

## Future Improvements

* **Advanced filtering and searching:** Allow users to filter and search todos based on various criteria.
* **Priority levels:** Add priority levels to todos.
* **Due dates:** Allow users to set due dates for todos.
* **Reminders:** Implement a reminder system to notify users of upcoming deadlines.
* **Integration with calendar apps:** Sync todos with popular calendar apps.

## Contributing

Feel free to contribute to this project by:

1. **Forking the repository.**
2. **Creating a new branch.**
3. **Making your changes.**
4. **Pushing your changes to your fork.**
5. **Submitting a pull request.**

Please follow the existing code style and write clear and concise commit messages.

## License

This project is licensed under the MIT License.
