# taskMangementService

A simple, extensible task management system built with Go using microservice architecture principles.  
It allows users to Create, Read, Update, Delete tasks** with additional features like **pagination, filtering**, and user validation via inter-service communication.

## Problem Breakdown

### Goal

Build a task service that supports:
- CRUD operations on tasks
- Filtering by status (e.g., completed, pending)
- Pagination
- Clean code structure following microservice principles
- External service communication (e.g., User validation)
- Asynchronous event handling (e.g., task notification via Kafka


## Design Decisions

| Concern               | Design Choice                                                        |
|-----------------------|----------------------------------------------------------------------|
| Language              | Go                                  |
| Framework             | [Gin](https://github.com/gin-gonic/gin) – HTTP router                |
| ORM                   | [GORM](https://gorm.io) – database handling                          |
| DB                    | SQLite                        |
| User Validation       | External REST call to User Service (simulated)                                   |
| Notifications         | Kafka used for async email/SMS trigger (simulated)                   |
| Project Structure     | Follows Separation of Concerns principle                         |
| Dependency Management | Go Modules (`go mod`)                                                |

## Instructions to run the service

Command - Go run main.go  (It will run on localhost 8081)
Install Dependencies - go mod tidy

### Api Documentation

Base URL - http://localhost:8081/tasks

1) Create task - 

Request -  

curl --location 'http://localhost:8081/tasks' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Task 2",
    "description" : "Learn SQS",
    "userId" : 1433,
    "status": "pending"
}'

Response - status (201)

{
    "id": 4,
    "title": "Task 2",
    "description": "Learn SQS",
    "status": "pending",
    "userId": 1433,
    "created_at": "2025-04-25T16:44:15.31241+05:30",
    "updated_at": "2025-04-25T16:44:15.31241+05:30"
}

2) Get all tasks

Query Parameters (optional):
  page: integer (e.g., 1)
  limit: integer (e.g., 5)
  status: string (e.g., completed)

Request - 
  
  curl --location 'http://localhost:8081/tasks' \
--header 'Content-Type: application/json' \
--data ''

Response - 

  [
    {
        "id": 1,
        "title": "Task 8",
        "description": "Learn microservice",
        "status": "completed",
        "userId": 123,
        "created_at": "2025-04-25T16:28:50.595776+05:30",
        "updated_at": "2025-04-25T16:28:50.595776+05:30"
    },
    {
        "id": 2,
        "title": "Task 3",
        "description": "Learn microservice",
        "status": "completed",
        "userId": 1433,
        "created_at": "2025-04-25T16:29:27.151034+05:30",
        "updated_at": "2025-04-25T16:29:27.151034+05:30"
    }
]

3) Get Task by ID 

Request - 
  curl --location 'http://localhost:8081/tasks/1' \
--header 'Content-Type: application/json' \
--data ''

Response - 
  {
    "id": 1,
    "title": "Task 8",
    "description": "Learn microservice",
    "status": "completed",
    "userId": 123,
    "created_at": "2025-04-25T16:28:50.595776+05:30",
    "updated_at": "2025-04-25T16:28:50.595776+05:30"
}

4) update task

Request - 
curl --location --request PUT 'http://localhost:8081/tasks/4' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Task 2",
    "description" : "Learn SQS",
    "userId" : 1433,
    "status": "completed"
}'

Response - 

{
    "id": 4,
    "title": "Task 2",
    "description": "Learn SQS",
    "status": "completed",
    "userId": 1433,
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "2025-04-25T16:50:04.393835+05:30"
}

5) Delete task

Request - 
curl --location --request DELETE 'http://localhost:8081/tasks/4' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Task 2",
    "description" : "Learn SQS",
    "userId" : 1433,
    "status": "completed"
}'

Response - 

{
    "message": "Deleted the given task"
}


### Microservice Demonstration

1) Call another microservice (user service) through the REST API.
2) Push event notification to Kafka to send event details via email and SMS.





