# 📖 DevBook – Social Network in Go

## Overview
**DevBook** is a full-stack social network application developed during the Go (Golang) Expert course. The project focuses on building a robust API following industry best practices and a dedicated Frontend to consume it. The application allows users to register, follow other users, create posts, and interact within a secure platform.

## Project Objectives
* Implement a **RESTful API** using the Go language.
* Handle authentication and authorization using **JWT (JSON Web Tokens)**.
* Manage data persistence with a relational database (**MySQL**).
* Develop a **Frontend** in Go using the `html/template` package.
* Apply security concepts such as password hashing with **Bcrypt**.
* Use a clean package architecture to ensure scalability and maintainability.

## Technologies Used
* **Go (Golang) 1.2x+**
* **Gorilla Mux** (HTTP Routing)
* **MySQL** (Relational Database)
* **JWT-go** (Authentication)
* **Bcrypt** (Password Encryption)
* **Godotenv** (Environment Variable Management)
* **HTML/Templates** (User Interface)

## Project Structure
The repository is organized into two main parts to separate backend logic from the user interface:

```text
DevBook/
│
├── api/                # Go REST API (Backend)
│   ├── src/            # Source code (Models, Repositories, Controllers)
│   ├── sql/            # Database creation and setup scripts
│   └── .env            # API environment configurations
│
├── web/                # Web Application (Frontend)
│   ├── src/            # Route logic and API consumption
│   ├── views/          # HTML Templates
│   ├── assets/         # CSS and JavaScript files
│   └── .env            # Frontend environment configurations
│
├── README.md           # Project documentation
└── main.go             # Entry point (based on repo organization)
