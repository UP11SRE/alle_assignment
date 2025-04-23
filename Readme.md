# Task Management Service

This is a simple task management microservice built in Go using the Gin framework and PostgreSQL. It allows you to create and retrieve tasks via a REST API.

## Problem Breakdown and Design Decisions

### Problem

- Build a backend service that:

- Accepts and stores task data via HTTP APIs.

- Stores data in PostgreSQL.

- Returns stored tasks on request.

### Design Decisions

- Framework: Chose Gin for its simplicity and performance in building RESTful APIs.

- Database: PostgreSQL for relational storage and reliability.

- MVC architecture for clear separation of concerns

## Instructions to Run the Service

- Clone the repo

```
git clone https://github.com/UP11SRE/alle_assignment.git
```

- Please copy the `.env.example` file to `.env` before running the application:

- Run The Service:

```
go run main.go
```

The server will start at http://localhost:8080

## API Documentation

### POST Create a Task

- http://localhost:8080/tasks

- Request Body(JSON)

```
{
  "title": "Write README",
  "description": "Add instructions to the repo",
  "user_name": "Naman"
}
```

- Task Status By Deafult Will be "Pending"

- Response Body:

```{
  "id": 1,
  "title": "Write README",
  "description": "Add instructions to the repo",
  "status": "pending",
  "user_name": "Naman",
  "created_at": "2025-04-23T12:34:56Z",
  "updated_at": "2025-04-23T12:34:56Z"
}
```

### Get All Tasks

- GET http://localhost:8080/tasks

You can send page in queryparam like: page = 1, By default page will be 1.

- Response:

```
{
    "meta": {
        "current_page": 1,
        "has_next": true,
        "page_size": 10,
        "total_items": 23,
        "total_pages": 3
    },
    "tasks": [
        {
            "id": 1,
            "title": "First Task",
            "description": "This is for testing",
            "status": "completed",
            "user_name": "Naman",
            "created_at": "2025-04-22T10:57:26.346103Z",
            "updated_at": "2025-04-22T13:28:11.750968Z"
        },
        {
            "id": 2,
            "title": "Second Task",
            "description": "Testing",
            "status": "pending",
            "user_name": "Naman1",
            "created_at": "2025-04-22T13:01:34.7092Z",
            "updated_at": "2025-04-22T13:01:34.7092Z"
        },
        {
            "id": 3,
            "title": "Second Task",
            "description": "Testing",
            "status": "pending",
            "user_name": "Naman1",
            "created_at": "2025-04-22T13:11:55.481978Z",
            "updated_at": "2025-04-22T13:11:55.481978Z"
        },
        {
            "id": 4,
            "title": "Second Task",
            "description": "Testing",
            "status": "pending",
            "user_name": "Naman1",
            "created_at": "2025-04-22T13:11:56.705447Z",
            "updated_at": "2025-04-22T13:11:56.705447Z"
        },
        {
            "id": 5,
            "title": "Second Task",
            "description": "Testing",
            "status": "pending",
            "user_name": "Naman1",
            "created_at": "2025-04-22T13:11:57.941042Z",
            "updated_at": "2025-04-22T13:11:57.941042Z"
        },
        {
            "id": 6,
            "title": "Second Task",
            "description": "Testing",
            "status": "pending",
            "user_name": "Naman1",
            "created_at": "2025-04-22T13:11:58.995972Z",
            "updated_at": "2025-04-22T13:11:58.995972Z"
        },
        {
            "id": 7,
            "title": "Second Task",
            "description": "Testing",
            "status": "pending",
            "user_name": "Naman1",
            "created_at": "2025-04-22T13:11:59.973893Z",
            "updated_at": "2025-04-22T13:11:59.973893Z"
        },
        {
            "id": 8,
            "title": "Second Task",
            "description": "Testing",
            "status": "pending",
            "user_name": "Naman1",
            "created_at": "2025-04-22T13:12:00.913421Z",
            "updated_at": "2025-04-22T13:12:00.913421Z"
        },
        {
            "id": 9,
            "title": "Second Task",
            "description": "Testing",
            "status": "pending",
            "user_name": "Naman1",
            "created_at": "2025-04-22T13:12:03.271144Z",
            "updated_at": "2025-04-22T13:12:03.271144Z"
        },
        {
            "id": 10,
            "title": "Second Task",
            "description": "Testing",
            "status": "pending",
            "user_name": "Naman1",
            "created_at": "2025-04-22T13:12:04.394361Z",
            "updated_at": "2025-04-22T13:12:04.394361Z"
        }
    ]
}
```

### GET Filter Tasks

- localhost:8080/tasks/filter?status=pending

here int he status you can send pending or completed, task either can be pending or completed.

- with status you can send page in queryparam as well

- Response:

```
{
    "meta": {
        "current_page": 1,
        "has_next": true,
        "page_size": 10,
        "total_items": 21,
        "total_pages": 3
    },
    "tasks": [
        {
            "id": 2,
            "title": "Second Task",
            "description": "Testing",
            "status": "pending",
            "user_name": "Naman1",
            "created_at": "2025-04-22T13:01:34.7092Z",
            "updated_at": "2025-04-22T13:01:34.7092Z"
        },
        {
            "id": 3,
            "title": "Second Task",
            "description": "Testing",
            "status": "pending",
            "user_name": "Naman1",
            "created_at": "2025-04-22T13:11:55.481978Z",
            "updated_at": "2025-04-22T13:11:55.481978Z"
        },
        {
            "id": 4,
            "title": "Second Task",
            "description": "Testing",
            "status": "pending",
            "user_name": "Naman1",
            "created_at": "2025-04-22T13:11:56.705447Z",
            "updated_at": "2025-04-22T13:11:56.705447Z"
        },
        {
            "id": 5,
            "title": "Second Task",
            "description": "Testing",
            "status": "pending",
            "user_name": "Naman1",
            "created_at": "2025-04-22T13:11:57.941042Z",
            "updated_at": "2025-04-22T13:11:57.941042Z"
        },
        {
            "id": 6,
            "title": "Second Task",
            "description": "Testing",
            "status": "pending",
            "user_name": "Naman1",
            "created_at": "2025-04-22T13:11:58.995972Z",
            "updated_at": "2025-04-22T13:11:58.995972Z"
        },
        {
            "id": 7,
            "title": "Second Task",
            "description": "Testing",
            "status": "pending",
            "user_name": "Naman1",
            "created_at": "2025-04-22T13:11:59.973893Z",
            "updated_at": "2025-04-22T13:11:59.973893Z"
        },
        {
            "id": 8,
            "title": "Second Task",
            "description": "Testing",
            "status": "pending",
            "user_name": "Naman1",
            "created_at": "2025-04-22T13:12:00.913421Z",
            "updated_at": "2025-04-22T13:12:00.913421Z"
        },
        {
            "id": 9,
            "title": "Second Task",
            "description": "Testing",
            "status": "pending",
            "user_name": "Naman1",
            "created_at": "2025-04-22T13:12:03.271144Z",
            "updated_at": "2025-04-22T13:12:03.271144Z"
        },
        {
            "id": 10,
            "title": "Second Task",
            "description": "Testing",
            "status": "pending",
            "user_name": "Naman1",
            "created_at": "2025-04-22T13:12:04.394361Z",
            "updated_at": "2025-04-22T13:12:04.394361Z"
        },
        {
            "id": 11,
            "title": "Second Task",
            "description": "Testing",
            "status": "pending",
            "user_name": "Naman1",
            "created_at": "2025-04-22T13:12:05.387287Z",
            "updated_at": "2025-04-22T13:12:05.387287Z"
        }
    ]
}
```

### GET Get Task by id

- localhost:8080/tasks/:id

- Response:

```
{
    "id": 24,
    "title": "New Title",
    "description": "Description of the new task",
    "status": "completed",
    "user_name": "Naman1",
    "created_at": "2025-04-23T04:59:56.281556Z",
    "updated_at": "2025-04-23T05:01:07.994987Z"
}
```

### Delete Delete Task

- localhost:8080/tasks/:id

- Response:

```
{
    "message": "Task deleted successfully"
}
```

### PATCH Update Task

- We are using patch here insted of put because here you can update only status, title or description. or all.
- Username can't be updated.
- status either can be pendning or completed only

- localhost:8080/tasks/:id

- Request Body:

```
{
    "title" : "New Title",
    "status" : "completed"
}
```

- Response:

```
{
    "message": "Task updated successfully"
}
```

## Structure

```
├── app
│   ├── controllers
│   │   └── task_controller.go
│   ├── models
│   │   └── task_model.go
│   ├── repositories
│   │   └── task_repository.go
│   ├── services
│   │   └── task_service.go
│   ├── routes
│   │   └── router.go
│   └── main.go
├── config
│   └── database.go
├── go.mod
└── go.sum

```

## Usage of the Structure

- `app/controllers`
  This folder contains the controllers responsible for handling HTTP requests and defining the application's behavior.

- `app/models`
  This folder defines the structure of the data in your app.
  For example, task_model.go contains the Task struct, which represents a task in the database.
  It is also used for validation when creating or updating tasks.

- `app/repositories`
  This layer is responsible for direct interaction with the database using raw SQL.

- `app/services`
  Implements the business logic of your app and acts as a bridge between controllers and repositories.
  Example: `task_service.go` validates data or applies rules before saving it.

- `app/routes`
  This folder contains the route configuration and maps HTTP endpoints to their corresponding controller methods.

- `app/main.go`
  This is the entry point of your application. It initializes the server and sets up the routes.

- `config/database.go`
  This file contains the database connection logic and any initial setup like table creation .

- `go.mod` & `go.sum` These files manage the Go modules and their dependencies.

## Scalability

- To scale the Task Management Service horizontally, we can deploy multiple instances (pods) of the application behind a load balancer. Each instance is stateless and communicates with a centralized PostgreSQL database, ensuring that scaling out does not affect data consistency. Kubernetes or any container orchestration platform can be used to manage these pods and ensure high availability and load distribution. Additionally, caching layers and read replicas of the database can be added for further performance improvement as needed.

## Inter-Service Communication

If we were to introduce another microservice — for example, a User Service — communication between services could happen in two ways:

- `Synchronous (Tightly Coupled)`
  If the Task Service depends on data from the User Service (e.g., to validate or fetch user info during task creation), we would use REST APIs or gRPC for real-time communication.

- `Asynchronous (Loosely Coupled)`
  If the services are independent and only need to notify each other of events (e.g., "Task Created" → notify User Service), then message queues like RabbitMQ, Kafka, or AWS SQS would be used. This allows parallel, fault-tolerant, and decoupled event-driven communication.

## Author:

Naman Kansal
