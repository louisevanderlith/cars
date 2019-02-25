# cars
Mango Web: Cars

## Run with Docker
*$ go build
*$ docker build -t avosa/cars:dev .
*$ docker rm carsDEV
*$ docker run -d --network host --name carsDEV avosa/cars:dev 
*$ docker logs carsDEV