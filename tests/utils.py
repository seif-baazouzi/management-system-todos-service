import uuid
import string
import random

import config
from testRestApi import testRoute, POST

def randomString(length):
    return ''.join(random.choices(string.ascii_letters, k=length))

def genUUID():
    return str(uuid.uuid4())
    
def createWorkspace():
    workspace = randomString(10)
    parentWorkspace = None

    body = { "workspace": workspace, "parentWorkspace": parentWorkspace }
    res = testRoute(POST, f"{config.workspacesService}/api/v1/workspaces", headers={ "X-Token": config.token }, body=body)
    
    workspaceID = False if "workspaceID" not in res.body else res.body["workspaceID"]

    return workspaceID
