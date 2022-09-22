# Delete Workspace Todos

Used to delete workspace todos.

**URL**: `/api/v1/todos/workspace/:workspaceID`

**Method**: `DELETE`

**Auth required**: Send User `JWT token` in `X-Token` header

### Success Response

**CODE**: `200`

```json
{
    "message": "success"
}
```

### Server Error Response

**CODE**: `500`

```json
{
    "message": "server-error"
}
```
