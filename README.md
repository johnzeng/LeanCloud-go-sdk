# Lean cloud go SDK

Leancloud go语言SDK，根据Leancloud公布的Restful API封装，并且使用godep管理依赖。

使用：

```shell
go get github.com/johnzeng/leancloud-go-sdk
cd $GOPATH/src/github.com/johnzeng/leancloud-go-sdk
godep restore
go install
```

# 目前支持的功能：

- Object create
- Object get by id
- Object delete by id
- Object update by id
- Object query support
- cql query support
- requestSmsCode, verifySmsCode（需要在控制台设置开放这两个接口）
- push notification(待测试)
- 支持LeanTime，LeanPointer，LeanByte, LeanRelation, LeanFile类型。
- User,Installation, Role等系统结构体的集成

具体使用方法请参考对应的test文件。里面每个test都是完整的请求。


# Todo:

- Cloud function 支持
- 事件流API
- 实时通信API
- 数据Schema
- 统计数据API

# How to test:
需要设置以下环境变量：

```shell
export LEAN_APPID=aawgeiyg3491v
export LEAN_APPKEY=sdb9bg9408
export LEAN_MASTERKEY=aslng18gbiq98
export LEAN_TEST_PHONE_NUMBER=1111111111
export LEAN_TEST_USER=john
export LEAN_TEST_PWD=11111111111
```

根据自己的测试app情况进行设置接口（当然上面的是假的，怎么可能将真的key丢出来呢）

