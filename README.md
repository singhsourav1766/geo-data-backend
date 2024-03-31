# Geo Data Backend Project

## Setup

1. Create an `uploads` folder at the base directory to upload files there.
2. In `models/db.go`, at line 15, add your PostgreSQL DSN.

## Running the Project

1. Run the following commands:
    ```bash
    go mod tidy
    go run main.go
    ```

## Curl Commands

### User Authentication APIs

- Register a new user:
    ```bash
    curl -X POST http://localhost:8080/auth/register \
         -H "Content-Type: application/json" \
         -d '{"username": "user123", "password": "password123"}'
    ```

- Login (get JWT token):
    ```bash
    curl -X POST http://localhost:8080/auth/login \
         -H "Content-Type: application/json" \
         -d '{"username": "user123", "password": "password123"}'
    ```

- Logout:
    ```bash
    curl -X GET http://localhost:8080/auth/logout \
         -H "Authorization: Bearer YOUR_JWT_TOKEN"
    ```

### File Upload APIs

- Upload a file:
    ```bash
    curl -X POST http://localhost:8080/files/upload \
         -H "Authorization: Bearer YOUR_JWT_TOKEN" \
         -F "file=@/path/to/your/file.geojson"
    ```

- Get a file by ID:
    ```bash
    curl -X GET http://localhost:8080/files/1 \
         -H "Authorization: Bearer YOUR_JWT_TOKEN"
    ```

- Delete a file by ID:
    ```bash
    curl -X DELETE http://localhost:8080/files/1 \
         -H "Authorization: Bearer YOUR_JWT_TOKEN"
    ```

### Geospatial APIs

- Create geospatial data:
    ```bash
    curl -X POST http://localhost:8080/geospatial/create \
         -H "Content-Type: application/json" \
         -H "Authorization: Bearer YOUR_JWT_TOKEN" \
         -d '{"title": "GeoData 1", "content": "This is GeoData 1 content"}'
    ```

- Get geospatial data by ID:
    ```bash
    curl -X GET http://localhost:8080/geospatial/1 \
         -H "Authorization: Bearer YOUR_JWT_TOKEN"
    ```

- Update geospatial data by ID:
    ```bash
    curl -X PUT http://localhost:8080/geospatial/1 \
         -H "Content-Type: application/json" \
         -H "Authorization: Bearer YOUR_JWT_TOKEN" \
         -d '{"title": "Updated GeoData 1", "content": "Updated content"}'
    ```

- Delete geospatial data by ID:
    ```bash
    curl -X DELETE http://localhost:8080/geospatial/1 \
         -H "Authorization: Bearer YOUR_JWT_TOKEN"
    ```

![Screenshot (619)](https://github.com/singhsourav1766/geo-data-backend/assets/121311381/2d17b96b-276e-48b5-a94b-9a3a1d4dbd3e)
![Screenshot (618)](https://github.com/singhsourav1766/geo-data-backend/assets/121311381/f1a00421-9d1e-4899-8101-e9b5b2634705)
![Screenshot (623)](https://github.com/singhsourav1766/geo-data-backend/assets/121311381/db47b6dd-bb85-4cfb-a417-461e0138b698)
![Screenshot (622)](https://github.com/singhsourav1766/geo-data-backend/assets/121311381/746c0e7a-6727-412d-a352-4ca9edea92a7)
![Screenshot (621)](https://github.com/singhsourav1766/geo-data-backend/assets/121311381/c7856c41-7feb-4661-8370-6e72cd5773c8)
![Screenshot (620)](https://github.com/singhsourav1766/geo-data-backend/assets/121311381/e66c5dc9-c421-4578-a000-9376c4a34c39)


