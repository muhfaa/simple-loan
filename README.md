# Simple Loan

Simple Loan is a simple CRUD for the loan system..

In this application I use the Hexagonal architecture with Golang as the programming language, MySQL and MongoDB as the database


# API
[Simple Loan API](https://github.com/muhfaa/simple-loan/blob/main/index.md)

# SQL Query
[SQl query to create table](https://github.com/muhfaa/simple-loan/blob/main/loan-database-init.sql)

# Postman collaction
[Postman](https://github.com/muhfaa/simple-loan/blob/main/Loan.postman_collection.json)

## Requirement

- go 1.16+
- Mysql
- Docker

## Config

For config file i use `json`

## Running App

`docker network create -d bridge my-project`

`docker-compose up --build`

- After that open [url](http://localhost:8090/?server=golang-loan-db&username=root&db=loan-db) and `password: root`
- Next execute this [SQl query to create table](https://github.com/muhfaa/simple-loan/blob/main/loan-database-init.sql) in that SQL query
