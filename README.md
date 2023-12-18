# Golang Fiber REST API

A simple RESTFul API built with Golang and [Fiber](https://gofiber.io).

## Table of Contents

- [Golang Fiber REST API](#golang-fiber-rest-api)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Requirements](#requirements)
  - [Getting Started](#getting-started)
    - [Installation](#installation)
    - [Before Starting](#before-starting)
    - [Running the Application](#running-the-application)
    - [Build Application](#build-application)
    - [API Documentation](#api-documentation)

## Introduction

This project is a simple REST API built using Golang and Fiber. It provides a basic structure for building RESTful services with Golang and includes essential features such as routing, middleware support, and more.

## Requirements

- Golang installed on your machine
- [Fiber](https://github.com/gofiber/fiber) library

## Getting Started

### Installation

Clone the repository:

```bash
git clone https://github.com/yadija/gofiber-rest-api.git
cd gofiber-rest-api
```

### Before Starting

Before running the application, you need to set up the necessary environment variables.

1. Open the `.env` file in the root directory of your project.

2. Add the following variables:

```bash
# MySQL Configuration
DB_PORT=3306
DB_HOST=localhost
DB_USER=your_mysql_username
DB_PASSWORD=your_mysql_password
DB_NAME=your_database_name

# TOKENIZE
ACCESS_TOKEN_KEY=your_access_token_secret_key
REFRESH_TOKEN_KEY=your_refresh_token_secret_key
```

3. Save the changes to the `.env` file.

### Running the Application

Run the following commands to start the server:

```bash
go run main.go
```

The server will start on `http://localhost:3000` by default.

### Build Application

Run the following command to build the application:

```bash
go build
```

### API Documentation

open `http://localhost:3000/swagger` in your browser to view the API documentation.