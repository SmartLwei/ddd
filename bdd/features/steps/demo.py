# -- FILE: features/steps/example_steps.py
from behave import given, when, then

from api.demo import DemoService


@given(u'新建客户端并且服务端已经开启')
def step_impl(context):
    context.demo_svc = DemoService()
    assert context.demo_svc.service_alive() is True
    context.count = dict()


@when(u'新建项目')
def step_impl(context):
    for row in context.table:
        item_name = row['name']
        resp = context.demo_svc.get_demos(item_name)
        assert resp is not None
        assert resp['code'] == 0
        context.count[item_name] = resp['data']['count']
        resp = context.demo_svc.add_demo(item_name)
        assert resp is not None
        assert resp['code'] == 0


@then(u'项目的数量增加了一个')
def step_impl(context):
    for row in context.table:
        item_name = row['name']
        resp = context.demo_svc.get_demos(item_name)
        assert resp is not None
        assert context.count[item_name] + 1 == resp['data']['count']
