eval `aws ecr get-login`
docker rm -f appserver
docker pull 532898105683.dkr.ecr.us-east-1.amazonaws.com/in-touch-server:latest
docker run --name appserver -p 80:80 -d 532898105683.dkr.ecr.us-east-1.amazonaws.com/in-touch-server:latest
