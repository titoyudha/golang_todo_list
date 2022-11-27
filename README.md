
# Golang To-Do-List

This is golang todo list API as my assignment.




## API Reference

#### Get all todo

```http
  GET /api/v1/todo
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API path |

#### Post a todo

```http
  POST /api/v1/todo
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `tile , email`      | `string` | **Required**. Your API path |


#### Update a todo

```http
  PUT /api/v1/todo/:id
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `todo_id` | `int` | **Required**. Id of todo to update |

#### Delete todo

```http
  DELETE /api/v1/todo/:id
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `todo_id` | `int` | **Required**. Id of todo to DELETE |





## Run

Step to run this api to your local machine

```bash
  mkdir todo_list
  cd todo_list
  git clone https://github.com/titoyudha/golang_todo_list.git
  change .env files with your configuration
  run go mod tidy
  go test (to run unit testing)
  go run main.go
```
    
## License

[MIT](https://choosealicense.com/licenses/mit/)

