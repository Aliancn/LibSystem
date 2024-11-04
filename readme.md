# Project Name

## Description
This project is a Go application that uses Redis and PostgreSQL as its data stores. The application is containerized using Docker and managed with Docker Compose.

## Prerequisites
- Go 1.23.2
- Docker
- Docker Compose

## Setup

1. Clone the repository:
    ```sh
    git clone <repository-url>
    cd <repository-directory>
    ```

2. Start the Docker containers:
    ```sh
    ./script/run.sh start
    ```

3. Stop the Docker containers:
    ```sh
    ./script/run.sh stop
    ```

## Directory Structure
```
.
├── common
│   ├── enum.go
│   └── utils
│       ├── encrypt.go
│       ├── generice.go
│       ├── jwt.go
│       └── oss.go
├── config
│   ├── application-dev.yaml
│   ├── application-release.yaml
│   └── config.go
├── data
│   ├── postgres
│   │   ├── config
│   │   ├── data
│   │   └── logs
│   ├── redis
│   │   ├── conf
│   │   ├── data
│   │   └── logs
│   ├── running
│   │   ├── config
│   │   └── logs
│   └── uploads
├── docker-compose.yaml
├── dockerfile
├── global
│   ├── global.go
│   └── tx
│       ├── gormTx.go
│       └── transactionManager.go
├── go.mod
├── go.sum
├── initialize
│   ├── enter.go
│   ├── gorm.go
│   ├── redis.go
│   └── router.go
├── internal
│   ├── api
│   │   ├── controller
│   │   ├── request
│   │   └── response
│   ├── model
│   │   ├── Book.go
│   │   ├── Borrow.go
│   │   ├── Paper.go
│   │   ├── User.go
│   │   ├── file.go
│   │   └── model.go
│   ├── repository
│   │   ├── book_repo.go
│   │   ├── borrow_repo.go
│   │   ├── dao
│   │   ├── paper_repo.go
│   │   └── user_repo.go
│   ├── router
│   │   ├── admin
│   │   ├── common
│   │   ├── router.go
│   │   └── user
│   └── service
│       ├── book.go
│       ├── borrow.go
│       ├── paper.go
│       └── user.go
├── logger
│   ├── log.go
│   ├── slog.go
│   └── systemLog.txt
├── main.go
├── middle
│   └── jwt_middle.go
├── readme.md
├── script
│   ├── db.sql
│   ├── run.bash
│   └── run.sh
└── uploads
    └── papers
```



## License
This project is licensed under the MIT License.

