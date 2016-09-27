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
- requestSmsCode, verifySmsCode（需要在控制台设置开放这两个接口）
- push notification(待测试)
- 支持LeanCloud的Date，Byte格式。

具体使用方法请参考对应的test文件。里面每个test都是完整的请求。


# Todo:
以下事项将会按照顺序进行：

- Object query support
- cql query support
- 实时通信API
- User,Installation, role等系统结构体的集成
- Pointer，File以及Relation数据结构支持。
- 事件流API
- 数据Schema
- 统计数据API


