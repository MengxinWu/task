echo "start build docker images..."
docker build -t admin:beta -f dockers/admin.Dockerfile .
docker build -t dispatch:beta -f dockers/dispatch.Dockerfile .
docker build -t monitor:beta -f dockers/monitor.Dockerfile .
docker build -t executor:beta -f dockers/executor.Dockerfile .

echo "remove docker container..."
docker rm -f task_admin
docker rm -f task_dispatch
docker rm -f task_monitor
docker rm -f task_executor

echo "run docker container..."
docker run -d --name task_admin --network tasknet --network-alias task_admin -p 48080:8080  admin:beta
docker run -d --name task_dispatch --network tasknet --network-alias task_dispatch dispatch:beta
docker run -d --name task_monitor --network tasknet --network-alias task_monitor monitor:beta
docker run -d --name task_executor --network tasknet --network-alias task_executor executor:beta