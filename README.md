# Distributed System Simulator

The Distributed System Simulator is a lightweight Go application designed to simulate a set of nodes in a distributed system. This project showcases key concepts in Go, such as concurrency with goroutines, thread-safe data access with mutexes, and HTTP server setup with basic RESTful endpoints.

## Key Features

- **HTTP Server**: The application creates an HTTP server with endpoints for retrieving node data and providing a welcome message.
- **Concurrency with Goroutines**: A goroutine periodically updates a random node to simulate a distributed system's behavior.
- **Thread-Safe Data Access**: The code uses `sync.RWMutex` to ensure thread-safe operations when accessing shared data.
- **Error Handling**: Proper error handling with appropriate HTTP status codes and log messages.
- **Unit Tests**: Comprehensive unit tests to validate the behavior of key functions, including `GetNodeData`, `RootHandler`, and `UpdateNode`.

<img width="1796" alt="Screenshot 2024-05-01 at 3 02 12 PM" src="https://github.com/shuddha2021/distributed-system-simulator-in-golang/assets/81951239/7e3703b9-33af-4fbe-af4c-ad82f5499e54">

<img width="937" alt="Screenshot 2024-05-01 at 3 02 32 PM" src="https://github.com/shuddha2021/distributed-system-simulator-in-golang/assets/81951239/a840c2c9-8cdf-44b8-85dd-24da5e5de4f4">

<img width="1174" alt="Screenshot 2024-05-01 at 3 02 52 PM" src="https://github.com/shuddha2021/distributed-system-simulator-in-golang/assets/81951239/cd0d0640-1743-48f0-8f66-cd8620d112f8">


## Core Logic

The core logic of the application revolves around simulating a set of nodes in a distributed system and providing HTTP endpoints to interact with them. Here's a brief overview of how it works:

- **Node Data Structure**: The `NodeData` struct represents a node with fields for `ID`, `Name`, `Value`, and `Time`.
- **Initialization**: The `InitNodes` function initializes a slice of nodes with random data.
- **HTTP Endpoints**: 
  - `/nodes`: Returns the current state of all nodes in JSON format.
  - `/`: Provides a welcome message with instructions for users.
- **Concurrency**: A goroutine periodically updates a random node's data every 5 seconds, demonstrating concurrency.
- **Synchronization**: The code uses `sync.RWMutex` to ensure thread-safe operations, and `sync.WaitGroup` to manage goroutine synchronization.

## Technologies Used

- **Go (Golang)**: The primary programming language used for building the application.
- **Net/HTTP**: Used to create the HTTP server and define HTTP handlers.
- **Sync Package**: Provides synchronization primitives such as `RWMutex` and `WaitGroup`.
- **Encoding/JSON**: For converting data to and from JSON format.
- **Testing Package**: Used to implement unit tests for critical functions.

## Getting Started

To run the Distributed System Simulator locally, follow these steps:

1. Clone the repository: `git clone <your-repo-url>`
2. Navigate to the project directory: `cd <your-project-directory>`
3. Start the application: `go run <your-go-file>.go`
4. Open your browser and visit `http://localhost:8080/` for the welcome message and `http://localhost:8080/nodes` to get the node data.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request. Be sure to follow the existing code style and write appropriate tests for any new functionality.

## License

This project is licensed under the MIT License. For more information, see the LICENSE file.

## Acknowledgments

- **Go Documentation**: A comprehensive resource for learning Go and its standard library.
- **Open-Source Community**: For providing resources and platforms that enable collaborative projects.
