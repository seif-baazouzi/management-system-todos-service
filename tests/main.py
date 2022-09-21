from testRestApi import runTests

import routes.getTodayTodos as getTodayTodos
import routes.getMonthTodos as getMonthTodos
import routes.createTodo as createTodo
import routes.updateTodo as updateTodo
import routes.deleteTodo as deleteTodo
import routes.deleteWorkspaceTodos as deleteWorkspaceTodos

if __name__ == "__main__":
    runTests([
        *getTodayTodos.tests,
        *getMonthTodos.tests,
        *createTodo.tests,
        *updateTodo.tests,
        *deleteTodo.tests,
        *deleteWorkspaceTodos.tests,
    ])
