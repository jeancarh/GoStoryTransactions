# Stori Structure

## Project Directory Structure

- `main.go`             : The entry point of your application
- `.env`             : you can put your custom variables here
- **email/**              : Package for email-related code
   - `email.go`        : Email sending logic
   - `templates/`      : HTML email templates
- **db/**                 : Package for database (MongoDB) related code
   - `db.go`           : Database interaction logic
- **csv/**                : Package for CSV file processing
   - `csv.go`          : Logic for reading and parsing CSV files
- **models/**             : Data models and structs
   - `transaction.go`  : Structs to represent transactions
- `Dockerfile`          : Dockerfile for building your application

## Installation
You have to install this version golang 1.21.3
you can install the dependencies with this commands

```bash
go mod download
```
Or
```bash
go get
```

## Usage

If you want to use this without docker you can run the project like this:
```go
go run main.go
```
If you want to use this project with docker, follow next steps

```docker
# this will build the project in a docker image
docker build -t storiapp .

# this will run the project in docker
docker run storiapp
```

Please keep in mind that you have to put in the enviroment variables your recipient email so you can check the emails in your personal email, you have to change only that in dockerfile

```docker
ENV \
    emailRecipient="<your-email>" \
```

If you are going to use this project only with go, please go to .env file and change the same variable there

```text
emailRecipient=<your-email>
```

## Notes

The project is saving the data of each transaction into the database of mongo in atlas, the email server is a free email server from brevo you have 300 emails per day in a free tier.