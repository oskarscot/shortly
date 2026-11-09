# Shortly

A simple URL shortener built to learn Go backend development.

Requires Go 1.26 or later.

Build and run from the project root:

```sh
go build -o shortly ./cmd
./shortly
```

The server runs at http://localhost:8080.

Create a short link (`POST /api/v1/links`):

```sh
curl -X POST http://localhost:8080/api/v1/links \
  -H 'Content-Type: application/json' \
  -d '{"url":"https://example.com"}'
```

The response contains the link's `id`, `code`, and `url`.

List all links (`GET /api/v1/links`):

```sh
curl http://localhost:8080/api/v1/links
```

To get redirected, open `http://localhost:8080/s/<code>` in your browser,
replacing `<code>` with the code from the response. Or follow the redirect with curl:

```sh
curl -L 'http://localhost:8080/s/<code>'
```
