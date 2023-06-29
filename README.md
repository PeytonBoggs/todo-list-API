# todo-list-API

## Setup

### Run the server:
 
 From root,

    `go run .`

This command requires local .env file in this format:

```
DBUSER = "<username>"
DBPASS = "<password>"
NET = "tcp"
ADDR = "127.0.0.1:3306"
DBNAME = "tasks"
```

For an example, see .env.example

### Login to SQL:

From root in a separate terminal,

    mysql -u root -p
    Password: pass

Set database:

    use tasks;

### Initialize Swagger

From root in a separate terminal,

    swag init

Then, as server is running, go to http://localhost:8080/swagger/index.html

## Terminal commands:

For all commands, 

`<ID>` is an `int` value

`<Title>` is a `string` value

`<Complete>` is a `boolen` value


### Get health:
    curl http://localhost:8080/health
    
### Get all tasks:
    curl http://localhost:8080/tasks

### Get task by ID:
    curl http://localhost:8080/tasks/<ID>

### Post task:
    curl http://localhost:8080/tasks \
        --include \
        --header "Content-Type: application/json" \
        --request "POST" \
        --data '{"ID": <ID>,"Title": "<Title>", "Complete": <Complete>}'

### Put task at ID:
    curl http://localhost:8080/tasks/<ID> \
        --include \
        --header "Content-Type: application/json" \
        --request "PUT" \
        --data '{"ID": <ID>,"Title": "<Title>", "Complete": <Complete>}'

### Delete task by ID:
    curl http://localhost:8080/tasks/<ID> \
        --include \
        --header "Content-Type: application/json" \
        --request "DELETE"

### Get table:

In SQL logged-in terminal, having set the database,

    SELECT * FROM tasks;