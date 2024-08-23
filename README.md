# CRUD App

This repository contains a simple CRUD (Create, Read, Update, Delete) application. The application is containerized using Docker and is designed to be deployed on an AWS EC2 instance via a CI/CD pipeline.

## Overview

The architecture of the application is illustrated below:

![Architecture Diagram]
![Screenshot from 2024-08-23 21-37-25](https://github.com/user-attachments/assets/0863ec30-be97-4877-a554-347509e10450)



### Workflow:

1. **CRUD Application:** The core application providing CRUD functionalities.
2. **Docker:** 
   - The application is containerized by creating a Docker image.
   - The Docker image is pushed to Docker Hub.
3. **Docker Hub:** 
   - The Docker image is stored in a Docker Hub repository (`xyz/image:latest`).
4. **AWS EC2:**
   - The EC2 instance pulls the Docker image from Docker Hub.
   - The EC2 instance spins up a container using the pulled image.

## Prerequisites

Before running the project, ensure you have the following installed:

- Docker
- AWS CLI configured with your credentials
- An AWS EC2 instance running and accessible
- Docker Hub account

## Setup

### 1. Clone the Repository
```bash
git clone https://github.com/yourusername/yourrepository.git
cd yourrepository
