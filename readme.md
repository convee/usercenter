# 基于go gin框架 用户中心

## 导入mysql 数据
```
    mysql > create database test;
    mysql > use user;
    msyql > source /home/convee/go/src/usercenter/user.sql;

```
## mysql默认用户密码 test 123456

## 运行项目
```
go build -o usercenter
./usercenter.exe
```
## 访问用户中心主页

http://127.0.0.1:8002/member