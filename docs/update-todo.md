# Create Todo

Used to update an existing todo.

**URL**: `/api/v1/todos/:todoID`

**Method**: `PUT`

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
    "message": "success"
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
