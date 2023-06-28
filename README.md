Terminal commands:

Go:
    Run:
        export DBUSER='root'
        export DBPASS='(password)'
        go run .

Curl:
    Get:
        curl http://localhost:8080/tasks

    Get by ID:
        curl http://localhost:8080/tasks/2

    Post:
        curl http://localhost:8080/tasks \
            --include \
            --header "Content-Type: application/json" \
            --request "POST" \
            --data '{"ID": "4", "Title": "Go to sleep", "Complete": "false"}'

    Put:
        curl http://localhost:8080/tasks/2 \
            --include \
            --header "Content-Type: application/json" \
            --request "PUT" \
            --data '{"ID": "2","Title": "Stay at home", "Complete": "false"}'

    Delete:
        curl http://localhost:8080/tasks/2 \
            --include \
            --header "Content-Type: application/json" \
            --request "DELETE"

SQL:
    Login to SQL:
        mysql -u root -p

    Set database:
        use tasks;

    Get table:
        SELECT * FROM tasks;

Swagger:
    Init:
        swag init

    URL:
        http://localhost:8080/swagger/index.html