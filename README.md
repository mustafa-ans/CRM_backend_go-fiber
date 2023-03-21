# CRM Backend

This is a backend server for a Customer Relationship Management (CRM) system, written in Golang using the Go Fiber library.
Installation

    Clone the repository to your local machine:

bash

git clone https://github.com/your_username/your_repo.git

    Install the necessary dependencies using go modules:

go

go mod download

    Run the server:

go

go run main.go

The server should now be running on http://localhost:3000.
Endpoints

The following endpoints are available:
Customers

    GET /customers - Get all customers
    GET /customers/:id - Get a specific customer by ID
    POST /customers - Create a new customer
    PUT /customers/:id - Update an existing customer by ID
    DELETE /customers/:id - Delete a customer by ID


# Contributing

Contributions are welcome! If you have any feature requests or bug reports, please submit them as issues on the repository.
License

This project is licensed under the MIT License. See the LICENSE file for more information.
