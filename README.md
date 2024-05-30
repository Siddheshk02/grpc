# Golang gRPC User Service

This project implements a Golang gRPC service for managing user details with functionalities to fetch user details by ID, retrieve a list of user details by IDs, and search for user details based on specific criteria.

## Features

- Fetch user details by ID
- Retrieve a list of user details by IDs
- Search for user details based on specific criteria (e.g., first name, city, phone number, marital status)
- Dockerized for easy deployment

## Prerequisites

- Docker

## Getting Started

### Pulling the Docker Image

To run the gRPC service, you can pull the pre-built Docker image from Docker Hub.

```bash
docker pull Siddheshk02/grpc:latest

Running the Docker Container

```
docker run -p 50051:50051 Siddheshk02/grpc:latest


You can use grpcurl to test the gRPC endpoints:

1. GetUser:
```
grpcurl -plaintext -d '{"id": 1}' localhost:50051 pb.UserService/GetUser

2. ListUsers:
```
grpcurl -plaintext -d '{"ids": [1, 2]}' localhost:50051 pb.UserService/ListUsers

3. SearchUser:
```
grpcurl -plaintext -d '{"fname": "Mark", "city": "CA"}' localhost:50051 pb.UserService/SearchUsers

