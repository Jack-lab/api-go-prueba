# api-go-prueba
CRUD in Golang to create, request, update and delete a note. With sqlite3 database connection

### Build project
to build the project. it must be placed in the root folder of the project /api-go-test. and execute the command:
`go build`
Once compiled we run it with the -migrate flag so that the notes table is created: 
`./go-web-server -migrate`
With this we should see in the terminal the message "Running at http://localhost:8080".

To make requests to our API I will use cURL, feel free to use the Rest client of your preference such as Postman where you add 
a collection of requests in the repository if you want to use postman. While the server is running, from another terminal we proceed to make the requests.

### Creating a note (POST)
`curl -X POST http://localhost:8080/notes -H 'Content-Type: application/json' -d '{"title": "First test note", "description": "This is a test note..."}'`

### Getting notes (GET)
`curl -X GET http://localhost:8080/notes`

### Editing a note (PUT)
`curl -X PUT http://localhost:8080/notes -H 'Content-Type: application/json' -d '{"id": 1, "title": "First edited note", "description": "Edited test note"}'`

### Deleting a note (DELETE)
`curl -X DELETE http://localhost:8080/notes?id=1`

