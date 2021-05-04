# Surf Report
Surf Report API

# Push to docker
``` docker login ```
``` docker build . -t reactor-lab ```
``` docker tag reactor-lab bradmccoydev/reactor-lab:latest ```
``` docker push bradmccoydev/reactor-lab:latest ```

# Change out Docker image in Kubernetes
``` kubectl set image deployment.apps/surf-report surf-report=bradmccoydev/surf-report:latest -n surf-report ```

# Liveiness Probes & Readiness Probes
These probes can be found in the Startup.cs
``` endpoints.MapHealthChecks("/healthz") ```
``` endpoints.MapHealthChecks("/ready") ```


