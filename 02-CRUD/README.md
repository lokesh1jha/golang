# CRUD Operations with Gorilla Mux in Go

This project demonstrates how to implement CRUD (Create, Read, Update, Delete) operations using the Gorilla Mux router in Go.

## Prerequisites

- Go 1.16 or higher
- Gorilla Mux package

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/lokesh1jha/02-CRUD.git
    ```
2. Navigate to the project directory:
    ```sh
    cd 02-CRUD
    ```
3. Install the dependencies:
    ```sh
    go get -u github.com/gorilla/mux
    ```

## Running the Application

To run the application, use the following command:
```sh
go run main.go
```

## API Endpoints

- **Create**: `POST /movies`
- **Read**: `GET /movies`
- **Read One**: `GET /movies/{id}`
- **Update**: `PUT /movies/{id}`
- **Delete**: `DELETE /movies/{id}`


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.