<<<<<<< HEAD
# Kubernetes Init Container Injector Operator

An operator that automatically injects init containers into Kubernetes Deployments based on annotations.

## Overview

This operator watches Kubernetes Deployments and automatically injects init containers based on annotations. It provides a simple, declarative way to add initialization containers to your deployments without modifying their core specifications.

## Features

- Annotation-based configuration
- Automatic init container injection
- Configurable container image and registry
- Customizable commands and arguments
- Non-intrusive deployment modification

## Installation

### Prerequisites

- Kubernetes cluster 1.16+
- kubectl configured to communicate with your cluster
- Go 1.19+ (for building from source)

### Deploy the Operator

```bash
# Clone the repository
git clone https://github.com/yourusername/initcontainer-injector-operator
cd initcontainer-injector-operator

# Build and deploy
make build
make deploy
```

## Usage

To use the init container injector, add the following annotations to your Deployment:

### Required Annotation

- `initcontainer_injector_args`: Comma-separated list of arguments for the init container

### Optional Annotations

- `initcontainer_injector_image`: Container image name (defaults to "default")
- `initcontainer_injector_registry`: Container registry (defaults to "docker.io")
- `initcontainer_injector_command`: Comma-separated command (defaults to "/bin/sh,-c,echo")

### Example Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: example-deployment
  annotations:
    initcontainer_injector_args: "-u,dep1,-u,dep2,-m,3600"
    initcontainer_injector_image: "nginx"
    initcontainer_injector_registry: "docker.io"
    initcontainer_injector_command: "/bin/sh,-c,echo"
spec:
  replicas: 1
  selector:
    matchLabels:
      app: example-app
  template:
    metadata:
      labels:
        app: example-app
    spec:
      containers:
      - name: main-app
        image: nginx
```

This will result in an init container being injected with:
- Image: `docker.io/nginx:latest`
- Command: `["/bin/sh", "-c", "echo"]`
- Args: `["-u", "dep1", "-u", "dep2", "-m", "3600"]`

## How It Works

1. The operator watches for Deployments in the cluster
2. When a Deployment with the required annotation (`initcontainer_injector_args`) is detected
3. The operator reads all related annotations
4. An init container is created with the specified configuration
5. The init container is injected into the Deployment
6. The Deployment is updated with the new configuration

## Configuration

### Operator Flags

- `--metrics-bind-address`: The address the metric endpoint binds to (default ":8080")
- `--health-probe-bind-address`: The address the probe endpoint binds to (default ":8081")
- `--leader-elect`: Enable leader election for controller manager (default "false")

## Development

### Building from Source

```bash
# Clone the repository
git clone https://github.com/yourusername/initcontainer-injector-operator
cd initcontainer-injector-operator

# Build
make build

# Run locally
make run
```

### Running Tests

```bash
# Run tests
make test
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

- Inspired by Kyverno's init container injection policy
- Built with [Kubebuilder](https://github.com/kubernetes-sigs/kubebuilder)

## Support

Please [open an issue](https://github.com/yourusername/initcontainer-injector-operator/issues) for support.

## Project Status

This project is under active development. Feedback and contributions are welcome!

## Roadmap

- [ ] Add metrics and monitoring
- [ ] Implement webhook validation
- [ ] Add support for more configuration options
- [ ] Improve error handling and recovery
- [ ] Add comprehensive test coverage
=======
# deployment-injector-operator
// TODO(user): Add simple overview of use/purpose

## Description
// TODO(user): An in-depth paragraph about your project and overview of use

## Getting Started

### Prerequisites
- go version v1.22.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/deployment-injector-operator:tag
```

**NOTE:** This image ought to be published in the personal registry you specified.
And it is required to have access to pull the image from the working environment.
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/deployment-injector-operator:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## Project Distribution

Following are the steps to build the installer and distribute this project to users.

1. Build the installer for the image built and published in the registry:

```sh
make build-installer IMG=<some-registry>/deployment-injector-operator:tag
```

NOTE: The makefile target mentioned above generates an 'install.yaml'
file in the dist directory. This file contains all the resources built
with Kustomize, which are necessary to install this project without
its dependencies.

2. Using the installer

Users can just run kubectl apply -f <URL for YAML BUNDLE> to install the project, i.e.:

```sh
kubectl apply -f https://raw.githubusercontent.com/<org>/deployment-injector-operator/<tag or branch>/dist/install.yaml
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

**NOTE:** Run `make help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

>>>>>>> develop
