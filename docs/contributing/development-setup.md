---
layout: page
title: Development Setup
permalink: /contributing/development-setup/
parent: Contributing
---

# Development Setup

This guide will help you set up your development environment for PGQL.

## Prerequisites

- Go 1.21 or later
- PostgreSQL 12 or later
- Git

## Setting Up Your Development Environment

1. **Fork and Clone**

   ```bash
   git clone https://github.com/Kinveil/pgql.git
   cd pgql
   ```

2. **Install Dependencies**

   ```bash
   go mod download
   ```

3. **Set Up PostgreSQL**

   - Install PostgreSQL
   - Create a test database:
     ```sql
     CREATE DATABASE pgql_test;
     ```

4. **Environment Setup**

   ```bash
   # Linux/Mac
   export PGQL_TEST_DB="postgres://user:pass@localhost:5432/pgql_test"

   # Windows
   set PGQL_TEST_DB=postgres://user:pass@localhost:5432/pgql_test
   ```

## Development Workflow

1. **Create a Branch**

   ```bash
   git checkout -b feature/your-feature
   ```

2. **Run Tests**

   ```bash
   go test ./...
   ```

3. **Generate Code**

   ```bash
   go generate ./...
   ```

4. **Run Linters**
   ```bash
   make lint
   ```

## Common Issues

1. **PostgreSQL Connection**

   - Check port number
   - Verify credentials
   - Ensure database exists

2. **Go Module Issues**
   ```bash
   go mod tidy
   ```

## Need Help?

- Check existing issues
- Join our discussions
- Ask in pull requests
