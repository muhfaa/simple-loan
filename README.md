# Simple Loan

Simple Loan is a simple CRUD for the loan system..

In this application I use the Hexagonal architecture with Golang as the programming language, MySQL and MongoDB as the database


# API
[Simple Loan API](https://github.com/muhfaa/Game-Currency/blob/main/index.md)

# SQL Query
[SQl query to create table]()

# Postman collaction
[Postman](https://github.com/muhfaa/Game-Currency/blob/1482240a72326f4c2c778287a37e0d9513aabf61/ATTN.postman_collection.json)

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
- Next execute this [SQl query to create table]() in that SQL query
