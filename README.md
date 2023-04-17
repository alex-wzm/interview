# The Interview Project

Technical interviews tend to be hand-wavy exercises focused on toy applications or over emphasize on lofty concepts and techniques. This project provides realistic scenarios for assessing a candidate's fundamental engineering skills: remote/asynchronous collaboration, code review, version control, and testing, to name a few.

Refer to the [contribution guide](/CONTRIBUTING.md) to get started.

## Architecture

Currently, the project consists of a Go microservice and a corresponding client application. They communicate using [gRPC](https://grpc.io/).

```
-------------------            -------------------
| Client          |            | Server          |
| --------        |            |                 |
| | SDK* |        | <--gRPC--> |                 |
| --------        |            |                 |
-------------------            -------------------

* Unimplemented
```

### Potential next steps

| Feature                | Description                            |
| ---------------------- | -------------------------------------- |
| Authentication         | Authenticate API requests.             |
| Client SDK             | Manage client-side authentication.     |
| Configuration          | Inject config values from environment. |
| Logging                | gRPC middleware, log levels.           |
| Testing                | Unit tests, integration tests.         |
| Continuous integration | GitHub Actions for tests.              |
