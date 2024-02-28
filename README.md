# Rate Limiter Service

This repository contains a Golang service for rate limiting requests using the token bucket algorithm. The service is designed to efficiently manage incoming requests based on predefined rate limits, ensuring fair access to resources and preventing abuse.

## Features

- **Token Bucket Algorithm**: Implements the token bucket rate limiting algorithm to control the rate of incoming requests.
- **Middleware**: Utilizes middleware to extract headers from incoming requests and inject them into the context for traceability and identification purposes.
- **Graceful Crash Handling**: Includes recovery middleware to handle crashes gracefully, ensuring the service remains operational even in the face of unexpected errors.
- **Configurability**: Provides a configuration package to load service configuration parameters, enabling easy customization of rate limit settings and other service parameters.
- **Health Check Endpoint**: Includes handlers with a health check endpoint to monitor the service's operational status.

## Getting Started

To get started with the rate limiter service, follow these steps:

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/AhGhazey/rate_limiters

2. Access the health check endpoint to verify the service is running:

    ```bash
   curl --location 'localhost:9090/health' \
    --header 'Trace-ID: 00000000-0000-0000-0000-000000000001' \
    --header 'X-Forwarded-From: 203.128.17.42' \
    --header 'User-ID: 00000000-0000-0000-0000-000000000002'

## Middleware
The service utilizes middleware to extract headers from incoming requests and inject them into the context. The following headers are extracted:

- **Trace-ID**: Used for tracing requests across services.
- **X-Forwarded-From**: Represents the IP address of the client making the request.
- **User-ID**: Identifies the user associated with the request.

These headers are essential for traceability, identification, and logging purposes within the service.