# Goal

- Turn project into a service that scans across multiple computers for availablity and scalbility
- Allow multi tenancy for service
- Provide accessible interfaces that are heuristic


# Why gRPC?
- `Simplicity`: We want to focus on the business problems and not the complexity and technical minutiae of request response serilization while making these lower level serilization abstractions accessible.

- `Maintainability`: Writing the first version of a service is a brief period of the total time you’ll spend working on the service. Once your service is live and people depend on it, you must maintain backward compatibility. With request-response type APIs, the simplest way to maintain backward compatibility is to version and run multiple instances of your API.

- `Security`: gRPC supports Secure Sockets Layer/Transport Layer Security(`SSL/TLS`) to encrypt all data exchanged between the client and server. Further, we can attach crednetials so we know the identity of each user

- `Ease of use`: Writing the contracts in protobufs gives us an interface definition language that we can point to and easily understand without having to look into the implementation code. Further, we can elaborate on our protos to extend into other objects like database ddl and GraphQL contracts. 

- `Scalability`: Specifically this applies to two cases. First, the ability for many devs to contribute to a project. Two, the load balancing that can occur across multiple computers which is becoming the new normal for applications. 