import utils
import config
from testRestApi import testRoute, Test, DELETE

def testToken():
    workspaceID = utils.createWorkspace()
    res = testRoute(DELETE, f"{config.server}/api/v1/todos/workspace/{workspaceID}")
    
    return res.equals({ "message": "invalid-token" })

def testDeleteNotExistingWorkspace():
    res = testRoute(DELETE, f"{config.server}/api/v1/todos/workspace/{utils.genUUID()}", headers={ "X-Token": config.token })
    return res.status == 400

def testDeleteWorkspaceTodos():
    workspaceID = utils.createWorkspace()

    res = testRoute(DELETE, f"{config.server}/api/v1/todos/workspace/{workspaceID}", headers={ "X-Token": config.token })
    return res.equals({ "message": "success" })

tests = [
    Test("Delete Workspace Todos Invalid Token", testToken),
    Test("Delete Workspace Todos Not Existing Workspace", testDeleteNotExistingWorkspace),
    Test("Delete Workspace Todos", testDeleteWorkspaceTodos),
]
