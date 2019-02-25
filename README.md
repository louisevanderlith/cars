# cars
Mango Web: Cars

## Run with Docker
*$ go build
*$ docker build -t avosa/cars:dev .
*$ docker rm carsDEV
*$ docker run -d -p 8081:8081 --network mango_net --name carsDEV avosa/cars:dev 
*$ docker logs carsDEV