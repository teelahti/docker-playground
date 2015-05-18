
## To build container

    docker build -t tbay/client . 

## To run and link to Redis container

see https://docs.docker.com/examples/running_redis_service/

    docker run --link redis:db -i -t tbay/client