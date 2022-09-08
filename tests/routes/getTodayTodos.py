import utils
import config
from testRestApi import testRoute, Test, GET

def testToken():
    res = testRoute(GET, f"{config.server}/api/v1/todos/today")
    return res.equals({ "message": "invalid-token" })

def testGetTodayTodos():
    res = testRoute(GET, f"{config.server}/api/v1/todos/today", headers={ "X-Token": config.token })
    return res.hasKeys("todos")

tests = [
    Test("Get Todos Invalid Token", testToken),
    Test("Get Today Todos", testGetTodayTodos),
]
