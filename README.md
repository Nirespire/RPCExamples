# gRPC and Twirp examples

> Constructed from: 
- [Twirp Wiki](https://github.com/twitchtv/twirp/wiki/Usage-Example:-Haberdasher)
- [Google gRPC example](https://github.com/grpc/grpc-java/blob/v1.9.x/examples/README.md)


## Demo

### gRPC Java
- navigate into the `examples` directory
- `./gradlew installDist`
  - You might need to modify `grpcVersion` in the `build.gradle` to a non-SNAPSHOT version for Gradle to resolve all dependencies
- To run the server`./build/install/examples/bin/hello-world-server`
- To run the client`./build/install/examples/bin/hello-world-client`


### Twirp GO

- Move folders in `twirp-go` into your `GOPATH` directory
- 

