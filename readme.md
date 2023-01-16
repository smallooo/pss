
### 编译可执行文件
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pss .

### build docker
docker build -t pss-docker-scratch .

### run 
docker run -p 8000:8000 pss-docker-scratch

### swagger
swag init