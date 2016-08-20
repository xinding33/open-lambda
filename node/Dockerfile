# use LTS
FROM ubuntu:trusty

# DOCKER (and other scripts)
RUN apt-get update -y
RUN apt-get install -y docker.io python netcat wget curl
VOLUME ["/docker_vol"]

# RETHINKDB
RUN echo "deb http://download.rethinkdb.com/apt trusty main" > /etc/apt/sources.list.d/rethinkdb.list
RUN wget -qO- https://download.rethinkdb.com/apt/pubkey.gpg | apt-key add -
RUN apt-get update
RUN apt-get install -y rethinkdb
VOLUME ["/data"]

# OL WORKER
COPY bin /open-lambda/bin

# SCRIPTS
COPY startup.py /open-lambda/startup.py
COPY kill.py /open-lambda/kill.py

# BE READY TO RUN PROCESSES (Instead of additional containers)
RUN apt-get update -y
RUN apt-get install -y python-pip python-dev build-essential
RUN pip install rethinkdb
RUN pip install Flask
COPY /pyserver/server.py /open-lambda/pyserver/server.py
COPY /pyserver/lambda_func.py /open-lambda/pyserver/lambda_func.py
COPY /pyserver/config.json /open-lambda/pyserver/config.json

# INIT
CMD ["python", "/open-lambda/startup.py"]
