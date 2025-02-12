# Meme Coin API
Develop APIs for managing meme coins and containerize the application for
deployment.

![Swagger Presentation](assets/swagger_presentation.png)

## Prerequisites
- **Go**: Version 1.20+
- **Docker**: Version 20.10+
- **Docker Compose**: Version 2.20+
- optional for local setup
  - **MySQL**
  - **Redis**
  - **Swaggo**

## Local Setup
1. Clone the repository and navigate to the project directory:
    ```
    git clone https://github.com/wwieo/meme_coin_api.git
    ```
2. To set up locally, install the `swaggo/swag` by following the instructions on
   the [website](https://github.com/swaggo/swag):
    ```
    go install github.com/swaggo/swag/cmd/swag@latest
    ```
3. Modify the [local configuration file](config/local.json) if necessary.
4. Ensure that both Redis and MySQL are running with your configuration settings.
5. This project does not support automatic database migration locally. Please create a database manually in MySQL.
    ```
    CREATE DATABASE IF NOT EXISTS meme_coin;
    ```
6. Start the application, and you should see the output below.
    ```
    make up
   
    # Output Like:
    2025/02/12 18:08:42 Pinged successfully mysql database: meme_coin
    2025/02/12 18:08:42 Pinged successfully redis: meme_coin
    Meme Coin API starts at :8000
    ```
7. Then you can open the local Swagger interface to explore the API.
    ```
    http://localhost:8000/meme_coin_api/swagger/index.html
    ```
   
## Running with Docker
1. Build
   ```
   make docker-build
   ```
2. Start
   ```
   make docker-up
   ```
3. Then you can open the local Swagger interface to explore the API.
   ```
   http://localhost:8000/meme_coin_api/swagger/index.html
   ```
