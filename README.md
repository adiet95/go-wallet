# Go-Order-API
Simple app product order with Echo HTTP Response Framework, GORM for Object relation model, PostgreSQL for database.

## 🔗 Description

This Backend Application is used for simple order product, in this application there are two models / ERD Schema likes User / Costumer & Products.
Also have several features like JWT, Authentication & Authorization.
There are 3 main modules :
1. Customer Management (Get with paginate, Get Detail, Insert, Update, Delete,
Search)
2. Order Management (Get with paginate, Get Detail, Insert, Update, Delete,
Search)
3. Authentication Management (Get Login Data, Insert Login Data)

Notes :
1In this application there are two types of users (Roles). admins and costumer. 
   Admin can do *Costumer Management* but Role Costumer can't, Registration page can only register Costumer roles, Admins can only be registered through seeding data.

<h2 align="center">
 ERD (Entity Relation Database)
</h2>
<p align="center"><img src="https://res.cloudinary.com/dw5qffbop/image/upload/v1665874871/erd_c15gne.png" alt="erd.jpg" /></p>

<h2 align="center">
 Table Specification
</h2>

<h3 align="center">Costumer's Table</h3>
<p align="center"><img src="https://res.cloudinary.com/dw5qffbop/image/upload/v1665882605/table-cost_miwqjk.png" alt="cost.jpg" /></p>
<h3 align="center">Order's Table</h3>
<p align="center"><img src="https://res.cloudinary.com/dw5qffbop/image/upload/v1665882605/table-user_sz493j.png" alt="order.jpg" /></p>

## Several command you must know in this app :
```bash
1. go run . serve //to run the app / server
2. go run . migrate -u //for database migration
# or
go run . migrate -d //for rollback
3. go run . seed // to seeding data Role admin if u want Email : "admin@gmail.com" Pass : admin12345678
```

## 🛠️ Installation Steps

1. Clone the repository

```bash
https://go-wallet.git
```

2. Install dependencies

```bash
go mod tidy
```
> Wait a minute, if still error run 

```bash
go mod vendor
```

3. Add Env File

```sh
  DB_USER="postgres"
  DB_HOST="localhost"
  DB_NAME="go-order"
  DB_PASS="root"
  JWT_KEYS="Rahasiaaaa"
  PORT=":8080"
  DB_PORT="5432"
```

4. Database Migration and Rollback

```bash
go run main.go migrate --up //for database migration table
# or
go run main.go migrate --down //for rollback the database
```

5. Seeding data admin

```bash
go run . seed
```
_Purpose to login as Admin's Role_
```
{
  "email": "admin@gmail.com",
  "password": "admin12345678"
}
```

6. Run the app

```bash
go run . serve
```

### 🚀 You are all set

## 🔗 RESTful Documentation Endpoints
[API Documentation](https://documenter.getpostman.com/view/28477373/2sA3e2hA7c) <- Click Link


## 💻 Built with

- [Golang](https://go.dev/): Go Programming Language
- [Echo](https://echo.labstack.com/): for handle http request
- [Postgres](https://www.postgresql.org/): for DBMS


## 🚀 About Me

- Linkedin : [Achmad Shiddiq](https://www.linkedin.com/in/achmad-shiddiq-alimudin/)
