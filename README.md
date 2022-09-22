

# REST API service 

A simple Rest API service Using Golang, which exposes endpoints to List, Show and Update data

## Available Features

| Feature                      | URL                                                                                      |
|------------------------------|------------------------------------------------------------------------------------------| 
| List All products            | http://localhost:3000/api/v1/products?offset=<StartingFrom>&limit=<No of items per page> |
| List Single Product          | http://localhost:3000/api/v1/products/<ProductId>                                        |
| Update title of the Product  | http://localhost:3000/api/v1/products/<ProductId>                                                     |

## Prerequisite

- [Go] - go1.19.1 darwin/amd64

## Steps to Run the App

```sh
cd Golang-REST-Service
go run main.go  
```

## Steps to Run Unit test App

```sh
cd Golang-REST-Service/app/tests
go test
```

## API Doc

https://documenter.getpostman.com/view/1756122/2s7ZE8o36G

## Sample Requests:

**Get /products**

```sh
curl --location --request GET 'http://localhost:3000/api/v1/products?offset=1&limit=10' \
--header 'Content-Type: application/json'`
 ```

**GET /products/{prodId}**
```sh
curl --location --request GET 'http://localhost:3000/api/v1/products/1' \
   --header 'Content-Type: application/json'`
```

**PUT /products/{prodId}**

```sh
curl --location --request PUT 'http://localhost:3000/api/v1/products/1' \
--header 'Content-Type: application/json' \
--header 'Authorization: Basic YWRtaW46YWRtaW4=' \
--data-raw '{
"title": "Apple - wwww"
}'`
```
