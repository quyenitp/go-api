# go-api
Simple Go RESTful API 

# Building
`go build`

# Run 
`go run main.go` 

# Dockerize 
docker build -t go-api .

# Docker run command 
docker run -d -p 8080:8080 go-api 

# API 
`http://localhost:8000/users/1`

`http://localhost:8000/users/1/accounts`

`http://localhost:8000/accounts/1`