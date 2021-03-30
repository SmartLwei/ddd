# -- FILE: features/environment.py
import json
import os
import time


def before_all(context):
    print("before_all")
    print("开始时间：", int(time.time()))


def before_feature(context, feature):
    print("before_feature")


def before_scenario(context, scenario):
    print("before_scenario")


def before_step(context, step):
    print("before_step")


def after_step(context, step):
    print("after_step")


def after_scenario(context, scenario):
    print("after_scenario")


def after_feature(context, feature):
    print("after_feature")


def after_all(context):
    print("结束时间：", int(time.time()))


def download_swagger():
    cmd = "curl -sS --connect-timeout 3 --retry 5 --retry-delay 2 -o ./swagger.json http://localhost:8000/swagger/doc.json"
    if os.system(cmd) != 0:
        print("Download swagger.json failed.")
        exit(1)

    with open('swagger.json', 'r+') as f:
        data = json.load(f)
        data['host'] = "localhost:8000"
        data['basePath'] = "/"
        del data['info']['license']
        f.seek(0)
        json.dump(data, f, indent=4)
        f.truncate()
