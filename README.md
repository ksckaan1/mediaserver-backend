# MediaServer

Backend for Movie and Series Management

## Services
- API Documentation http://localhost:9189
- BFF (API Gateway - Backend For Frontend) http://localhost:9190
    - Movie
    - Series
    - Season
    - Episode
    - Media
    - TMDB
- MinIO http://localhost:9001
- Couchbase http://localhost:8091
- Jaeger http://localhost:16686

## Environment Variables

Create `.env` file in the root directory.

```env
# COUCHBASE
COUCHBASE_HOST=couchbase
COUCHBASE_USER=cbadmin
COUCHBASE_PASSWORD=cbpass
COUCHBASE_BUCKET=media_server

# SERVICES
MEDIA_SERVICE_ADDR=media_service:8080
TMDB_SERVICE_ADDR=tmdb_service:8080
SERIES_SERVICE_ADDR=series_service:8080
EPISODE_SERVICE_ADDR=episode_service:8080
SEASON_SERVICE_ADDR=season_service:8080
MOVIE_SERVICE_ADDR=movie_service:8080


# API KEYS
TMDB_API_KEY=<TMDB API KEY>

# TRACER
TRACER_ENDPOINT=jaeger:4318

# S3
S3_ENDPOINT=minio:9000
S3_REGION=eu-central-1
S3_BUCKET=media
S3_ACCESS_KEY=minioadmin
S3_SECRET_KEY=minioadmin
S3_USE_SSL=false
```

> [!IMPORTANT]  
> Replace `<TMDB API KEY>` with your actual TMDB API key.

## Run

```bash
docker compose up
```
	