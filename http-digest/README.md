# HTTP Digest Authentication

> [!WARNING]
> This is highly insecure and is considered to be of equal strength to HTTP
> Basic Authentication. The original algorithm used in this authentication
> technique is MD5, which is widely considered to be trivially broken at this
> point.
>
> This technique requires passwords to be stored in plaintext or as reversible
> hashes on the server, increasing the total risk significantly if the server 
> is compromised.

HTTP Digest Authentication improves on HTTP Basic Authentication by hashing the
password, avoiding transmitting the password over the network in plain text.

## Workflow (with original MD5)

1. The client requests access to the protected resource.
1. The server responds with a 401 Unauthorized status code and a WWW-Authenticate header with the following:
    * Realm
    * A nonce (number used only once) that changes with each authentication request
1. The client responds by calculating: 
    * HA1 = MD5(username:realm:password)
    * HA2 = MD5(method:digestURI)
    * response = MD5(HA1:nonce:HA2) 
1. The client responds back with an Authorization header containing:
    * Username
    * Realm
    * Nonce
    * URI
    * Response (the calculated digest)
    * Algorithm (typically "MD5")
    * Opaque (if provided by server)
1. The server calculates the same digest using the user's password.
1. The server grants access if the digest matches.

At the time, this was considered a major improvement over HTTP Basic
Authentication. However, MD5 is now vulnerable to collision attacks and rainbow
table attacks, making this method significantly weaker than originally intended.

If the algorithm is upgraded to SHA-256, this solves the problem of hash
collisions being trivial, but the server must still store plain text passwords
for this technique to work. If the server is compromised, the passwords will
all be compromised, leaving the user susceptible to password spraying attacks.

Special attention must also be given to the nonce. The nonce must be securely
generated and of sufficient length.
