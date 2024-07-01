# golang-stock-api

This repository is implementation of stock tracker by golang


## setup stock db

```shell
CREATE DATABASE stocksdb;
CREATE TABLE IF NOT EXISTS stocks (
  stockid SERIAL PRIMARY KEY,
  name TEXT,
  price INTEGER,
  company TEXT
);
```

## setup router with 

