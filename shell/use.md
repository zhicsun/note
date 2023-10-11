# 常用命令

## 添加白名单

```shell
#!/bin/bash
ips=(ip1 ip2)

for v in ${ips[@]}
do
 ufw allow from $v to any port 端口号
done

ufw status
```

## 删除白名单

```shell
#!/bin/bash
ips=(ip1 ip2)

for v in ${ips[@]}
do
 ufw delete allow from $v to any port 端口号
done

ufw status
```

## 输出所有客户端 ip

```shell
netstat -naltp | grep 服务端口号 | awk '{print $5}' | awk -F ':' '{print $1}' | sort | uniq
```

## tcpdump 查看 post 和 get 请求数据
```shell
# get
tcpdump -s 0 -A 'tcp dst port 服务端口号 and (tcp[((tcp[12:1] & 0xf0) >> 2):4] = 0x47455420)' -i 网卡名

# post
tcpdump -s 0 -A 'tcp dst port 服务端口号 and (tcp[((tcp[12:1] & 0xf0) >> 2):4] = 0x504f5354)' -i 网卡名

tcpflow -cp -i 网卡名 port 服务端口号
```