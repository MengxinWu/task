# 通用任务系统

## zookeeper & kafka

> docker pull zookeeper

> docker run -d --name zookeeper --network tasknet --network-alias zookeeper -p 2181:2181 -t zookeeper

> docker pull wurstmeister/kafka

> docker run -d --name kafka0 --network tasknet --network-alias kafka0 -p 9092:9092 -e KAFKA_BROKER_ID=0 -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://kafka0:9092 -e KAFKA_LISTENERS=PLAINTEXT://kafka0:9092 -t wurstmeister/kafka

> docker run -d --name kafka1 --network tasknet --network-alias kafka1 -p 9093:9093 -e KAFKA_BROKER_ID=1 -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://kafka1:9093 -e KAFKA_LISTENERS=PLAINTEXT://kafka1:9093 -t wurstmeister/kafka


## admin 任务管理服务

> docker build -t admin:beta -f dockers/admin.Dockerfile .

> docker run -d --name task_admin --network tasknet --network-alias task_admin -p 48080:8080  admin:beta


## dispatch 任务调度服务

> docker build -t dispatch:beta -f dockers/dispatch.Dockerfile .

> docker run -d --name task_dispatch --network tasknet --network-alias task_dispatch dispatch:beta


## monitor 任务监听服务

> docker build -t monitor:beta -f dockers/monitor.Dockerfile .

> docker run -d --name task_monitor --network tasknet --network-alias task_monitor monitor:beta


## executor 任务执行器

> docker build -t executor:beta -f dockers/executor.Dockerfile .

> docker run -d --name task_executor --network tasknet --network-alias task_executor executor:beta
