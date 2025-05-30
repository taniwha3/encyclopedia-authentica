# Encyclopedia Authentica

Welcome to the Encyclopedia Authentica! We list different kinds of authentication techniques and provide:

* A high level description of how they work
* Architecture Diagrams
* A short, likely insecure implementation to demonstrate how they work.
* Positives
* Negatives
* Common Traps and Pitfalls

Of course, this is all non-authorative, do your own additional research. No one but yourself is responsible for your use or misuse of any of this.

## Basic Authentication

* **INSECURE** [HTTP Basic Authentication](http-basic-authentication/README.md) - username/password sent as base64
* **INSECURE** HTTP Digest Authentication - password hashed with md5sum (or later algorithm)
* **INSECURE** [Form based authentication](form-based/README.md) - password sent over POST via form

## Token Based Authentication

* **JSON Web Tokens (JWT)** json tokens with claims and signatures embedded
* **OAuth 2.0** Authorization workflow for third party applications to authenticate against an IdP (Identity Provider)
* **OpenID Connect (OIDC)** identity layer extending oauth 2.0. oauth 2.0 provides authorization to login, oidc provides identity.
* **Bearer Tokens/API Keys** bearer tokens typically included in HTTP headers in a request.
* **Session Tokens** Server generated bearer tokens typically stored in cookies or local storage

## Certificate based authentication

* **Public Key Infrastructure (PKI)**
* **x.509 Certificates**
* **TLS**
* **Mutual TLS (mTLS)**
* **Diffie-Hellman**


## Things to sort through
* SAML
* Webauthn/FIDO2
* Secure Remote Password (SRP) - zero knowledge proof
* Kerberos
* HOTP
* TOTP
* SMS/Email Verification Codes - Some are literally just links that log you in :cry
* PUsh Notifications
* Passkeys
* Hardware Security Keys
* WS-Federation
* DID
* Verifiable Credentials (VCs)
* Adaptive Authentication
* Biometrics
* NTLM
* Cookie Based Authentication
