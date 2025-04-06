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

### Core Logic

- Mimics an e-commerce platform with **products** and **categories**
- Two separate servers: one for **API requests**, one for **monitoring**
- Built with **Fiber v3**, using embedded HTML templates (`embed` package)
- Custom configuration loader based on `envconfig`
- Structured **Zap logger** and integrated **Prometheus** metrics
- Uses **PostgreSQL** and **Redis** as primary data stores
- Clean Hexagonal Architecture (business logic separated from infrastructure)
- Custom database migration system (`cmd/migration/schemas`)
- Internationalization (i18n) support for API responses

### Tests

Zanbil includes a variety of tests to ensure quality and performance:

- **Unit Tests** – Use `testify`, `sqlmock`, and `miniredis` to test components in isolation
- **Benchmark Tests** – Evaluate performance of services (note: benchmarks use mocks)
- **Functional Tests** – Validate cache behavior with Redis using Docker Compose
- **Integration Tests** – End-to-end tests using real HTTP requests and cookies
- **Load Testing** – Performed via `k6` to simulate real-world traffic

Tests are organized to reflect realistic development environments:

- Functional and integration tests use the `deployments/docker/compose.local.yml` setup
- Redis and PostgreSQL expectations can be verified via actual drivers for more robust assertions

### Provisioning

- Use [**cluster** Ansible tool](https://github.com/k3s-io/k3s-ansible) to provision a Kubernetes cluster with `kind` for local testing

### Deployment

#### Use github Action for the build process

- Automated build and release pipeline using GitHub Actions

#### Use compose file with podman for container deployment

- **Local:** Use `.local` compose file to run services on your local machine
- **Production:** Use `.prod` compose file for deploying to production environments

#### use k8s helm chart via helmsman and helm-secret

- Deploy with Helm via `helmsman` and `helm-secrets`
- Kubernetes manifests and Helm chart available in the `k8s/` directory

### TODOs 📝

- [ ] Add idempotent admin APIs
- [ ] Implement request rate limiting
