# bookstore-items-api
- User API
- Framework used 
    - gin-gonic/gin 
        - upto 40 times faster perofrmance
    - Provides http engine to be used
- When returning `nil` from a function return has to be pointer 
- the memory block containing the slice elements is passed by "reference".

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