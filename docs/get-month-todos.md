# Get Month Todos

Used to get user todos for specified month.

**URL**: `/api/v1/todos/:workspaceID/month/:month` month is year-month

**Method**: `GET`

**Auth required**: Send User `JWT token` in `X-Token` header

### Success Response

**CODE**: `200`

```json
{
    "todos": {
        "year-month-day": [
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
        ],
        ...
    }
}
```

### Server Error Response

**CODE**: `500`

```json
{
    "message": "server-error"
}
```
