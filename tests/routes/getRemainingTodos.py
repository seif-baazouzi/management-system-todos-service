import utils
import config
from testRestApi import testRoute, Test, GET

def testToken():
    res = testRoute(GET, f"{config.server}/api/v1/todos/remaining")
    return res.equals({ "message": "invalid-token" })

def testGetTodayTodos():
    res = testRoute(GET, f"{config.server}/api/v1/todos/remaining", headers={ "X-Token": config.token })
    return res.hasKeys("todos")

tests = [
    Test("Get Remaining Todos Invalid Token", testToken),
    Test("Get Remaining Todos", testGetTodayTodos),
]
