# High-Load TCP Service for Game Clients

## English Version

For English-speaking users, documentation can be found [here](https://github.com/ascenmmo/tcp-server/blob/master/README.md).

## Description

This project is a high-load **TCP service** for connecting game clients. It is designed for free deployment on servers, making it an optimal solution for game developers.

## Key Features

- **High Performance**: Optimized for handling a large number of simultaneous TCP connections.
- **Free Deployment**: Easily deployable on any server without additional costs.
- **Docker Support**: Available as both a binary and a Docker container, ensuring compatibility across platforms.
- **Flexibility and Scalability**: Configurable and easily scalable to meet your needs.

## Installation

### Installation via Docker

1. Make sure Docker is installed. If it’s not installed, follow the Docker installation instructions.
2. Run the command:
   ```bash
   docker compose up -d --force-recreate --build && docker image prune -f
	```

## Configuration

The project uses the env/env.go package to define key configuration parameters. All services interacting with this TCP server must use the same token to ensure secure connections and authentication.

### Configuration Parameters

```go
package env

var (
   ServerAddress       = "0.0.0.0" // Server IP address
   TCPPort             = "8083"    // Port for TCP connections
   TokenKey            = "_remember_token_must_be_32_bytes" // Unique token for authentication
   MaxRequestPerSecond = 5         // Max requests per second
)
```

* ServerAddress: Specifies the IP address where the server will operate.
* TCPPort: The port on which the server will listen for TCP connections.
* TokenKey: A unique token that must be the same across all services interacting with this TCP server. This ensures the security and integrity of connections.
* MaxRequestPerSecond: A limit on the maximum number of requests the server can handle per second.



##  Importance of a Single Token
### Using a single token allows:

* **Security Assurance:** All services check the token before establishing a connection, helping to prevent unauthorized access.
* **Simplified Authentication:** A single token for all services simplifies the authentication and access management process.
* **Ease of Maintenance:** If the token needs to be changed, it can be done in one place, and all services will be updated simultaneously.

Make sure that all your services are configured to use this token to ensure the correct operation and security of the system.



## Troubleshooting

If you encounter issues when starting the service, check the following:

- Ensure that port 8083 is not in use by other applications.
- Verify that the configuration parameters in env/env.go are correctly specified.
- If you are using Docker, ensure that it is running and properly configured.






## Теги

`TCP`, `игровой сервер`, `высоконагруженный`, `бесплатное развертывание`, `Docker`, `кроссплатформенный`, `игровая разработка`, `сеть`, `многопользовательская игра`, `сервис для игр`, `настройка сервера`, `аутентификация`, `токены`, `Golang`, `open-source`

## Tags

`TCP`, `game server`, `high-performance`, `free deployment`, `Docker`, `cross-platform`, `game development`, `network`, `multiplayer game`, `game service`, `server setup`, `authentication`, `tokens`, `Golang`, `open-source`
