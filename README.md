# docker-playground
Docker playground for multi-container communication practice. 

Building the machines could be done completely with Docker commands, but for clarity both containers have a Dockerfile.

If you need an ad hoc Redis CLI client for testing, run:

    docker run --link redis:db -i -t ubuntu:14.04 /bin/bash
	
...and then install redis with:

	$ sudo apt-get update
	$ sudo apt-get install redis-server
	$ sudo service redis-server stop

Now you can connect to our Redis instance ("DB" is the link name):

    redis-cli -h $DB_PORT_6379_TCP_ADDR

Detailed instructions on [Docker website](https://docs.docker.com/examples/running_redis_service/).
