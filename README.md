# Mahasiswa JobHun

## Name

Muhammad Faris Hadi Mulyo - Backend Developer - Technical Test Jobhun Internship 2023

## Description
Repository yang digunakan untuk techincal test  Jobhun Internship 2023

## Requirement & Stack

-   Golang >=1.17
-   MySQL 5.7

## Local Installation

1. Clone this project
    ```
    git@gitlab.dot.co.id:dpr/dpr-sniper-be.git
    ```

2. Copy `.env.example` to `.env`
    ```
    cp .env.example .env
    ```
3. Configure environment variables for database connection
    ```
    DB_CONNECTION=mysql
    DB_HOST=127.0.0.1
    DB_PORT=3306
    DB_DATABASE=jobhun_mahasiswa
    DB_USERNAME=root
    DB_PASSWORD=
    ```

4.  Run the application
    ```
    go run main.go
    ```