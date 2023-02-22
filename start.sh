#后台启动github加速器
nohup ../fastgithub/fastgithub_linux-x64/fastgithub &

#切换要部署的分支
git checkout dev

#拉取最新代码
git pull origin

#停止删除容器
docker-compose stop
docker-compose rm -f

#删除镜像
docker rmi tiktok_user
docker rmi tiktok_video
docker rmi tiktok_gateway

#重新启动容器
docker-compose up -d

#显示容器
docker ps

#显示日志
echo "<<<<<<<<<<<<<< user 模块错误日志 >>>>>>>>>>>>>>>"
docker logs tiktok_user | grep error

echo "<<<<<<<<<<<<  video 模块错误日志 >>>>>>>>>>>>>>"
docker logs tiktok_video | grep error

echo "<<<<<<<<<<<<<< gateway 模块错误日志 >>>>>>>>>>>>>>>"
docker logs tiktok_gateway | grep error