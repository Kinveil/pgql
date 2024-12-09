---
layout: page
title: Code Style
permalink: /contributing/code-style/
parent: Contributing
---

# Code Style Guide

## General Principles

1. **Readability First**

   - Clear over clever
   - Self-documenting code
   - Meaningful variable names

2. **Consistency**
   - Follow Go conventions
   - Use standard project layouts
   - Maintain existing patterns

## Go Style Guidelines

### Naming

1. **Packages**

   - Short, concise names
   - No plurals
   - No underscores

   ```go
   // Good
   package parser
   package analyzer

   // Bad
   package sqlParser
   package sql_analyzer
   ```

2. **Variables**

   - Descriptive but concise
   - Avoid abbreviations except for common ones

   ```go
   // Good
   var queryParser *Parser
   var tableNames []string

   // Bad
   var qp *Parser
   var tn []string
   ```

### Code Organization

1. **File Structure**

   ```go
   // Imports grouped by standard, external, internal
   import (
       "context"
       "strings"

       "github.com/lib/pq"

       "github.com/your/project/internal/parser"
   )
   ```

2. **Type Definitions**

   ```go
   // Group related types together
   type Query struct {
       // ...
   }

   type QueryOption func(*Query)
   ```

### Error Handling

1. **Error Types**

   ```go
   // Define custom errors
   type ParseError struct {
       Line    int
       Message string
   }
   ```

2. **Error Wrapping**
   ```go
   if err != nil {
       return fmt.Errorf("parsing query: %w", err)
   }
   ```

## SQL Style Guidelines

1. **Query Formatting**

   ```sql
   SELECT
       user_id,
       full_name
   FROM users
   WHERE status = $1
   ```

2. **Comments**
   ```sql
   -- Explain complex queries
   WITH active_users AS (
       SELECT user_id
       FROM users
       WHERE last_active > now() - interval '30 days'
   )
   ```

## Documentation

1. **Package Documentation**

   ```go
   // Package parser provides SQL parsing functionality for PGQL.
   package parser
   ```

2. **Function Documentation**
   ```go
   // ParseQuery parses a SQL query and returns its AST representation.
   // It returns an error if the query is invalid.
   func ParseQuery(query string) (*AST, error)
   ```

## Testing

1. **Test Names**

   ```go
   func TestParser_SimpleSelect(t *testing.T)
   func TestParser_ComplexJoin(t *testing.T)
   ```

2. **Table Tests**
   ```go
   func TestParse(t *testing.T) {
       tests := []struct {
           name    string
           input   string
           want    *AST
           wantErr bool
       }{
           // test cases
       }
       // ...
   }
   ```
