# General explanation

Tools:
- Github Actions for running CI/CD workflows
- checkity.io for code coverage
- helm charts for templating k8s manifests
- minikube for doing sanity checks after service is deployed

The approach that I used while working on this is to keep things simple and free ($). Hence, the tools above are chosen. Github Actions has a free 2000 workflow run/month. checkity.io is the only free code coverage tool.

What each folder stores:
- deploy/helm/http -> store helm chart
- .github/workflows -> store CI/CD workflows
- services/ -> store application code (nodejs & golang)
- services/<app>/deploy -> store Dockerfiles

The general CI/CD workflow:
1. Build & Test
2. Docker build and push
3. Deploy to k8s cluster
4. `curl` service URLs

The triggers for CI/CD workflows:
- push to `main` branch and PR raised against `main` branch
    - If only 1 service is modified, then only the CI/CD pipeline for that service will be run
    - If the shared helm chart is modified, both CI/CD pipelines will be run
