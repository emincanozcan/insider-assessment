# Insider Assesment Project

This project implements an automatic message sending system based on the given requirements.

## Getting Started

### Prerequisites

- Docker
- Git

### Installation

1. Clone the repository:
   ```bash
   git clone github.com/emincanozcan/insider-assessment
   cd insider-assessment
   ```

2. Start the services using docker compose:
    ```bash
    docker compose up
    ```
This will start all the necessary services, including a custom webhook.site imitation server.

### Configuration

Configuration parameters can be updated within the environment section of the `docker-compose.yml` file.

- `MESSAGE_SEND_INTERVAL`   Message sending interval in seconds (e.g., 120 to send messages every 2 minutes).
- `MESSAGE_SEND_BATCH_SIZE` Number of messages in each sending batch (e.g., 2 to send 2 messages in every interval).
- `WEBHOOK_URL`             URL to send the messages. By default, this is set to the webhook imitation server. It can also be set to an external URL like `https://webhook.site/bla-bla`.
- `WEBHOOK_AUTH_KEY`        Value of the `x-ins-auth-key` header.

### API Endpoints

The project includes the following API endpoints:

1. `GET /swagger/index.html#`

   Opens the Swagger UI. It's suggested to use this endpoint in a browser. The endpoints below [2-5] can be tried out through this interface.

2. `GET /messages/sent`

   Returns a list of all sent messages.

3. `POST /messages`

   Adds a new message to the system. Body parameters are recipient and content with max character limit validations.

4. `POST /messages/processing/start`

   Starts the background service that sends unsent messages every 2 minutes.

5. `POST /messages/processing/stop`

   Stops the background message sending service.


## Technologies Used

1. Programming Language: Go

2. Database: PostgreSQL

3. Cache: Redis

4. Database query generator: [sqlc](https://github.com/sqlc-dev/sqlc), specifically [sqlc-gen-go](https://github.com/sqlc-dev/sqlc-gen-go)

5. Database migrataion manager: [golang-migrate](https://github.com/golang-migrate/migrate)

## Troubleshooting

### Port Issues

By default, the docker compose file exposes:

- port 8080 for the assessment project
- port 8081 for the webhook.site imitation server
- port 5432 for postgresql
- port 6379 for redis

If you encounter any issues due to unavailable ports, feel free to customize these ports.

Service-to-service communication is done via Docker networking, so the ports are opened to the outside only for testing purposes.
