# Using GORM with Golang Microservices
This repository provides example code for a blog post that guides users on integrating GORM with Go microservices. It covers the essentials of setting up GORM and using it effectively in a microservices architecture.

## Running Locally

To get started with the example code, you'll first need to set up the database. You can do this by running:

```bash
docker-compose up
```

This command will launch a PostgreSQL container and configure a database for testing purposes.

Next, to start the HTTP server locally, use the following command:

```bash
make start
```

You can then make requests to your server using tools like Postman, cURL, or any HTTP client of your choice.

Happy coding!
