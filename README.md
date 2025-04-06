# Zanbil

![Go Version](https://img.shields.io/badge/Golang-1.23-66ADD8?style=for-the-badge&logo=go)
![App Version](https://img.shields.io/github/v/tag/mohammadne/zanbil?sort=semver&style=for-the-badge&logo=github)
![Repo Size](https://img.shields.io/github/repo-size/mohammadne/zanbil?logo=github&style=for-the-badge)
<!-- ![Coverage](https://img.shields.io/codecov/c/github/mohammadne/zanbil?logo=codecov&style=for-the-badge) -->

> Zanbil is a persian name means basket 🧺

Zanbil is a lightweight and efficient Go program designed to demonstrate how to write, test, build, and release Go applications across multiple platforms, including Docker (Podman) and Kubernetes which follows best practices for error handling, internationalization and clean architecture.

Whether you’re building a small project or an enterprise-level application, Zanbil provides valuable insights into structuring and developing Go programs with scalability and maintainability in mind.

![Home](https://github.com/user-attachments/assets/017a32e9-8f40-4228-882f-64f0e700748c)

## Development

```bash
# setup requirements
cd hacks/release/docker && podman compose -f compose.local.yml up -d

go run cmd/migration/main.go --direction=up # setup the database
go run cmd/server/main.go # run the server

# teardown application
CTRL+C # to terminate the go application
hacks/release/docker && podman compose -f compose.local.yml down
```

## Features

### Logic

- mimic the ecommerce platform by having products and categories
- two seperate servers for requests and monitoring stuffs
- using fiber v3 with middlewares and handlers (template files are embedded using the `embed` package)
- custom config loader via `envconfig`
- zap logger and prometheus monitoring
- Redis and postgresql database
- Hexagonal architecture by seperating business layers
- Custom DB migration handling (`cmd/migration/schemas`)
- i18n and translation handling for api responses

### Tests

I have written unit tests for different modules, for mocking redis I have used `miniredis` and for mysql database I have used the package `sqlmock`, also for mocking api calls to this repositories I have used the `testify` package to mock the behavior of this repositories.

For the service layer and the item's service I have added some sample `benchmark` tests to measure the performance, altough this method uses the mocks and mocks do not represent real performance.

There is a functional test and an integration test in the tests directory which are very basic test cases spinned up by docker-compose (the one in deployments/docker/compose.local.yml). in the functional part we test the cache module and in the integration test we test the behavior of the whole application (set cookie and send the request to the applicaion, for more advanced integration tests we can examine the expectations like the value of items in redis and mysql by an actual redis and mysql driver and assert if something goes wrong).

I have also added `k6` for the `load` testing.

So In my project, I have the following tests:

- unit tests
- benchmark tests
- functional
- integration
- load test via k6

### Provisioning

- use `cluster` Ansible to provision k8s cluster via kind for testing purposes

### Deployment

#### Use github Action for the build process

#### Use compose file with podman for container deployment

The application can be deployed using `compose`:

- **Local Development:** Use the `.local` compose file to resolve dependencies locally.
- **Production:** The `.prod` compose file is designed for deploying the application to a server.

#### use k8s helm chart via helmsman and helm-secret

A Kubernetes setup (`k8s/`) is provided for deploying the application with `helmsman` and `helm` charts.

### TODOs

- use idempotent admin APIs
- implement rate limiting
