
# Log-ingestor

The Log Ingestor is a simple Go application that demonstrates log handling and querying across PostgreSQL and MongoDB databases. It provides HTTP endpoints for logging


## Tech Stack


**Server:** Golang, PostgreSQL, MongoDB


## Prerequisites

Before running the Log Ingestor, ensure you have the following dependencies installed:

1. Go (Golang)
2. PostgreSQL
- Linux
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```
- macOS
```bash
brew install postgresql
```
- Windows

    Download and install PostgreSQL from the    official website: [PostgreSQL Downloads](https://www.postgresql.org/download/)

3. MongoDB

- Linux
```bash
sudo apt update
sudo apt install mongodb
```
- macOS
```bash
brew tap mongodb/brew
brew install mongodb-community
```
- Windows

    Download and install MongoDB from the official website: [MongoDB Downloads](https://www.mongodb.com/try/download/community)
## Installation

1. Clone the repository

```bash
git clone https://github.com/dyte-submissions/november-2023-hiring-swag2716.git
cd log-ingestor    
```
2. Install necessary Go packages
```bash
go get -u github.com/gorilla/mux
go get -u github.com/lib/pq
go get -u go.mongodb.org/mongo-driver/mongo    
``` 
3. Set up your environment variables:

 Create a .env file in the project root and configure the following variables for postgreSQL connection string:
 ```bash
 HOST=<your-host-name>
PORT=5432
USER=<your-user-name>
PASSWORD=<your-postgreSQL-password>
DBNAME=<your-dbname>
 ```
4. Run the application
```bash
go run main.go
```
## API Reference

### Store like event

**Description**: Store a log in both postregreSQL and mongoDB.


**Endpoint**: POST `http//localhost:3000/ingest`

**Request Body**:

```json
{
	"level": "error",
	"message": "Failed to connect to DB",
    "resourceId": "server-1234",
	"timestamp": "2023-09-15T08:00:00Z",
	"traceId": "abc-xyz-123",
    "spanId": "span-456",
    "commit": "5e5342f",
    "metadata": {
        "parentResourceId": "server-0987"
    }
}
```

**Response**:

- Status Code: 200 OK
- Response Body: `"Log ingested successfully!"`
## Documentation

[Documentation](https://linktodocumentation)

