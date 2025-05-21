# Case: Feedback Application

This repository contains a simple feedback application designed for internal use by developers to rate and provide feedback on the internal development portal. The application is built in Go, deployed on Google Cloud Run, and uses a PostgreSQL database for storing feedback. The infrastructure is managed using Terraform, and CI/CD pipelines are implemented using GitHub Actions.

---

## Features

- **Feedback Application**:
  - A Go-based web application that allows developers to submit anonymous feedback with a dice-roll rating (1-6) and an optional note.
  - Feedback is stored securely in a PostgreSQL database.
  - Includes a health check endpoint for monitoring.

- **Infrastructure as Code**:
  - Terraform is used to provision the required Google Cloud resources:
    - A VPC network.
    - A Cloud SQL PostgreSQL instance.
    - A Cloud Run service for deploying the application.

- **CI/CD Pipelines**:
  - GitHub Actions workflows automate the deployment process:
    - Terraform workflow to validate, plan, and apply infrastructure changes.
    - Cloud Run deployment workflow to build and deploy the application.

---

## CI/CD Pipelines

**Terraform Workflow**:
The terraform.yml workflow automates the validation, planning, and application of Terraform configurations. It is triggered on pushes to the main branch.

**Cloud Run Deployment Workflow**:
The workload.yml workflow automates the build and deployment of the application to Cloud Run. It is triggered on pushes to the main branch.

## Feedback Application Endpoints
Submit Feedback: POST /submit-feedback
Accepts a JSON payload with rating (1-6) and an optional note.
Health Check: GET /
Returns 200 OK if the application is running.

## License
This project is for educational purposes and is not licensed for production use.