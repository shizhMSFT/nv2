# Prototype 1 Scope

To support [Notary v2 requirements][notaryv2-requirements] for multiple signatures that don't change the artifact digest or tag, we must support new artifacts that link to existing artifacts.

## Goals & Intent

Experiment with storing linked artifacts in a registry, enabling the [reverse lookup pattern](https://github.com/notaryproject/requirements/blob/main/verification-by-reference.md).

The prototype should:

1. Validate the independent push, link, discover, pull experience.
2. Mock the end to end experience for building an image, signing it all offline, without having to first push to the registry to get the digest.
3. Not require updating any objects in a registry, including the tag
4. Not require tags or other pattern matching conventions to achieve the result
5. Enable the `nv2` client with the ability to find the required signature, based on the target artifact: eg: the `net-monitor:v1` image
6. Enable a user to locally validate the experience, without any dependency on any particular cloud provider.

The prototype should not:

1. Attempt to codify a solution that could be used in production, or at any scale. Can we prove the model could work? If so, we'll harden the implementation.
2. Focus on private or public key acquisition. Key management is not part of this prototype.
3. Validate with a container host. OPA validation will be done in a future prototype.

## Target Experience

The experience should focus on the e2e tooling, including the use of the docker cli.

[Demo Script](./docs/nv2/demo-script.md)

## Deliverables

1. Instance of cncf distribution to locally run. The prototype would be checked into: [notaryproject/distribution/prototype-1](https://github.com/notaryproject/distribution/tree/prototype-1).
2. A built image, users can run locally pushed to: [docker.io/notaryv2/registry:nv2-prototype-1](https://hub.docker.com/u/notaryv2)
3. A docker plug-in to simulate the docker build, sign, push experience

[notaryv2-requirements]:    https://github.com/notaryproject/requirements#goals