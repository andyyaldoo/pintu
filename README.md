### TODOs

- [ ] Create 2 folder to store the backend services
  - [ ] Golang
  - [ ] Nodejs
- [ ] infra folder to store the shared helm chart to be used by two backend services
- [ ] .github folder to store CI/CD pipeline
  - [ ] build
    - [ ] Add dependency caching
  - [ ] (Run unit test) run unit test
  - [ ] Docker build
  - [ ] Docker push
  - [ ] (Optiona) Deploy to k8s cluster
- [ ] Customize trigger.
  - [ ] Only deploy golang app when there are changes in the golang app
  - [ ] Deploy all apps when helm chart changes
- [ ] (Optional) Service made public using DNS (xip.io)
- [ ] Write explanations in README.md
