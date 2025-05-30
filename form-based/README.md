# Form Based Authentication

> TODO: Note - Form-based authentication doesn't have a specific RFC standard

> TODO: Make Golang example [Golang Example](form-based-authentication/main.go)

> [!WARNING]
> Consider using OAuth 2.0/OpenID Connect instead
> 
> Form-based authentication is considered insecure since it transmits the password over the network. If you must use form-based authentication, ensure you use a recent version of TLS.

> Use TLS 1.3 or later: https://datatracker.ietf.org/doc/html/rfc8446

Form-based authentication uses HTML-based forms to submit login information
over an HTTP POST to an endpoint on the server. The server uses a secure
hashing algorithm to validate against a hash stored in the database. If the
user successfully authenticates, a unique session identifier is returned to the
user, typically as a cookie, which the user's browser automatically includes in
subsequent requests. The server maintains a store of active session IDs and
their user mappings.

> [!WARNING]
> For the hashing algorithm, consider using bcrypt or argon2. These are
> intentionally computationally expensive, which makes a stolen
> hash database more expensive to crack.

## Technical Details

### Example HTTP Form

```html
<form method="POST" action="/login" enctype="application/x-www-form-urlencoded">
  <!-- Consider adding a CSRF token here. TODO: Add CSRF link when written -->
  <input type="email" name="username" required>
  <input type="password" name="password" required>
  <input type="submit" value="Login">
</form>
```

### Server Implementation Details

1. The server accepts POST requests to an endpoint, e.g., `/login`.
1. The server parses the form data and extracts the username/password.
1. The password is hashed, ideally with a salt.
1. The hash is compared with what's in the database.
1. If the comparison fails, it returns an error and the flow terminates.
1. If the comparison succeeds, the user is considered authenticated.
1. A mapping of the `Session ID` to the `User ID` is created and inserted in a session store. (P.S. Best practice is to also store an expiration time.)
1. The session ID is returned to the user and stored as a cookie (with HttpOnly, Secure, and SameSite flags).
1. The user's browser automatically submits the cookie with the session ID for all subsequent requests.
1. The server validates the session ID and expiration time from the submitted cookie.

### HTTP Form Authentication Weaknesses

There are several weaknesses that must be taken into account:

* Without TLS, credentials are transmitted in plaintext over the network. Always use HTTPS/TLS for form-based authentication.
* Session cookies are vulnerable to being stolen through XSS attacks, CSRF attacks, and session hijacking. Session cookies sent over unecrypted connections may be lifted on the network itself. Malicious JavaScript may also steal credentials.
* Session fixation may occur if session IDs are repeated or are predictable.
* The developer may in good faith attempt to implement controls to mitigate issues, but will likely introduce vulnerabilities.
* Server vulnerabilities such as SQL injections may be present, putting sensitive data accessible by that server at risk.
