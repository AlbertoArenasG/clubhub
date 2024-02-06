# clubhub
Backend application built on Golang with MongoDB and Fiber.

## Requirements

- Docker
- Docker Compose

## Environment Setup

1. Clone this repository to your local machine:

```bash
git clone https://github.com/AlbertoArenasG/clubhub.git
cd clubhub
```

2. Create folder for local database data:

```bash
mkdir .mongo-data
```

3. Create a `.env` file based on the provided example:

```bash
cp .env.example .env
```

4. Edit the `.env` file and configure the environment variables as needed, such as the database connection.

## Running the Project

To run the project, simply execute the following command:

```bash
docker-compose up --build
```

The project will run at http://localhost:3000.

## Documentation

https://documenter.getpostman.com/view/6768805/2s9YywfKjV
