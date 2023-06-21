# Blinkr

A REST-ful API with CRUD functionality written in Go.

This application uses MongoDB for data storage. It lets you create, read, update and delete an object called "Blink"from the MongoDB container using http requests.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Prerequisites
Before running the API, ensure you have the following prerequisites installed:

- [Go](https://go.dev/) (version 1.20.4)
- [Docker](https://www.docker.com/)

## Installation
1. Clone the repository:

```shell
https://github.com/mar-cial/blinkr
```

2. Change to the project directory:

```shell
cd blinkr
```

3. Install the dependencies:

```shell
go mod download
```

4. Set up the database and rest server configuration by creating a `.env` file in the root directory:

```plaintext
MONGOURI=mongodb://root:password@db:dbport
DBUSER=root
DBPASS=password
DBHOST=db
DBPORT=27017
DBNAME=testdb
DBCOLL=testcoll
SERVERPORT=8000
```

The Docker compose file defines default values for all the necessary env variables except DBPASS and MONGOURI. Those need to be set.

5. Install ui dependencies:

This projects uses [Next.js 13](https://nextjs.org/) with app router as UI.

```shell
cd ui
yarn
```

## Usage
1. Run the application:

With a single `docker compose up --build` you can get the three containers that make up this application up and running

```shell
docker compose up --build
```

2. The API server will start running on `http://localhost:8000` by default or whichever port defined in .env file
3. The front end application will start running on `http://localhost:3000` by default or whichever port defined in .env file

## API Endpoints
The following endpoints are available:

- `GET /blinks/list`: Retrieve all blinks
- `GET /blinks/list/{id}`: Retrieve a single blink by ID. ID should be a valid [Mongo ID](https://www.mongodb.com/docs/manual/reference/method/ObjectId/) Hex. 
- `POST /blinks/create/many`: Create multiple blinks when a valid array of blinks is sent as request's body.
- `POST /blinks/create/one`: Create one blink when a valid blink is sent as request's body.
- `PUT /blinks/update/{id}`: Update a specific blink by ID
- `DELETE /blinks/delete/{id}`: Delete a specific blink by ID

## Examples
Here are some examples of how to interact with the API using `curl`:

- Retrieve all blinks:
```shell
curl -X GET http://localhost:8000/blinks/list
```

- Retrieve a specific blink by ID:
```shell
curl -X GET http://localhost:8000/blinks/list/{id}
```

- Create a new blink:
```shell
curl -X POST -H "Content-Type: application/json" -d '{"title": "This is a sample blink title", "description": "This is a sample blink message"}' http://localhost:8000/blinks/create/one
```

- Update a specific blink by ID:
```shell
curl -X PUT -H "Content-Type: application/json" -d '{"title": "Updated title", "message": "Updated message"}' http://localhost:8000/blinks/update/{id}
```

- Delete a specific blink by ID:
```shell
curl -X DELETE http://localhost:8000/blinks/delete/{id}
```

## Contributing
Contributions to this project are welcome. To contribute, follow these steps:

1. Fork the repository.
2. Create a new branch.
3. Make your changes and commit them.
4. Push your changes to your forked repository.
5. Submit a pull request describing your changes.

## License
This project is licensed under the [MIT License](LICENSE).

