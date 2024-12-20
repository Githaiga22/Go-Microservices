# Hotel Booking System

A basic hotel booking system built using Go, showcasing microservice architecture. This system handles booking, customer, and room services, along with an API Gateway for routing requests to the appropriate microservices.

## Project Structure

The project is organized into a microservice architecture with the following structure:

```bash
hotel-booking/
├── main.go
├── services/
│   ├── booking/
│   │   ├── booking.go
│   │   ├── handler.go
│   │   └── routes.go
│   ├── customer/
│   │   ├── customer.go
│   │   ├── handler.go
│   │   └── routes.go
│   └── room/
│       ├── room.go
│       ├── handler.go
│       └── routes.go
├── api-gateway/
│   └── gateway.go
└── go.mod
