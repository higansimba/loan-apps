
# Loan APPS

This projects built with golang and mongodb. Just it...not much stack here.


## How To Run It
```
1. Preparing Environtment
Language: golang
Database: mongodb
DatabaseName: loan
```


```
2. Preparing Repository
rename: .env.example --> .env
# clone the repository
cd loan-app
cd go mod tidy
cd go mod vendor

go run main.go

```

## Documentation
import the postman json file
```Loan Apps.postman_collection.json```

## How To Use It
1. Get Token For Accessing The API
``` curl --location 'localhost:8888/token' ```

2. Use Seeder to make dummy user and admin
``` curl --location 'localhost:8888/seed' ```

3. Use the endpoint you need and get data from database


