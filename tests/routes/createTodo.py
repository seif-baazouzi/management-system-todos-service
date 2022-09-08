import utils
import config
from testRestApi import testRoute, Test, POST

def testToken():
    workspaceID = utils.createWorkspace()
    res = testRoute(POST, f"{config.server}/api/v1/todos/{workspaceID}")
    return res.equals({ "message": "invalid-token" })

def testEmptyFields():
    workspaceID = utils.createWorkspace()
    res = testRoute(POST, f"{config.server}/api/v1/todos/{workspaceID}", headers={ "X-Token": config.token })
    return res.equals({ "title": "Must not be empty" })

def testNotExistingWorkspace():
    res = testRoute(POST, f"{config.server}/api/v1/todos/{utils.genUUID()}", headers={ "X-Token": config.token })
    return res.status == 400

def testCreateTodo():
    workspaceID = utils.createWorkspace()
    body = { "title": utils.randomString(10), "body": utils.randomString(50), "startingDate": utils.today() }

    res = testRoute(POST, f"{config.server}/api/v1/todos/{workspaceID}", headers={ "X-Token": config.token }, body=body)
    return res.equals({ "message": "success"})


tests = [
    Test("Create Todo Invalid Token", testToken),
    Test("Create Todo Empty Fields", testEmptyFields),
    Test("Create Todo Not Existing Workspace", testNotExistingWorkspace),
    Test("Create Todo", testCreateTodo),
]
