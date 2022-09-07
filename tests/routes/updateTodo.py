import utils
import config
from testRestApi import testRoute, Test, PUT

def testToken():
    todoID = utils.createTodo()
    res = testRoute(PUT, f"{config.server}/api/v1/todos/{todoID}")
    return res.equals({ "message": "invalid-token" })

def testEmptyFields():
    todoID = utils.createTodo()
    res = testRoute(PUT, f"{config.server}/api/v1/todos/{todoID}", headers={ "X-Token": config.token })
    print(res.body)
    return res.equals({ "title": "Must not be empty" })

def testNotExistingTodo():
    res = testRoute(PUT, f"{config.server}/api/v1/todos/{utils.genUUID()}", headers={ "X-Token": config.token })
    return res.status == 404

def testUpdateTodo():
    todoID = utils.createTodo()
    body = { "title": utils.randomString(10), "body": utils.randomString(50), "done": True }

    res = testRoute(PUT, f"{config.server}/api/v1/todos/{todoID}", headers={ "X-Token": config.token }, body=body)
    return res.equals({ "message": "success"})


tests = [
    Test("Update Todo Invalid Token", testToken),
    Test("Update Todo Empty Fields", testEmptyFields),
    Test("Update Todo Not Existing Todo", testNotExistingTodo),
    Test("Update Todo", testUpdateTodo),
]
