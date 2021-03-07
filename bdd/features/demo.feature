# -- FILE: features/example.feature
Feature: Showing off behave

  Scenario: 对demo的接口进行测试
    Given 新建客户端并且服务端已经开启
    When 新建项目
      | name      |
      | test_item |
    Then 项目的数量增加了一个
      | name      |
      | test_item |