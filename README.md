# Maestro-DB
A distributed append only commit log used for quick writes and reads to any scale 


# Part 1 - Scaffolding
- Going to start off with building a simple Json over HTTP commit log
- Purge json use in favor for protobufs
- Wirte a log package which will be used for storing and looking up data

[Part-1 Notes](assets/Part-1%20Notes.md)

# Part 2 - Networking
- Setup and server and client to to define actions via gRPC
- Setup SSL/TLS to encrypt our data exchange between the client and the server
- Add in observibility...Metrics + Logs

# Part 3 - Distribute
- Make our services dicoverable
- Cordinate our microservices using consensus
- Load balancing

# Part 4 - Deployment
- Deploy codebase with Kubernetes locallay
- Add in terraform to deploy on GKE

# Part 5 - Profits
![alt text](assets/underpants_gnomes.jpeg)


# Technologies Used
[Cosmik Air](https://github.com/cosmtrek/air) - ☁️ Live reload for Go apps