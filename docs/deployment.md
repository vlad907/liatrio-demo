# Deployment Plan

Choose whichever cloud platform you are most comfortable with. Below is a suggested path for AWS ECS (Fargate). Adapt to GCP or Azure if preferred.

## Prerequisites

- Docker Hub repository matching the value used in `.github/workflows/ci.yml`
- AWS account with permissions to create ECS, ECR (optional), IAM roles, VPC networking, and Load Balancers
- AWS CLI configured locally or a GitHub Actions OIDC role

## Steps

1. **Networking**
   - Reuse an existing VPC or create a new one with at least two subnets across distinct AZs.
   - Ensure the subnets have internet access (NAT or Internet Gateway) for pulling images from Docker Hub.

2. **Task Execution Role**
   - Create an IAM role with the `AmazonECSTaskExecutionRolePolicy`.
   - Allow pulling Docker Hub images (no extra permissions required if the repo is public).

3. **Task Definition**
   - Define a Fargate task with one container.
   - Image: `docker.io/<DOCKERHUB_USERNAME>/liatrio-demo:<CI_TAG>`
   - Port mappings: container and host port 8080 (protocol TCP).
   - Environment: optionally set `PORT=8080` (defaults to `8080`).

4. **Service & Load Balancer**
   - Create an Application Load Balancer targeting the ECS service.
   - Register the service with `desiredCount=1` (scale later as needed).
   - Health check path: `/`.

5. **DNS (optional)**
   - Map a friendly DNS name to the ALB using Route53 or your domain provider.

## Automating Deployments (Extra Credit)

1. Store the last successful image tag as an artifact in the CI workflow or emit it via `workflow_run`.
2. Create `.github/workflows/deploy.yml` triggered by `workflow_run` or `push` to `main`.
3. In the deploy workflow:
   - Assume an AWS IAM role via OIDC (`aws-actions/configure-aws-credentials`).
   - Update the ECS service with the new image tag using `aws ecs update-service --force-new-deployment`.

Document the image tag you deploy so it can be shown during the live demo.
