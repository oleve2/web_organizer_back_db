# Web organizer project (backend and DB)
Disclaimer!  
That project was designed to be a private knowledge base and activity tracker, used by mostly a single person. That's why database and authentication methods were coded the way they are)  
Multiple users refeactor is in production.

## Installation
- In order to build a binary for your OS you need to have golang version >= 1.17.3 installed.
- Make and ``.env``` file in the root of the project (see ```.env file``` topic for details).
- To run server you can use command ```go run ./cmd/main.go``` from root project folder, or make a binary file.

### Binary compilation
Example commands for amd64 architecture for windows, linux, macos  
```
# windows
GOOS=windows GOARCH=amd64 go build -o wa3_back .\cmd\main.go
#maocs
GOOS=darwin GOARCH=amd64 go build -o wa3_back ./cmd/main.go
#linux
GOOS=linux GOARCH=amd64 go build -o wa3_back ./cmd/main.go
```

## Prepare database
For making and sqlite .db file we use python >= 3.7.6.  
You can inspect and interactively run notebook 'WA3 DB support (init and modify).ipynb'.    
You can name a database file however you want - make sure you set a correct path to it in a ```.env``` file.  

## .env file
In order to connect database and specify address/port - you need to create an ```.env``` file with the following structure. Note that keys names should not be changed.  
Make sure host/port correspond to frontend env variable.

#### .env file content
```
APP_HOST="0.0.0.0"
APP_PORT="9999"
APP_SQLITE_PATH="./db/prd.db"
```
APP_HOST - your host address
APP_PORT - port
APP_SQLITE_PATH - path to sqlite database

