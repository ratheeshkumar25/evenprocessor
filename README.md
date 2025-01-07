# Event Processor

## Overview

Event Processor is a Golang-based service that listens for events, processes them, and performs necessary actions such as storing them in a database or triggering other services. This service is designed to be simple, efficient, and highly configurable.

## Features

- **Event Handling**: Process incoming events from various sources.
- **Configurable**: Easily configurable with `.env` files and supported configuration libraries.
- **REST API**: Provides a simple API for interacting with the service (built using `gorilla/mux`).
- **File Watching**: Capable of monitoring file changes using `fsnotify`.
- **Extensible**: Designed to be extended with custom event types and processing logic.

## Project Directory 
Directory structure:
└── ratheeshkumar25-evenprocessor/
    ├── README.md
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── .gitignore 
    ├── cmd/
    │   ├── main.go
    │   └── .env
    ├── config/
    │   └── config.go
    ├── k8s/
    │   └── even-process.yaml
    ├── logger/
    │   └── logger.go
    └── pkg/
        ├── domain/
        │   ├── event.go
        │   └── interfaces.go
        ├── handlers/
        │   └── event_handler.go
        ├── models/
        │   └── events.go
        ├── repository/
        │   └── event_repository.go
        ├── server/
        │   └── server.go
        ├── usecase/
        │   └── event_usecase.go
        └── worker/
            └── worker.go


## Requirements

- Go 1.23.1 or later
- Docker (for containerization)
- Docker Compose (for running the service in a containerized environment)
- PostgreSQL or any other database for event storage (optional, based on your configuration)

## Setup

### 1. Clone the repository

```bash
git clone https://github.com/ratheeshkumar/event-processor.git
cd event-processor
