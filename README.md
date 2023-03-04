# crm-backend-project

Project for Udacity's Golang course

This project implements a RESTful CRUD API server for a CRM (Customer Relationship Management) system.

# Usage

## Running the API server
To run the server, run:
```
go run main.go
```

## Running Unit tests
To run the unit tests, run:
```
go test
```

# API Endpoints

## GET /customers

Retrieve all customer data.

## GET /customers/{id}

Retrive data for a customer

## POST /customers

Create a new customer

Request (JSON) Body fields:
* __Name__: (required, string): Customer's name
* __Role__: (required, string): Customer's role
* __Email__: (required, string): Customer's email address
* __Phone__: (required, number): Customer's phone number 
* __Contacted__: (required, boolean): True if customer has been contacted

## UPDATE /customers/{id}

Update an exiting customer's data

Request (JSON) Body fields:
* __Name__: (required, string): Customer's name
* __Role__: (required, string): Customer's role
* __Email__: (required, string): Customer's email address
* __Phone__: (required, number): Customer's phone number 
* __Contacted__: (required, boolean): True if customer has been contacted


## DELETE /customers/{id}

Deletes a customer 