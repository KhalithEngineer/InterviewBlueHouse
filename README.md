## Go HTTP Server with Gin
This repository contains a basic Go application featuring an HTTP server using the Gin web framework. The server listens for POST requests on the "/helloworld" endpoint and simulates a delayed response with the received request body.

# Getting Started

Run the server:

go run main.go
The server will start at http://localhost:8080 by default.

Make a POST request to http://localhost:8080/helloworld with a message in the request body, and the server will respond after a random delay.

# Example

curl -X POST -d "Hello, World!" http://localhost:8080/helloworld
Implementation Details
returnRequestBody function: This function reads the request body, simulates a random delay, and then responds with the received request body.

Server Configuration: The Gin web framework is used to create a default server instance, which listens for POST requests on the "/helloworld" endpoint.

## Concurrent HTTP Requests and Response Time Analysis
This Go application demonstrates concurrent sending of HTTP POST requests to a specified URL and analyzes the response time of each request. The program calculates and prints the minimum and maximum time taken for the requests.

# Table of Contents
Overview
How it Works
Usage
Implementation Details
Contributing
License
Overview
This application utilizes goroutines and channels in Go to concurrently send multiple HTTP POST requests to a specified endpoint (http://localhost:8080/helloworld in this example). It records the time taken for each request and calculates the minimum and maximum response times.

# How it Works
hitRequest function: This function sends a POST request to the specified URL, records the elapsed time, and communicates the time duration through a channel.

findMinMaxTime function: This function reads from the time duration channel and finds the minimum and maximum elapsed times.

Concurrency: The program launches multiple goroutines concurrently to simulate concurrent HTTP requests. It utilizes a wait group to ensure that all goroutines have completed before calculating the minimum and maximum response times.
