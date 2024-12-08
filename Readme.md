# Product Management System 
A backend system for a Product Management Application, this project highlights asynchronous image processing, caching, logging, and a scalable architecture. Developed using Golang, PostgreSQL, Redis, RabbitMQ, and S3, it prioritizes high performance and clean design.

---

## Features

1. **RESTful API Endpoints**:
   - `POST /products`: Create a new product with fields such as:
     - `user_id` (reference to the user table)
     - `product_name` (string)
     - `product_description` (text)
     - `product_images` (array of image URLs)
     - `product_price` (decimal)
   - `GET /products/:id`: Retrieve product details by ID, including processed image data.
   - `GET /products`: Fetch all products for a specific `user_id`, with optional filtering by price range and product name.

2. **Database**:
   - PostgreSQL is used for data storage.
   - Schema includes `users` and `products` tables with an additional column for `compressed_product_images`.

3. **Asynchronous Image Processing**:
   - Uses RabbitMQ or Kafka for message queuing.
   - Microservice for image processing downloads and compresses images, storing the processed images in S3 and updating the database.

4. **Caching**:
   - Redis is used to cache data for the `GET /products/:id` endpoint.
   - Cache invalidation ensures updates are reflected in real time.

5. **Logging**:
   - Structured logging with libraries like `logrus` or `zap`.
   - Logs include request details, response times, and error messages.

6. **Error Handling**:
   - Robust mechanisms for handling asynchronous processing failures with queue retry and dead-letter queues.

7. **Testing**:
   - Unit tests for API endpoints and core functions.
   - Integration tests for end-to-end functionality.
   - Benchmark tests for the `GET /products/:id` endpoint, with and without cache.

---

## System Requirements

1. **Languages and Tools**:
   - Golang (Backend API and microservices)
   - PostgreSQL (Database)
   - Redis (Caching)
   - RabbitMQ/Kafka (Message Queue)
   - S3 (Image Storage)

2. **Architecture**:
   - Modular structure for API, caching, and image processing services.
   - Scalable design to handle increased load and distributed services.
   - Ensures transactional consistency across all components.

---

## Setup Instructions

### Prerequisites

- Install Golang (v1.18 or later)
- Set up PostgreSQL and create the required database.
- Install Redis for caching.
- Set up RabbitMQ or Kafka for messaging.
- Configure S3 for image storage.

### Steps

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd product-management-system
   ```

2. Set up environment variables by creating a `.env` file:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=your_database
   REDIS_HOST=localhost
   REDIS_PORT=6379
   QUEUE_HOST=localhost
   QUEUE_PORT=5672
   S3_BUCKET_NAME=your_bucket_name
   S3_REGION=your_region
   S3_ACCESS_KEY=your_access_key
   S3_SECRET_KEY=your_secret_key
   ```

3. Run migrations to set up the database schema:
   ```bash
   go run scripts/migrate.go
   ```

4. Start the API server:
   ```bash
   go run main.go
   ```

5. Run the image processing microservice:
   ```bash
   go run services/image_processor/main.go
   ```

---

## Testing

1. Run unit tests:
   ```bash
   go test ./... -cover
   ```

2. Run integration tests:
   ```bash
   go test -tags=integration ./tests
   ```

3. Benchmark tests for caching:
   ```bash
   go test -bench=.
   ```

---

## Project Structure

```
|-- cmd/
|   |-- api/               # Main API server entry point
|-- config/                # Configuration files and environment handling
|-- db/                    # Database schema and queries
|-- internal/
|   |-- cache/             # Redis caching module
|   |-- logger/            # Logging module
|   |-- queue/             # RabbitMQ/Kafka integration
|-- services/
|   |-- image_processor/   # Asynchronous image processing service
|-- tests/                 # Unit and integration tests
|-- main.go                # API server bootstrap
|-- go.mod                 # Go module file
```

---

## Assumptions

- The system assumes valid image URLs during product creation.
- S3 bucket permissions are preconfigured for uploading and accessing images.
- Queue service is properly set up and running before starting the services.

---

## License

This project is licensed under the MIT License. See `LICENSE` for more details.
