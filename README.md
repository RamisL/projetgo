# projetgo

A Go project for creating and managing payments and products.

## Table of Contents

- [Installation](#installation)
- [Usage](#Usage)
- [Support](#Support)
- [Contributing](#Contributing)

## Installation

To run the project project, you will need to have Docker and Docker Compose installed on your machine.

First, clone the repository to your local machine using Git:

`git clone https://github.com/RamisL/projetgo.git`

Change into the project directory:

`cd projetgo`

Build the Docker images for the project using the docker-compose command:

`docker-compose build`

Run the project using the docker-compose command:

`docker-compose up`

This will start the containers for the project and run the application. You should see the logs for the application in the terminal.

To stop the containers, press **CTRL+C** in the terminal.

To start the containers again, use the **docker-compose up** command.

Then to install the project package, run the following command:

`go get github.com/RamisL/projetgo`

This will download and install the package and its dependencies into your Go workspace.

## Usage

To use the project project with Postman, you will need to have Go and Postman installed on your machine.

Build and run the project using the go run command:

`go run main.go`

This will start the server for the project and print the logs to the terminal.

Here is an example of how to create a new product:

Open Postman and create a new request. Set the method to POST and the URL to `http://localhost:3000/api/product/create`.

In the body of the request, enter a JSON object with the input fields for the product. For example:

```
{
    "name": "Desk",
    "price": "199.99",
}
```

Send the request by clicking the Send button

## Support

If you have any questions or issues with the project package, please file an issue in the project repository.

## Contributing

To contribute to the project package, fork the repository and create a pull request with your changes.

Thank you for considering contributing to project!
