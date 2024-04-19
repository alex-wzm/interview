# Auth Service

## Design

The goal of this change is to add an authorization service.  The idea behind the service is to place secret management and generating JWT in a single service.  Basically, everything in the current `jwt.go` file would be moved into the service.  For the sake of time, I only implemented the code to generate JWT.  A more complete service would also provide an interface for validating a JWT, which the `service-go` process would use.


Due to time constraints, I did not implement any password authentication in the new service.  In a more complete implementation, the connection would be encrypted and either basic auth or oAuth could be used to perform the authentication.  Additionally, the secrets would not be directly stored in the service.  Another service such as AWS Secret Manager or HashiCorp Vault


## Running Service

To start the service:

1. At the top level directory, run `./setup.sh` This will create the protobufs needed for the new auth service
2. cd into `auth-go`
3. run `go run cmd/auth/auth.go`
4. In another terminal start the `service-go` service
5. In a third terminal, run the `client-go` process
