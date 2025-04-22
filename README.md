# 📝 Task Manager Web App (Go + PostgreSQL)

A simple yet powerful task manager web app built using **Go**, **PostgreSQL**, and **Vue.js**.  
Users can add, update, delete, and mark tasks as completed via a responsive web UI.

## Features

- Add, update, delete tasks
- Mark tasks as completed/incomplete
- Clean and interactive UI using Vue.js + Bootstrap
- RESTful API built in Go using Chi router
- PostgreSQL database integration with `pgx`

---

## Tech Stack

| Layer        | Tech                     |
|--------------|--------------------------|
| Backend      | Go (Golang)              |
| HTTP Router  | Chi                      |
| ORM/Driver   | pgx (PostgreSQL driver)  |
| Frontend     | Vue.js + Bootstrap       |
| Renderer     | thedevsaddam/renderer    |
| Database     | PostgreSQL               |

## 📁 Project Structure

```bash
task-manager/
│
├── main.go                    # Entry point of the application
├── go.mod / go.sum            # Module definitions
│
├── static/                    # Frontend assets
│   └── home.tpl               # Main HTML page (Vue.js inside)
│
├── handlers/                 # All route handler logic
│   └── task.go                # Task API handlers (GET, POST, PUT, DELETE)
│
├── models/                   # Data models
│   └── task.go                # Task struct definition
```


## How to Run This App Locally

### 1. Prerequisites

- Go 1.18+ installed
- PostgreSQL installed and running
- Git

---

### 2. Set Up PostgreSQL

1. Create a new database:
```sql
CREATE DATABASE taskdb;
```

2. Inside the `taskdb` database, create the `tasks` table:
```sql
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

### 3. Clone the Project

```bash
git clone https://github.com/your-username/task-manager.git
cd task-manager
```

---

### 4. Configure the Database Connection

In `main.go`, update this line with your actual PostgreSQL password:

```go
const postgresURL = "postgres://postgres:<yourpassword>@localhost:5432/taskdb"
```

Replace `<yourpassword>` with your actual PostgreSQL password.

---

### 5. Install Dependencies

```bash
go mod tidy
```

---

### 6. Run the Application

```bash
go run main.go
```

Then open your browser and go to:

```
http://localhost:9000
```

## 🧪 API Endpoints (for testing with tools like Postman)

| Method | Endpoint     | Description        |
|--------|--------------|--------------------|
| GET    | `/task`      | Fetch all tasks    |
| POST   | `/task`      | Create a new task  |
| PUT    | `/task/{id}` | Update a task      |
| DELETE | `/task/{id}` | Delete a task      |
