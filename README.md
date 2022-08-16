# 通用任务系统

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
