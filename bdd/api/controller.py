import json

import requests

host = "http://localhost:8080"
headers = {"accept": "*/*", "Content-Type": "application/json;charset=UTF-8"}


class ControllerService:
    @classmethod
    def demo(cls):
        url = '%s/controller/v1/demo?hello=world' % host
        r = requests.get(url=url, headers=headers)
        resp = json.loads(r.content)
        print(json.dumps(resp, indent=4))
        return resp


if __name__ == "__main__":
    resp = ControllerService().demo()
