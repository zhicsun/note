# 常用 shell

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
netstat -naltp | grep 服务器端口号 | awk '{print $5}' | awk -F ':' '{print $1}' | sort | uniq
```