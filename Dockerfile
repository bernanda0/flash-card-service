# Use the official PostgreSQL image as the base image
FROM postgres:latest

# Set environment variables for the PostgreSQL image
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD postgres
ENV POSTGRES_DB spaced_repetition

