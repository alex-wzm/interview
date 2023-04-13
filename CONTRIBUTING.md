# Contribute to The Interview Project

This monorepo is an open-ended project that candidates of all skill/career levels can build upon to demonstrate their technical knowledge and abilities. Since the project has no specific functional objectives beyond being the foundation for interviews, you can build anything you can dream of!

There's no restriction on languages, frameworks, or tools. However, do limit any required "getting started" steps by being aware of non-automatable dependencies to keep the interview easily accessible to all.

## Requirements

- [Go 1.18+](https://go.dev/doc/install)
- [gRPC prerequisites](https://grpc.io/docs/languages/go/quickstart/#prerequisites)
    <details>
    <summary>TL;DR</summary>

    1. Install protobuf compiler
        * Linux
          ```
          apt install -y protobuf-compiler
          ```
    
        * Mac
          ```
          brew install protobuf
          ```

        * Windows: use [WSL](https://learn.microsoft.com/en-us/windows/wsl/install) or [install latest release](https://grpc.io/docs/protoc-installation/#install-pre-compiled-binaries-any-os)

    2. Install Go plugins
       ```
       go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
       go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
       ```

    3. Add plugins to your path
       ```
       export PATH="$PATH:$(go env GOPATH)/bin"
       ```
    </details>

## Setup

Run the setup script to generate protobuf stubs.

```sh
./setup.sh
```

## Workflow

1. Create a [fork](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/working-with-forks/about-forks) of the repository and clone it locally.

2. Create a [branch](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-and-deleting-branches-within-your-repository)<sup>1</sup> for your changes.

3. Create a [pull request](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request-from-a-fork) when you are ready for feedback. Include a reasonably detailed description of the changes to help reviewers contextualize their feedback.

4. Once approved, use the [squash and rebase](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/configuring-pull-request-merges/about-merge-methods-on-github#squashing-your-merge-commits)<sup>2</sup> feature to commit your work to `main`.

Practice [Scaled Trunk-Based Development](https://trunkbaseddevelopment.com/#scaled-trunk-based-development) to reduce long-running branches in favor of **short-lived feature branches**.

The trunk is called `main`.

![scaled trunk-based development workflow](https://trunkbaseddevelopment.com/trunk1c.png)

## Debug

#### Inspect service with [grpcurl](https://github.com/fullstorydev/grpcurl)

```sh
# describe the service API
grpcurl -plaintext localhost:8080 describe InterviewService

# make an API call
grpcurl -plaintext -d '{"Name": "Obi-Wan"}' localhost:8080 InterviewService/HelloWorld
```

#### Go runtime error: _package your_pkg is not in GOROOT_

A Go [workspace](https://go.dev/ref/mod#workspaces) is a collection of [modules](https://go.dev/ref/mod#modules-overview) defined by the `go.work` file. For a module to be properly indexed by `gopls` or the go compiler, it must be referenced in the workspace file.

```
use ./your-module
```

---

## Appendix

### 1. Style guide

#### Branch names

Prefer `kebab-case` for branch names.

#### Linting and formatting

- [gopls](https://pkg.go.dev/golang.org/x/tools/gopls)

### 2. Why "Squash and rebase"?

Squashing commits into a single commit keeps the main code history focused on significant chunks of contributions while preserving as much detail about commits as desired by the author.

Using the rebase strategy to fold changes into the main code history (instead of using merge) maintains a chronological history best suited for easy code review and reversion when necessary.

As a rule of thumb, PR authors should manage the squashing and rebasing of their own branches. However, since this project is an interview format, it is likely that you will leave a PR up and not return to the repo. In this case, the interviewer should complete the PR on the author's behalf.
