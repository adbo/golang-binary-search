# Number Index Service

## Description
Number Index Service is a web application written in Go that allows users to search for the index of a given numeric value in a sorted list of numbers. The application handles HTTP requests and returns the index of the found value or the closest value within a specified tolerance.

## Features
- Loading a sorted list of numbers from a text file.
- Searching for the index of a given numeric value.
- Returning the index of the closest value if the exact value is not found, within a defined tolerance level.

## Getting Started

### Prerequisites
- Go (version 1.15 or later)
- A text file containing sorted numbers (named `input.txt` by default)

### Running the Service
1. To start the service, in the project folder run:
```go run main.go

2. The service will start on the default port (e.g., 8080). You can access the service at `http://localhost:8080`.

### Usage
- Send a GET request to `/endpoint/[number]` where `[number]` is the numeric value you want to search for.
- The service will respond with the index of the number or the closest match within the tolerance.

## Configuration
You can configure the service by editing the `config.yaml` file. Available configurations include:
- `server.port`: Port on which the service will run.
- `logging.level`: Level of logging (e.g., Info, Debug, Error).

## Testing
To run the unit tests, use the following command:
```make test