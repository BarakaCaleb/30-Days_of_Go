## Testing the API

You can test the API endpoints using a tool like [Postman](https://www.postman.com/) or `curl`.

### Example with `curl`

#### Get all books
```bash
`curl http://localhost:8080/books`

Get a single book

`curl http://localhost:8080/book?id=1`

###Create a new book

`curl -X POST -H "Content-Type: application/json" -d '{"id":"4","title":"Brave New World","author":"Aldous Huxley"}' http://localhost:8080/book/create`

###Update a book

`curl -X PUT -H "Content-Type: application/json" -d '{"id":"4","title":"Brave New World Revisited","author":"Aldous Huxley"}' http://localhost:8080/book/update?id=4`

###Delete a book

`curl -X DELETE http://localhost:8080/book/delete?id=4`
