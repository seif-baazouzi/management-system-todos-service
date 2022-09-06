import json
import requests

# logging

RED    = "\033[31m"
BOLD   = "\033[1m"
RESET  = "\033[0m"
GREEN  = "\033[32m"
YELLOW = "\033[33m"


def green(msg):
  return GREEN + BOLD + str(msg) + RESET

def yellow(msg):
  return YELLOW + BOLD + str(msg) + RESET

def red(msg):
  return RED + BOLD + str(msg) + RESET

def bold(msg):
  return BOLD + str(msg) + RESET

# testing

GET    = "GET"
POST   = "POST"
PUT    = "PUT"
PATCH  = "PATCH"
DELETE = "DELETE"

class Response:
    def __init__(self, res):
        self.raw     = res
        self.status  = res.status_code
        self.headers = res.headers
        self.body    = json.loads(res.text)
    
    # check if the following keys are exist in the body
    def hasKeys(self, *keys):
        for key in keys:
            if key not in self.body:
                return False
        
        return True
    
    # check if the following keys are exist in the body and they are equals to there values
    def equals(self, keyValuePairs):
        for key, value in keyValuePairs.items():
            if key not in self.body or self.body[key] != value:
                return False
        
        return True

def testRoute(method, url, headers={}, body={}):
    res = requests.request(method, url=url, headers=headers, json=body)
    res = Response(res)

    print(f"{yellow(method)} ({bold(res.status)}) {url}")

    return res

class Test:
    def __init__(self, title, function):
        self.title = title
        self.function = function

    def run(self):
        print(f"Testing {self.title}")
        res = self.function()
        resMessage = green("success") if res else red("failed")
        
        print(f"Test {self.title} {resMessage}\n")
        
        return res

def runTests(tests):
    testsCount = 0
    successCount = 0

    for test in tests:
        if test.run():
            successCount += 1
        testsCount += 1

    failedCount = testsCount - successCount
    print(f"\nRun {yellow(testsCount)} tests, {green(successCount)} are success and {red(failedCount)} are failed")

    return 0 if failedCount == 0 else 1
