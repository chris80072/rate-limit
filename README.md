# Rate Limit

This is a simple rate limit service by using Redis. All the related services are base on docker.

### Description
The user can send 60 requests in the past 60 seconds. If the limit is exceeded, you will receive an error response.

### Why choose Redis?

Redis is a single thread non-relational database. It's good at read and write data because the data is saving in memory. Even if the data is lost due to accident, it will not cause harm to the main service.

### How to execute
1. download the project  
``` $ git clone https://github.com/chris80072/rate-limit ```
2. build docker image  
``` $ docker build -t ratelimit . --no-cache ```
3. execute the service  
``` $ docker-compose up -d ```
4. opne below the url in your browser  
    http://localhost:8080/