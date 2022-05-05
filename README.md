# bookstore-items-api
- User API
- Framework used
    - gin-gonic/gin
        - upto 40 times faster perofrmance
    - Provides necessary http engine

### General design principle
- If method is returning error, then method should log the error and then return.
Caller shouldn't be responsible for logging error received.
- Caller should be responsible for reacting to the error appropriately

### User API microservices
- Developed using MVC
#### MVC
    - Controller calls service and returns data based on presentation requirement
    - Service calls different business function
    - Model only responsible for holding and persisting data as needed
#### This microservices exposes following API's
  - GET "/ping"
  - GET "/users/search/"
	- GET "/users/:user_id"
	- POST "/users"
	- POST "/users/login"
	- PUT "/users/:user_id"
	- PATCH "/users/:user_id"
	- DELETE "/users/:user_id"

## MySQL DB
- To run MySQL as docker container, execute thiss
  ```
    sudo docker run \
    --detach \
    --name=msql-bookstore \
    --env="MYSQL_ROOT_PASSWORD=<<password>>" \
    --publish 6603:3306 \
    --volume=<<path-to-local-volume>>:/var/lib/mysql \
    mysql
  ```
- Log into my sql container, to initialize db
 ```
  docker exec -it <<container-id>> /bin/sh
  mysql -u root -p
  mysql -h localhost -p 6603 -u root -p
  create schema `users_db`
```
## When Running this microservices in developer mode on your local
- Make sure to set this env variables
``` 
export mysql_password=<<MYSQL_ROOT_PASSWORD>>
export mysql_username=root
export mysql_host=127.0.0.1:6603
export mysql_schema=users_db
```
- go run *.go

## When Running user-api microservices as docker
```
docker run -i -t -d -e mysql_username=root  -e mysql_password=<<MYSQL_ROOT_PASSWORD>>  -e mysql_host=172.17.0.2:3306 -e  mysql_schema=users_db -e auth_url=http://172.17.0.5:8082 -p 8081:8081 bookstore-user-api
```
- This sets env variabes as needed

## go clean cache
- go clean -modcache
- go get "x.y.z.lib" in case one has to use updated code 

## Docker command
- Use the docker network inspect bridge command. This will show you the containers currently attached to the bridge network:
  - ref: https://www.tutorialworks.com/container-networking/
- docker rm -f $(docker ps -a -q)
- docker rm $(docker ps --filter status=created -q)
- docker rm $(docker ps --filter status=exited -q)
