# Pelter-Backend

## Overview

Pelter-Backend is the backend service for the Pelter application. It provides APIs for managing user data, authentication, and other core functionalities.

## Features

- Product CRUD
- User manangement
- Transaction management

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/pelter-backend.git
   ```
2. Navigate to the project directory:
   ```sh
   cd pelter-backend
   ```
3. Install dependencies:
   ```sh
   go install
   ```

## Usage

1. Start the development server:
   ```sh
   docker build .
   docker compose up
   make run
   ```
2. Access the API at `http://localhost:8080`

## Configuration

Create a `.env` file in the root directory and add the following environment variables in .env.sample

## API Documentation

For detailed API documentation, refer to the [API Docs](https://documenter.getpostman.com/view/39871792/2sAYBSiYV7).
