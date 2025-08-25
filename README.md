# URL Shortener API
I just developed for fun (to learn concepts of web servers in golang) not tested and due to no authorization or proper seeding it's not production ready.
## Usage Examples

### Basic request:
```bash
curl -X POST http://127.0.0.1:8080/url/ \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com", "ttl": 10}'
```

### With different TTL:
```bash
curl -X POST http://127.0.0.1:8080/url/ \
  -H "Content-Type: application/json" \
  -d '{"url": "https://google.com", "ttl": 86400}'
```

### Minimal request (if TTL is optional):
```bash
curl -X POST http://127.0.0.1:8080/url/ \
  -H "Content-Type: application/json" \
  -d '{"url": "https://github.com"}'
```

### Empty body (should now work without panic):
```bash
curl -X POST http://127.0.0.1:8080/url/ \
  -H "Content-Type: application/json" \
  -d '{}'
```