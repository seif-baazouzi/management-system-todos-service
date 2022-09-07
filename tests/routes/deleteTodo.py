import utils
import config
from testRestApi import testRoute, Test, DELETE

def testToken():
    todoID = utils.createTodo()
    res = testRoute(DELETE, f"{config.server}/api/v1/todos/{todoID}")
    
    return res.equals({ "message": "invalid-token" })

def testDeleteNotExistingTodo():
    res = testRoute(DELETE, f"{config.server}/api/v1/todos/{utils.genUUID()}", headers={ "X-Token": config.token })
    return res.equals({ "message": "This todo does not exist" })

def testDeleteTodo():
    todoID = utils.createTodo()

    res = testRoute(DELETE, f"{config.server}/api/v1/todos/{todoID}", headers={ "X-Token": config.token })
    return res.equals({ "message": "success" })

tests = [
    Test("Delete Todo Invalid Token", testToken),
    Test("Delete Todo Not Existing Todo", testDeleteNotExistingTodo),
    Test("Delete Todo", testDeleteTodo),
]
