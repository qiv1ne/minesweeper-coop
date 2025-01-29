# Minesweeper

This project is aim to create minesweeper game for cooperative play.

Project contains multiple services:

1. Minesweeper service - main service contain Minesweeper logic.
2. Auth service - do auth and verification jobs.
4. Frontend service - main service for graphic user interface.
4. Postgres service - provide interaction with PostgreSQL database.
5. Redis service - provide interaction with Redis database for store game and user sessions.
6. PostgreSQL - database for storing user profiles, game history, leaderboards.
7. Redis - database for storing user and game sessions.

All services use REST endpoints for communication

## Frontend service
Listen on port *8080*

## Minesweeper service
Listen on port *8081*

## Auth service
Listen on port *8082*

## Redis service
Listen on port *8083*

## Postgres service
Listen on port *8084*
