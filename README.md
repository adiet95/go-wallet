
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
