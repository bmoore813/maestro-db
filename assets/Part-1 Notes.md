# Assumptions

### Protobuf Compiler is installed:
[Install Guide](InstallProtoCompiler.md)


### API Design:
We will only need two end points(handlers):

1. `Produce`: For writing to the log
2. `Consume`: For reading from the log

Steps for Handlers should be the same

1. Unmarshall the request's JSON body into a native GO Struct
2. Run that end points business logic with the request to obtain a result
3. Marshal and write that result to the response

**Any further complication in the handlers should be abstracted out**

