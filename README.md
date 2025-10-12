# Liatrio Apprenticeship Exercise

This repository contains a minimal Go web service packaged with Docker and automated with GitHub Actions per the Liatrio apprenticeship interview exercise.

## Application

The application is built with [Fiber](https://gofiber.io/) and exposes a single endpoint:

```
GET /
```

Response (minified JSON):

```json
{"message":"My name is Vladimir Avdeev","timestamp":1716239023}
```

`timestamp` is generated dynamically each request.

### Running locally

```bash
go run .
```

Visit http://localhost:8080/.

### Tests

```bash
go test ./...
```

### Docker

Build and run:

```bash
docker build -t liatrio-demo:local .
docker run --rm -p 8080:80 liatrio-demo:local

The container sets `PORT=80` by default, so adjust the host mapping if you need a different external port.
```

## GitHub Actions

Workflow: `.github/workflows/ci.yml`

Pipeline steps:

1. Run unit tests (`go test ./...`)
2. Build the Docker image
3. Start the container on port 80 and validate it with the [`liatrio/github-apprentice-action`](https://github.com/liatrio/github-apprentice-action)
4. Push a uniquely versioned tag (`:${GITHUB_RUN_NUMBER}`) and `:latest` to Docker Hub

### Required secrets

| Name | Purpose |
| ---- | ------- |
| `DOCKERHUB_USERNAME` | Docker Hub username that owns the target repository |
| `DOCKERHUB_TOKEN` | Docker Hub access token with `repo` scope |

Set the Docker repository name to match `DOCKERHUB_USERNAME/liatrio-demo` or update `IMAGE_NAME` in the workflow.

## Deployment

After the image is pushed you can deploy it to any container platform. A simple approach:

1. **AWS ECS (Fargate)** – create a task definition pinned to the pushed image tag, wire it to a service behind an Application Load Balancer, and expose container port 80 (map to any listener you prefer).
2. **GCP Cloud Run** – deploy the image with `gcloud run deploy`, make sure to pass the fully qualified Docker Hub image reference and keep the container port at 80 (set a different `PORT` env var if needed).
3. **Azure Container Apps** – use `az containerapp up` pointing at the tag produced by the CI workflow.

For extra credit, create a second workflow (e.g., `deploy.yml`) that triggers on `push` to `main`, pulls the image tag produced by the CI workflow, and updates the chosen platform automatically.

Document any infrastructure as code or cloud console steps you take so they can be demonstrated later.
