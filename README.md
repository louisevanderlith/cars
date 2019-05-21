# cars
Mango Web: Cars
Cars will act as the central platform for advertising vehicles.

## Run with Docker
* $ docker build -t avosa/cars:dev .
* $ docker rm carsDEV
* $ docker run -d -e RUNMODE=DEV -p 8081:8081 --network mango_net --name CarsDEV avosa/cars:dev 
* $ docker logs carsDEV