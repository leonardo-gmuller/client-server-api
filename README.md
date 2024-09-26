# Client-Server API

## Overview

This project implements a simple client-server API built using Go. The API provides an endpoint to retrieve the exchange rate of USD to BRL. The server interacts with a SQLite database, and the client can request data from the server.

## Built With:
- **Go**: The core programming language used for the API and client-server logic.
- **SQLite**: Lightweight database for storing exchange rate data.

## Folder Structure

📂data<br>
📂client<br>
    ┗ 📜client.go<br>
    📂server<br>
 ┣ 📂database            
 ┣ 📂dto                
 ┣ 📂entity            
 ┣ 📂handler            
 ┗ 📜server.go          



 
`client`: Client to interact with the server

`database`: Database-related files (repositories, migrations, connections)

`dto`: Data Transfer Objects (for data exchange)

`entity`: Database entities/models

`handler`: API handlers (business logic and route handlers)


## Endpoints
- `GET /cotacao`

**Description**: Fetch the current exchange rate of USD to BRL (Brazilian Real).

**Response**: JSON containing the exchange rate.

## Run Locally
### Prerequisites
1. Install Go on your machine.
2. Set up a SQLite database and place the .db file into the data folder.

### Steps
1. Open the project in your editor (e.g., VSCode).
2. Modify the database path if necessary:
    - In the file server/database/db.go, change the database path to match your setup if needed.
3. Run the server:

    Use Go to run the server:
```bash
go run server/server.go
```
4. Run the client:

    Use Go to run the client:
```bash
go run client/client.go
```
Alternatively, if using VSCode, you can run and debug both files using its integrated debug environment.




