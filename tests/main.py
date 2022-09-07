from testRestApi import runTests

import routes.getTodayTodos as getTodayTodos
import routes.createTodo as createTodo
import routes.updateTodo as updateTodo

if __name__ == "__main__":
    runTests([
        *getTodayTodos.tests,
        *createTodo.tests,
        *updateTodo.tests,
    ])
