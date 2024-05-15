# Tron 安装

## 准备

```shell
# bt https://www.hostcli.com/
apt-get install vim git openjdk-8-jdk supervisor
mkdir -p ~/tron
```

## 数据

```shell
# http://3.219.199.168/
vim d.txt 
http://3.219.199.168/backup2024/FullNode_LevelDB_08.tar.gz
http://3.219.199.168/backup2024/FullNode_LevelDB_07.tar.gz
http://3.219.199.168/backup2024/FullNode_LevelDB_06.tar.gz
http://3.219.199.168/backup2024/FullNode_LevelDB_05.tar.gz
http://3.219.199.168/backup2024/FullNode_LevelDB_04.tar.gz
http://3.219.199.168/backup2024/FullNode_LevelDB_03.tar.gz
http://3.219.199.168/backup2024/FullNode_LevelDB_02.tar.gz
http://3.219.199.168/backup2024/FullNode_LevelDB_01.tar.gz
http://3.219.199.168/backup2024/FullNode_LevelDB_00.tar.gz
http://3.219.199.168/backup2024/FullNode_LevelDB.md5sum
wget -i d.txt -bc

vim tar.sh
#!/bin/bash
cat FullNode_LevelDB_0* | tar zxv
nohup sh tar.sh 2>&1 &
```

## kafka

```shell
git clone https://github.com/tronprotocol/event-plugin.git
cd event-plugin
./gradlew build

cd ~/tron
wget https://downloads.apache.org/kafka/3.7.0/kafka_2.13-3.7.0.tgz
tar -xzvf kafka_2.13-3.7.0.tgz
mv kafka_2.13-3.7.0 kafka

vim kafka/config/server.properties
# listeners=PLAINTEXT://0.0.0.0:9092
# advertised.listeners=PLAINTEXT://ip:9092

supervisord
vim /etc/supervisor/conf.d/tron.conf
[program:zookeeper]
process_name=%(program_name)s_%(process_num)02d
directory=/root/tron/kafka
command=/root/tron/kafka/bin/zookeeper-server-start.sh /root/tron/kafka/config/zookeeper.properties
numprocs=1
startsecs=5
redirect_stderr=true
autostart=true
autorestart=true
user=root
log_stdout=true
log_stderr=true
stdout_logfile=/data/logs/zookeeper.log
stopasgroup=true
killasgroup=true

[program:kafka]
process_name=%(program_name)s_%(process_num)02d
directory=/root/tron/kafka
command=/root/tron/kafka/bin/kafka-server-start.sh /root/tron/kafka/config/server.properties
numprocs=1
startsecs=5
redirect_stderr=true
autostart=true
autorestart=true
user=root
log_stdout=true
log_stderr=true
stdout_logfile=/data/logs/kafka.log
stopasgroup=true
killasgroup=true

mkdir -p /data/logs/
supervisorctl update

kafka/bin/kafka-topics.sh --create --topic block --bootstrap-server localhost:9092
kafka/bin/kafka-topics.sh --create --topic transaction --bootstrap-server localhost:9092
kafka/bin/kafka-topics.sh --create --topic contractevent --bootstrap-server localhost:9092
kafka/bin/kafka-topics.sh --create --topic contractlog --bootstrap-server localhost:9092
kafka/bin/kafka-topics.sh --create --topic solidity --bootstrap-server localhost:9092
kafka/bin/kafka-topics.sh --create --topic solidityevent --bootstrap-server localhost:9092
kafka/bin/kafka-topics.sh --create --topic soliditylog --bootstrap-server localhost:9092

kafka/bin/kafka-console-consumer.sh --topic  --from-beginning --bootstrap-server localhost:9092
```

## 节点

```shell
wget https://github.com/tronprotocol/java-tron/releases/download/GreatVoyage-v4.7.4/FullNode.jar

vim /etc/supervisor/conf.d/tron.conf
[program:tron]
process_name=%(program_name)s_%(process_num)02d
directory=/root/tron/
command=java -Xmx80g -XX:+UseConcMarkSweepGC -jar FullNode.jar -c main_net_config.conf --es
numprocs=1
startsecs=5
redirect_stderr=true
autostart=true
autorestart=true
user=root
log_stdout=true
log_stderr=true
stdout_logfile=/data/logs/tron.log
stopasgroup=true
killasgroup=true

supervisorctl update
```