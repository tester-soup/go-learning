*go学习写的一些代码。不值得一看。*
<!-- TOC -->

- [familyaccount 家庭收支记账软件](#familyaccount-%E5%AE%B6%E5%BA%AD%E6%94%B6%E6%94%AF%E8%AE%B0%E8%B4%A6%E8%BD%AF%E4%BB%B6)
    - [需求](#%E9%9C%80%E6%B1%82)
    - [用户界面](#%E7%94%A8%E6%88%B7%E7%95%8C%E9%9D%A2)
    - [扩展功能](#%E6%89%A9%E5%B1%95%E5%8A%9F%E8%83%BD)
- [customermanage 客户信息管理系统软件](#customermanage-%E5%AE%A2%E6%88%B7%E4%BF%A1%E6%81%AF%E7%AE%A1%E7%90%86%E7%B3%BB%E7%BB%9F%E8%BD%AF%E4%BB%B6)
    - [需求](#%E9%9C%80%E6%B1%82)
    - [用户界面](#%E7%94%A8%E6%88%B7%E7%95%8C%E9%9D%A2)
- [chatroom 多人聊天系统](#chatroom-%E5%A4%9A%E4%BA%BA%E8%81%8A%E5%A4%A9%E7%B3%BB%E7%BB%9F)
    - [用户界面](#%E7%94%A8%E6%88%B7%E7%95%8C%E9%9D%A2)
    - [扩展功能](#%E6%89%A9%E5%B1%95%E5%8A%9F%E8%83%BD)

<!-- /TOC -->

## [familyaccount](./familyaccount/) 家庭收支记账软件
### 需求
1. 模拟实现基于文本界面的《家庭记账软件》
2. 该软件能够记录家庭的收入、支出，并且能够打印收支明细

### 用户界面
```dotnetcli
----------------家庭收支记账软件----------------
                   1 收支明细
                   2 登记收入
                   3 登记支出
                   4 退出软件
请选择（1-4）
1
----------------当前收支明细记录----------------
收支    账户金额        收支金额        说明
收入      10100           100        收到红包
支出      10050            50         吃饭
----------------家庭收支记账软件----------------
                   1 收支明细
                   2 登记收入
                   3 登记支出
                   4 退出软件
请选择（1-4）
```
### 扩展功能
- 对以上项目网一个转账功能
- 在使用该软件前有一个登录功能，只有输入正确的用户名和密码才能操作。

## [customermanage](./customermanage/) 客户信息管理系统软件
### 需求
1. 模拟实现基于文本界面的《客户信息管理软件》
2. 该软件能够实现对客户对象的插入、修改和删除（用切片实现），并且能够打印客户明细表.
### 用户界面
```dotnetcli
------------------客户信息管理系统-----------------
                     1 添加客户
                     2 修改客户
                     3 删除客户
                     4 客户列表
                     5 退    出
请选择1-5
1
------------------添加客户-----------------
姓名：test
性别：male
年龄：18
电话：1234567
邮件：123@gmail.com
------------------添加完成-----------------
------------------客户信息管理系统-----------------
                     1 添加客户
                     2 修改客户
                     3 删除客户
                     4 客户列表
                     5 退    出
请选择1-5
4
------------------客户列表-----------------
编号    姓名    性别    年龄    电话    邮箱
1       张三    男      20      111     zs@qq.com
2       test    male    18      1234567 123@gmail.com
------------------客户列表end-----------------
------------------客户信息管理系统-----------------
                     1 添加客户
                     2 修改客户
                     3 删除客户
                     4 客户列表
                     5 退    出
请选择1-5
2
------------------修改客户-----------------
请输入要修改的客户ID(-1退出)：2
姓名(test)：<直接回车则不修改>
性别(male)：
年龄(18)：
电话(1234567)：1234567890
邮件(123@gmail.com)：test@gmail.com

编号    姓名    性别    年龄    电话    邮箱
2       test    male    18      1234567890      test@gmail.com
------------------修改完成-----------------
------------------客户信息管理系统-----------------
                     1 添加客户
                     2 修改客户
                     3 删除客户
                     4 客户列表
                     5 退    出
请选择1-5
3
------------------删除客户-----------------
请输入要删除的客户ID(-1退出)：
2
编号    姓名    性别    年龄    电话    邮箱
2       test    male    18      1234567890      test@gmail.com
确定要删除吗？y/n
y
------------------删除完成-----------------
------------------客户信息管理系统-----------------
                     1 添加客户
                     2 修改客户
                     3 删除客户
                     4 客户列表
                     5 退    出
请选择1-5
5
退出系统.
```
## [chatroom](./chatroom/) 多人聊天系统
### 用户界面
```dotnetcli
-----------欢迎进入多人聊天系统-----------
              1 登录聊天室
              2 注册用户
              3 退出系统
请选择（1-3）
1
请输入用户id
1
请输入用户密码
123
------------恭喜test01[ID:1]登录成功------------
          1 显示在线用户列表
              2 发送消息
              3 信息列表
              4 退出系统
请选择（1-4）
1
------------当前在线用户列表------------
ID：1
ID：2
------------end------------
```
### 扩展功能
- 实现私聊【点对点聊天】
- 用户离线，把他从在线列表剔除
- 实现离线留言。用户不在线，可以接收离线消息，登录后可接收
- 发送一个文件
