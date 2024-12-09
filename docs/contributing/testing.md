---
layout: page
title: Testing
permalink: /contributing/testing/
parent: Contributing
---

# Testing Guide

## Test Categories

### 1. Unit Tests

- Test individual components
- Mock dependencies
- Fast execution

### 2. Integration Tests

- Test component interaction
- Use test database
- Slower execution

### 3. End-to-End Tests

- Test complete workflows
- Use real PostgreSQL
- Slowest execution

## Writing Tests

### Test Structure

```go
func TestFeature(t *testing.T) {
    // Setup
    ctx := context.Background()
    db := setupTestDB(t)
    defer cleanupDB(t, db)

    // Test cases
    tests := []struct {
        name    string
        input   string
        want    *Result
        wantErr bool
    }{
        {
            name:    "simple select",
            input:   "SELECT * FROM users",
            want:    expectedResult,
            wantErr: false,
        },
    }

    // Execute tests
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Feature(ctx, tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("unexpected error: %v", err)
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}
```

## Test Utilities

### Database Setup

```go
func setupTestDB(t *testing.T) *sql.DB {
    t.Helper()
    // Setup code
}

func cleanupDB(t *testing.T, db *sql.DB) {
    t.Helper()
    // Cleanup code
}
```

## Running Tests

### All Tests

```bash
go test ./...
```

### Specific Package

```bash
go test ./internal/parser
```

### With Coverage

```bash
go test -cover ./...
```

### Generate Coverage Report

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Test Data

1. **Fixtures**

   - Place in testdata/ directory
   - Use meaningful names
   - Include documentation

2. **Generated Data**
   - Use consistent seeds
   - Document generation process
   - Handle edge cases

## Mocking

1. **Interface Mocking**

   ```go
   type MockParser struct {
       mock.Mock
   }
   ```

2. **Database Mocking**
   - Use sqlmock for database tests
   - Document mock expectations
   - Clean up resources

## CI Integration

1. **GitHub Actions**

   - Run on pull requests
   - Include all test categories
   - Generate coverage reports

2. **Local Verification**
   ```bash
   make test      # Run all tests
   make test-ci   # Run CI test suite
   make cover     # Generate coverage
   ```

## Best Practices

1. **Test Naming**

   - Clear and descriptive
   - Indicate what's being tested
   - Show expected behavior

2. **Test Size**

   - Keep tests focused
   - One assertion per test
   - Clear setup and teardown

3. **Test Coverage**

   - Aim for >80% coverage
   - Focus on critical paths
   - Document uncovered cases

4. **Test Maintenance**
   - Update tests with code changes
   - Remove obsolete tests
   - Keep tests readable

````
## Test Utilities

### Database Setup
```go
func setupTestDB(t *testing.T) *sql.DB {
    t.Helper()
    // Setup code
}

func cleanupDB(t *testing.T, db *sql.DB) {
    t.Helper()
    // Cleanup code
}
````

## Running Tests

### All Tests

```bash
go test ./...
```

### Specific Package

```bash
go test ./internal/parser
```

### With Coverage

```bash
go test -cover ./...
```

### Generate Coverage Report

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Test Data

1. **Fixtures**

   - Place in testdata/ directory
   - Use meaningful names
   - Include documentation

2. **Generated Data**
   - Use consistent seeds
   - Document generation process
   - Handle edge cases

## Mocking

1. **Interface Mocking**

   ```go
   type MockParser struct {
       mock.Mock
   }
   ```

2. **Database Mocking**
   - Use sqlmock for database tests
   - Document mock expectations
   - Clean up resources

## CI Integration

1. **GitHub Actions**

   - Run on pull requests
   - Include all test categories
   - Generate coverage reports

2. **Local Verification**
   ```bash
   make test      # Run all tests
   make test-ci   # Run CI test suite
   make cover     # Generate coverage
   ```

## Best Practices

1. **Test Naming**

   - Clear and descriptive
   - Indicate what's being tested
   - Show expected behavior

2. **Test Size**

   - Keep tests focused
   - One assertion per test
   - Clear setup and teardown

3. **Test Coverage**

   - Aim for >80% coverage
   - Focus on critical paths
   - Document uncovered cases

4. **Test Maintenance**
   - Update tests with code changes
   - Remove obsolete tests
   - Keep tests readable

```

```
