# Config Resolution Service

A simple Go service that resolves environment variable values via a REST API endpoint.

## Getting Started

### Build and Run

```bash
go build -o config-resolution
./config-resolution
```

The service will start on port 9090 and display:
```
2025/10/28 21:09:13 Started service on port: :9090
```

## API Endpoint

### POST /test/resolve/values

Resolves environment variable values from the system.

#### Request Format

```json
{
  "resolve": ["HOME", "USER", "PATH"]
}
```

The `resolve` field accepts an array of environment variable names to look up.

#### Response Format

```json
{
  "resolvedValues": {
    "HOME": "/Users/hasithah",
    "USER": "hasithah",
    "PATH": "/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin"
  }
}
```

The response contains a map where:
- **Keys**: Environment variable names from the request
- **Values**: Resolved values from the system (empty string if not found)

## Example Usage

```bash
curl -X POST http://localhost:9090/test/resolve/values \
  -H "Content-Type: application/json" \
  -d '{"resolve": ["HOME", "USER"]}'
```

### Sample Output

```json
{
  "resolvedValues": {
    "HOME": "/Users/hasithah",
    "USER": "hasithah"
  }
}
```

## Logging

The service logs each resolved variable:

```
2025/10/28 21:09:33 Resolved HOME = /Users/hasithah
2025/10/28 21:09:33 Resolved USER = hasithah
```
