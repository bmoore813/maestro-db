# Security

[List of Actors](https://en.wikipedia.org/wiki/Alice_and_Bob#Cast_of_characters)

Security in distributed services can be broken down into three steps

1. Encrypt data in-flight to protect against man in the middle attacks
2. Authenticate to identify clients; and
3. Authorize to determine the permissions of the identified clients


### Encryption In-Flight Data

TLS encryption:
1. Specify which version of TLS they'll use
2. Decide which cipher suites(the set of encryption algorithms)
3. Authenticate the identity of the server via the servers private key and the certificate authority's digital signature; and
4. Generate session keys for symmetric enryption after the handshake is complete

Once this is done the client and server should be able to communicate securely

### Authenticate to Identify Clients

Most web services use TLS for one-way authentication and on authenticate the server. The authentication of the client is left to the application to work out usually by some username/password + tokens. However, we leverage TLS two-way authentication, in which both the server and the client validate the others communication. In this setup both the server and the client use a certificate to prove their identity. 


### Authorize to Determine the Permissions of the Clients

Differentiating between authentication and authorization is necessary when you have a resource with shared access across varying levels of ownership. Maestro-DB has controls where Alice might be the owner and have both read and write access to the contents of the log, whereas Bob might be allowed to read the contents, but isn't able to write. In this type of situation,  you need authorization with granular access control. Maestro has a built in access control list-based authorization to control whether a client is allowed to read, write or both to the log.


### CloudFlare For CA
Since this is an internal tool