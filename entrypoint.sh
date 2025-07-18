#!/bin/sh

echo "⏳ Waiting for PostgreSQL..."
until nc -z postgres 5432; do
  sleep 1
done
echo "✅ PostgreSQL is up!"

echo "📦 Running db migrate deploy..."
make migrate-sql

echo "🌱 Seeding data..."
make run-data-seed

echo "🚀 Starting the app..."
/golang-auth-app