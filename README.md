# JWT-GO-DEMO

这个 demo 用 Go 语言进行 JWT 认证

所使用的中间件 [jwt-go](https://pkg.go.dev/github.com/dgrijalva/jwt-go?tab=doc)

参考教程 [Securing Your Go REST APIs with JWTS](https://www.youtube.com/watch?v=-Scg9INymBs)

## GET 参数说明

使用默认的 mySigningKey 执行 generateJWT() 函数后，GET Header 的 Token 参数为: 

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1OTczNzg1MzcsInVzZXIiOiJ3aW5uZXJ3aW50ZXIifQ.PPGqKZe1tFanH9YQyG28HHUV-HLRC1aLkD5xGwST-30
```