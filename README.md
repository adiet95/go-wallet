
## Several command you must know in this app :
```bash
1. go run . serve //to run the app / server
2. go run . migrate -u //for database migration
# or
go run . migrate -d //for rollback
3. go run . seed // to seeding data Role admin if u want Phone : 081388355301 Pin : 123987
```

## 🛠️ Installation Steps

1. Clone the repository

```bash
https://github.com/adiet95/go-wallet.git
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
  DB_USER=""
  DB_HOST=""
  DB_NAME=""
  DB_PASS=""
  DB_PORT=""

  REDIS_URL=""
  REDIS_PASSWORD=""
  REDIS_DB=""

  JWT_KEYS="e48840cae9bb5eeef8e627d61e165e2e0029991c1bc1ac3829f7e87c0c78e569"
  PORT=":8080"
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
  "phone_number": "081388355301",
  "pin": "123987"
}
```

6. Run the app

```bash
go run . serve
```

### 🚀 You are all set

## 🔗 RESTful Documentation Endpoints
[API Documentation](https://documenter.getpostman.com/view/22320158/2sAYdZtZ9i) <- Click Link
API Documentation
```bash

https://documenter.getpostman.com/view/22320158/2sAYdZtZ9i

```


## 💻 Built with

- [Golang](https://go.dev/): Go Programming Language
- [Echo](https://echo.labstack.com/): for handle http request
- [Postgres](https://www.postgresql.org/): for DBMS
- [Redis](https://redis.io/): Transactional Report


## 🚀 About Me

- Linkedin : [Achmad Shiddiq](https://www.linkedin.com/in/achmad-shiddiq-alimudin/)
