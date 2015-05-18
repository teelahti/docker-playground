## To build redis container

    docker build -t tbay/redis --rm=true .
    
## To run redis container

Use -d to run in the background

    docker run --name redis -d tbay/redis