# HTTP Basic Authentication

[RFC 7617](https://datatracker.ietf.org/doc/html/rfc7617)

[Golang Example](http-basic-authentication/main.go)

> [!WARNING]
> HTTP Basic Auth is considered to be insecure. If you **must** use it, ensure you are using a TLS connection since it sends the password over the wire.

HTTP Basic Authentication transmits credentials by including a Base64 encoded string of "username:password" in the Authorization header of each HTTP request.

## Technical Details

### Header Format

```
Authorization: Basic <credentials>
```

### Credential Encoding

The credentials are encoded as follows

```
base64($username:$password)
```

Note: We use Standard Base64 encoding, not URL Encoding. Typically, the authorization should work with or without padding.

Example:

```sh
username="admin"
password="secret"

credentials="${username}:${password}"

echo -n "$credentials" | base64
```

### Challenge-Response Flow

* When a protected resource is requested without credentials, the server responds with HTTP status code 401: Unauthorized.
* The response icludes a `WWW-Authenticate` header requesting Basic authentication.

```
WWW-Authenticate: Basic realm="example"
```

* The client resends the request with the `Authorization` header.
* The server validates the credentials and grants access if valid.

### Lack of security

The biggest disadvantage to this approach is the username and password are sent
as HTTP headers over the wire. There are multiple authentication schemes that
are significantly more secure that developers should use if they are writing a
new application. HTTP Basic Authentication should only be used in the following
conditions.

* The program is legacy and has no path towards supporting anything newer
* The connection is encrypted with TLS.

### Example curl command

```sh
curl -v admin:secret http://localhost:8080/protected
curl -v -u admin:secret http://localhost:8080/protected
```
