# User management

```bash
curl -X POST -H "Content-Type: application/json" -d '{"Name": "John Doe", "Email": "john@example.com"}' http://localhost:8888/createUser

curl "http://localhost:8888/readUser?email=john@example.com"
```