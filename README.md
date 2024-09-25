# ArgoCD Client

A Kubernetes client for ArgoCD, generated using `client-gen`, that simplifies interactions with ArgoCD resources without importing the entire ArgoCD project.

## Description

This project provides a typed Kubernetes client for ArgoCD by extracting and generating ArgoCD API types. The core challenge is that ArgoCD does not offer a standalone API and Client module, forcing Go applications to import the full ArgoCD codebase. This introduces significant issues, such as:

- Pulling in private and outdated Kubernetes libraries.
- Version conflicts with newer Kubernetes libraries.
- The need for extensive version overrides in Go modules.

By decoupling ArgoCD's API types from the main repository, we provide a lightweight Client and API alternative design for projects that need to create and manage ArgoCD resources without the bloat.

## Usage

This Go module can be imported into your project to provide access to ArgoCD's API types without requiring the entire ArgoCD project. It supports easier integration with modern Kubernetes controllers.

## Limitations

- **Source of API Types:** The API types in this project were directly copied from the ArgoCD project (specifically from [ArgoCD's application APIs](https://github.com/argoproj/argo-cd/tree/master/pkg/apis/application)).
- **Potential Maintenance Burden:** As ArgoCD evolves, this project may become difficult to maintain, requiring regular updates as the upstream ArgoCD API changes.
- **Hacky Approach:** This method is a workaround due to the absence of a dedicated API and Client repository for ArgoCD. Use it with caution, particularly if your project depends on long-term stability.

## Note

While this library provides an immediate solution to manage ArgoCD resources, we strongly recommend that users track changes in ArgoCD's upstream API to keep this module up to date.


