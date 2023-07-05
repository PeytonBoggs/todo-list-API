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

### Swagger:

As server is running, go to http://localhost:8080/swagger/index.html

To update swag docs (in docs folder),

    swag init

### SQL:

From root in a separate terminal,

    mysql -u root -p
    Password: pass

Set database:

    use tasks;

Then run:

    SELECT * FROM tasks;

## Terminal commands:

For all commands, 

`{id}` is an `int` value

`{title}` is a `string` value

`{complete}` is a `boolean` value


### Get health:
    curl http://localhost:8080/health
    
### Get all tasks:
    curl http://localhost:8080/tasks

### Get task by ID:
    curl http://localhost:8080/tasks/id/{id}

### Get task by Complete:
    curl http://localhost:8080/tasks/complete/{complete}

### Post task:
    curl http://localhost:8080/tasks \
        --include \
        --header "Content-Type: application/json" \
        --request "POST" \
        --data '{"Title": "{title}", "Complete": {complete}}'

### Put task at ID:
    curl http://localhost:8080/tasks/id/{id} \
        --include \
        --header "Content-Type: application/json" \
        --request "PUT" \
        --data '{"ID": {id}, "Title": "{title}", "Complete": {complete}}'

### Delete task by ID:
    curl http://localhost:8080/tasks/id/{id} \
        --include \
        --header "Content-Type: application/json" \
        --request "DELETE"

### Patch complete by ID:
    curl http://localhost:8080/tasks/{id} \
        --include \
        --header "Content-Type: application/json" \
        --request "PATCH"

### Get table:

In SQL logged-in terminal, having set the database,

    SELECT * FROM tasks;