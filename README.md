# ci-cd-example-pipeline

A GitHub Actions pipeline template you can drop into any project. Copy the four workflows, adjust two lines, and you get lint, tests, Docker builds, semantic versioning, and automated GitHub Releases.

The Go API in this repo is just the example application the pipeline runs on: you can replace it with anything.

## Workflows

```
push / PR
    └── ci.yml
        ├── lint       checks code style and common errors
        ├── test       runs tests and measures coverage
        └── comment    posts a summary on the PR with coverage %

merge to main
    └── cd.yml
        ├── build.yml    compiles and pushes a Docker image to GHCR
        └── release.yml  creates a semver tag, CHANGELOG, and GitHub Release
```

### Versioning (Conventional Commits)

| Commit prefix | Bump |
|---------------|------|
| `feat:` | minor — `v1.0.0 → v1.1.0` |
| `fix:` | patch — `v1.1.0 → v1.1.1` |
| `feat!:` | major — `v1.x.x → v2.0.0` |
| `chore:` / `docs:` / `refactor:` | patch (default) |

## Adopting the pipeline

### 1. Copy the workflows

```bash
cp -r .github/workflows/ your-project/.github/workflows/
```

### 2. Swap the language-specific steps

Open `ci.yml` and replace the `lint` and `test` steps with your language's tooling:

| Language | Lint | Test |
|----------|------|------|
| Go | `golangci/golangci-lint-action` | `go test -race -coverprofile=coverage.out ./...` |
| Node.js | `npm run lint` | `npm test -- --coverage` |
| Python | `ruff check .` | `pytest --cov=. --cov-report=term` |
| Rust | `cargo clippy` | `cargo test` |
| Java | `checkstyle` | `mvn test` |

The `build.yml` and `release.yml` workflows are language-agnostic and work as-is for any Dockerized project.

### 3. Update the Docker image name in `build.yml`

The default tag is `ghcr.io/${{ github.repository }}`, which resolves automatically from your repo name, so no change needed in most cases.

### 4. Configure code quality tools

This pipeline integrates with SonarCloud and Codecov for static analysis and coverage tracking.

#### SonarCloud

1. Go to [sonarcloud.io](https://sonarcloud.io) and log in with GitHub
2. Click **"+"** → **"Analyze new project"** and import your repository
3. Choose **"With GitHub Actions"** and copy the generated token
4. Add `sonar-project.properties` to your repo root:

```properties
sonar.projectKey=YOUR_ORG_YOUR_REPO
sonar.organization=YOUR_ORG

sonar.sources=.
sonar.exclusions=**/*_test.go

sonar.tests=.
sonar.test.inclusions=**/*_test.go

sonar.go.coverage.reportPaths=coverage.out
```

#### Codecov

1. Go to [codecov.io](https://codecov.io) and log in with GitHub
2. Add your repository and copy the `CODECOV_TOKEN`

#### Coverage file by language

| Language | Command | Coverage file |
|----------|---------|---------------|
| Go | `go test -coverprofile=coverage.out ./...` | `coverage.out` |
| Node.js | `npx jest --coverage` | `coverage/lcov.info` |
| Python | `pytest --cov=. --cov-report=xml` | `coverage.xml` |
| Rust | `cargo tarpaulin --out Xml` | `cobertura.xml` |
| Java | `mvn test` (JaCoCo) | `target/site/jacoco/jacoco.xml` |

Update the `files` field in the Codecov step and `sonar.go.coverage.reportPaths` in `sonar-project.properties` to match your language's coverage file.

### 5. Configure GitHub secrets

Go to **Settings → Secrets and variables → Actions** and add:

| Secret | Description |
|--------|-------------|
| `SONAR_TOKEN` | Generated on SonarCloud |
| `CODECOV_TOKEN` | Generated on Codecov |

### 6. Configure GitHub branch protection

**Settings → Branches → Add rule for `main`**:
- ✅ Require a pull request before merging
- ✅ Require status checks: `lint`, `test`
- ✅ Require branches to be up to date before merging
- ✅ Do not allow bypassing

## Example application

The `handler/` and `main.go` files are a minimal Go HTTP API used to demonstrate the pipeline:

| Method | Route | Response |
|--------|-------|----------|
| GET | `/health` | `{"status":"ok"}` |
| POST | `/words` | `{"words":["hello","world"]}` |

Replace or delete them when adopting the pipeline for your own project.

## License

This project is licensed under the [MIT LICENSE](LICENSE).
