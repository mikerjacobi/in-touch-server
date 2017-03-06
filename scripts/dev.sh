go build
docker build -t in-touch-dev -f Dockerfile.dev . 
docker rm -f in-touch
docker run -p 80:8080 -d --name in-touch in-touch-dev
docker logs -f in-touch &
