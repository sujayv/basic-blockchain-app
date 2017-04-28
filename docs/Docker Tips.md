**a. Commands for first run or to remove previous data and docker images in order to run from scratch**
1. docker-compose down
2. docker rm $(docker ps -aq)
3. docker images | grep "dev-peer" | awk '{print $1}' | xargs docker rmi
4. docker-compose up -d (To start containers again)

**b. Commands to restart docker containers and run where you previously left off**
1. Run (a)
2. docker-compose kill (Run while exiting)
3. docker-compose up -d (Run the next time you want to start and then run node programs directly)
