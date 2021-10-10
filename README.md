##Tools & Testing. 
Anything to support Dev 

##Start all services 
docker-compose up 

## Start with profile 
# by implicitly enabling profile `dev` or profile config in docker-compose.yaml 
1. docker-compose --profile=dev up
2. docker-compose --profile=redis up

### With Mongodb change connection string 
mongodb://u_test:Him3d3jmuDGD@localhost:27017/services?authSource=admin