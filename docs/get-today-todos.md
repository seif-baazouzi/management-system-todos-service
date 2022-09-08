# Get Today Todos

Used to get user today todos from all his workspaces.

**URL**: `/api/v1/todos/today`

**Method**: `GET`

**Auth required**: Send User `JWT token` in `X-Token` header

### Success Response

**CODE**: `200`

```json
{
    "todos": [
        {
            "todoID": "todoID", 
            "title": "title", 
            "body": "body", 
            "done": true or false, 
            "startingDate": "starting date", 
            "endingDate": "ending date", 
            "userID": "userID", 
            "workspaceID": "workspaceID", 
            "createdAt": "created date" 
        },
        ...
    ]
}
```

### Server Error Response

**CODE**: `500`

```json
{
    "message": "server-error"
}
```
