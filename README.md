# Web scraping service - Test project

## Development Setup

# Prerequisites
- The application requires Node, Golang, Docker, Git

# analyzer-api
`cd analyzer-api`
`go mod download`
`go run main.go`

# web-client
`cd web-client`
`npm install`
`npm start`


### Build and deploy docker images 

# analyzer-api

# build api docker image
`cd analyzer-api`
`docker build -t analyzer-api .`

## start api docker container
`docker run -p 8080:8080 -it analyzer-api`

## web-client

# build web docker image
`cd web-client`
`docker build -t app:dev .`

## start web docker container
`docker run -it --rm -v ${pwd}:/app -v /app/node_modules -p 3000:3000 app:dev`


#### Assumptions and Decisions

-  As a Golang knowledge testing project, not worried about application security, UI design, and a hundred percent accuracy of the extracted data.
- Designed and tested with a few URLs.
- Designed as a separate service except for designing a monolithic web application.  

##### Suggestions 

- It's better to develop and deploy this as a microservice witch responsible for a single task and can follow a flat structure without a complex folder structure. And better to deploy in a serverless architecture, where we can handle lots of things without any overhead.