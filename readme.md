# RSS Post Aggregator

Small Go service that collects and serves RSS posts via HTTP routes.

## Build & Run

1. Build:
    go build -o rss-aggregator ./cmd/... || go build -o rss-aggregator main.go

2. Run:
    ./rss-aggregator

## Configuration (environment variables)

- PORT (default: 8080)  
- DATABASE_URL  
- RSS_POLL_INTERVAL (seconds or duration string)  
- LOG_LEVEL (e.g., debug, info, warn, error)

## Routes / Endpoints

- GET  /health               — health check
- GET  /feeds                — list registered RSS feeds
- POST /feeds                — add a new feed (JSON payload: feed URL, optional metadata)
- GET  /feeds/{id}           — get feed details
- DELETE /feeds/{id}         — remove a feed
- GET  /posts                — list aggregated posts (supports pagination & filters like feed_id, since)
- GET  /posts/{id}           — get a single post by ID
- POST /refresh              — trigger manual refresh of all feeds

Note: Replace or extend the above endpoints with the exact handlers found in main.go if they differ.


