from testRestApi import runTests

import routes.getTodayTodos as getTodayTodos

if __name__ == "__main__":
    runTests([
        *getTodayTodos.tests,
    ])
