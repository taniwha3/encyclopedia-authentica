# MD5-digest Authentication

[!WARNING]
> This is highly insecure, and is considered to be of equal strength
> to "HTTP Basic Authentication". MD5 is widely considered to be trivially
> broken at this point.

MD5-digets improves on HTTP Basic Authentication by hashing the password,
avoiding transmitting the password in plain text.

## Workflow

1. The client requests access to the protected resource.
1. The server responds with a 401 Unauthorized status code and a WWW-Authenticate with the following:
    * Realm
    * A nonce (number used only once) that changes with each authentication request
1. The client responds by calculating: 
    * HA = The MD5 hash of the username, realm, and password.
    * HB = Calculate the MD5 hash of the HTTP method and requested URI
    * Calculate the final digest by hashing HA + nonce + HB 
1. The client responds back with an Authorization header containing:
    * Username
    * Realm
    * Nonce
    * URI
    * Calculated digest (see previous section)
1. The server calculates the same digest using the user's password.
