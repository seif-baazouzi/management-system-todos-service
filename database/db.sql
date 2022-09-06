CREATE DATABASE todos;

\c todos

CREATE TABLE todos (
    todoID UUID PRIMARY KEY,
    title VARCHAR NOT NULL, 
    body VARCHAR NOT NULL, 
    userID UUID NOT NULL,
    workspaceID UUID NOT NULL,
    createdAt TIMESTAMP DEFAULT NOW() 
);

CREATE INDEX UserIdInex ON todos USING hash (userID);
CREATE INDEX WorkspaceInex ON todos USING hash (workspace);
