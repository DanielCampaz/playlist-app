# Playlist App (Api)

[Download and Install GO](https://go.dev/doc/install)

### Command to start the api

- go run main.go

### Authorization

In the header, the token must be sent for all endpoints except login and singup,
Example:

```json
{
  "Content-Type": "application/json",
  "Token": "Your Token"
}
```

### .Env

For the api to work, you must add the environment variables at the root of the project at the same level as main.go, create an .env file and put the following environment variables with your own data:

```
  MYSQL_USER = root
  MYSQL_PASSWORD = 123456789
  MYSQL_PORT = 3306

  MYSQL_NAMEDATABASE = playlist

  API_NAME = api-v1

  APLICATION_PORT = :8080

  SECRETKEYJWT = 123456789

  SMTPSERVER = smtp.gmail.com

  SMTPPORT = 587

  SMTPUSERNAME = example@gmail.com

  SMTPPASSWORD = "your password SMTP"

  SMTPFROMEMAIL = example@gmail.com
```

[Github PlaylistApp](https://github.com/DanielCampaz/playlist-app)

## All this code was developed by Daniel Campaz
