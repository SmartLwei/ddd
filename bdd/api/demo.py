import json

import requests

host_g = "http://localhost:8080"
headers_g = {"accept": "*/*", "Content-Type": "application/json;charset=UTF-8"}


class DemoService:

    def __init__(self):
        self.host = host_g
        self.header = headers_g

    def service_alive(self):
        txt = ""
        try:
            r = requests.get(url=self.host, headers=self.header)
            txt = r.json()
        except requests.exceptions.ConnectionError as e:
            pass
        return txt == "service is alive"

    def add_demo(self, name):
        full_path = '%s/demo/api/v1/demo' % self.host
        data = {"name": name}
        r = requests.post(url=full_path, headers=self.header, json=data)
        resp = json.loads(r.content)
        return resp

    def get_demos(self, name):
        full_path = '%s/demo/api/v1/demos?name=%s' % (self.host, name)
        r = requests.get(url=full_path, headers=self.header)
        resp = json.loads(r.content)
        return resp


if __name__ == "__main__":
    demo_svc = DemoService()
    if demo_svc.service_alive():
        print("service is alive, now run demos")
        demo_svc.add_demo("test_name")
        demo_svc.get_demos("test_name")
