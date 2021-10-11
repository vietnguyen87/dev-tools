##Tools & Testing. 
Wishing all our developers to have full services in the local environment. So this repo provides a `docker-compose` file build all services like mongo, kafka, es, redis. 
Hope you feel free to contribute as soon as there is an internet problem.

##Start all services 
docker-compose up 

## Start with profile 

##For Dev includes: mongo, redis, kafka. 
docker-compose --profile=dev up

##Redis 
docker-compose --profile=redis up

##kafka & kafkaDrop 
docker-compose --profile=kafka up

##MongoDB  
docker-compose --profile=mongodb up

### With Mongodb change connection string 
mongodb://u_test:Him3d3jmuDGD@localhost:27017/services?authSource=admin

##ElasticSearch   
docker-compose --profile=elasticsearch up
