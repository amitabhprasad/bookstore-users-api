# bookstore-items-api
- User API
- Framework used 
    - gin-gonic/gin 
        - upto 40 times faster perofrmance
    - Provides http engine to be used
- When returning `nil` from a function return has to be pointer 
- the memory block containing the slice elements is passed by "reference".

## General
- If method is returning error, then method should log the error and then return.
Caller shouldn't be responsible for logging error received.
- Caller should be responsible for reacting to the error appropriately

## MySQL DB
- 
    sudo docker run \
    --detach \
    --name=msql-bookstore \
    --env="MYSQL_ROOT_PASSWORD=<chamgeMe>" \
    --publish 6603:3306 \
    --volume=/Users/amitabhprasad/projects/goworkspace/src/github.com/amitabhprasad/bookstore-app/storage/docker/mysql-data:/var/lib/mysql \
    mysql

 docker exec -it ff89dca5d590 /bin/sh
 mysql -u root -p 
 mysql -h localhost -p 6603 -u root -p 

 run  
 export mysql_users_password= and set other env variables 
 export mysql_users_username=root
 export mysql_users_host=127.0.0.1:6603
export mysql_users_schema=users_db
## UserAPI
- Developed using MVC
### MVC
    - Controller calls service and returns data based on presentation requirement 
    - Service calls different business function
    - Model only responsible for holding and persisting data as needed 

## OAuth
- Developed using domain driven design


## Elastic Search
- 
    docker run 
    --name es01 
    --net elastic -p 9200:9200 -p 9300:9300
    --volume=/Users/amitabhprasad/projects/goworkspace/src/github.com/amitabhprasad/bookstore-app/elasticsearch/data
    docker.elastic.co/elasticsearch/elasticsearch:8.0.0

