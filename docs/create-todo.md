# Create Todo

Used to create new todo.

**URL**: `/api/v1/todos/:workspace`

**Method**: `POST`

**Auth required**: Send User `JWT token` in `X-Token` header

### Request

```json
{
    "title": "title", 
    "body": "body", // optional
    "done": true or false, // default is false
    "startingDate": "starting date", 
    "endingDate": "ending date" // optional
}
```

### Success Response

**CODE**: `200`

```json
{
    "message": "success",
    "todoID": "todoID"
}
```

### Invalid Input Response

**CODE**: `200`

```json
{
    "title": "title error message", 
    "startingDate": "starting date error message"
}
```

### Server Error Response

**CODE**: `500`

```json
{
    "message": "server-error"
}
```
