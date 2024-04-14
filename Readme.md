## Introduction

This project implements APIs for cardholder and transaction management. 
It is built using golang and grpc. The APIs are defined using protobuf and the server is implemented using gRPC. 
The APIs are exposed using REST.


## Project Structure

```bash
├── bin
│   ├── api
│   └── migration
├── cmd
│   ├── api
│   │   └── main.go
│   └── migration
│       └── main.go
├── deployment
│   ├── api.Dockerfile
│   └── migration.Dockerfile
├── docker-compose.yml
├── docs
│   └── openapiv2
│       └── protos
│           └── cardholder.swagger.json
│           └── transaction.swagger.json
│
├── internal
│   ├── app
│   │   ├── cardholder
│   │   └── transaction
│   └── mocks
│       ├── cardholder
│       └── transaction
├── Makefile
├── protos
├── README.md
└── rpc
```
The folder description for above structure is as follows:


| File/Folder | Description                                                                                                       |
|---|-------------------------------------------------------------------------------------------------------------------|
| bin | It contains the compiled binaries for this app                                                                    |
| cmd | It contains the main applications for this app. Each application has its own subdirectory named after it.         |
| deployment | It contains Dockerfiles.                                                                                          |
| docs | It contains the OpenAPI specification for the APIs (accounts and transactions) generated from the protobuf files. |
| internal | It contains the application code which is divided into different modules like app,config,db and mocks.            |
| protos | It contains the Protobuf definitions for the APIs (accounts and transactions).                                    |
| rpc | It contains the generated code from the Protobuf definitions.                                                     |

## Pre-requisites

- Docker
- Docker-compose

## How to setup (docker)


Build the docker images using the docker-compose file
```bash
docker-compose build
```

Run the docker-compose setup

```bash
docker-compose up
```

## How to generate OpenAPI spec
You can generate the open api spec for the existing apis using the following:

```bash
# This downloads the dependencies and generates openapi spec
make setup
```
It will be generated under `docs/openapiv2/protos` folder.


## How to use the APIs

The APIs can be accessed using the following postman link:

Postman link: [link](https://api.postman.com/collections/32800026-99a86293-2332-4d3c-8fb0-dad89b1c91b1?access_key=PMAT-01HVEWF2YBWY3597KGQ4XEJFRD)


## How to run the tests

```bash
 docker-compose -f docker-compose-testing.yml up
```