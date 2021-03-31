# Prototype 2 Scope

To support [Notary v2 requirements](https://github.com/notaryproject/requirements#goals) for multiple signatures that don't change the artifact digest or tag, we must support new artifacts that link to existing artifacts.

## Goals & Intent

With [Prototype-1]() complete, we must now harden the linked artifacts implementation with a spec and design that could be deployed for production validation. While the design will be prototyped in CNCF Distribution, the full implementation will not be completed in CNCF Distribution. Cloud providers have cloud optimized implementations, and implementers of CNCF Distribution have their unique product implementations. The purpose of the prototype is to prove the APIs are capable of being implemented.

The prototype should:

1. Implement the [oci.artifact.manifest spec](https://github.com/opencontainers/artifacts/pull/29), supporting linked manifests in CNCF distribution.
1. Implement the [oci.artifact.manifest links API](https://github.com/opencontainers/artifacts/pull/29), supporting retrieving linked manifests in CNCF distribution.
1. Implement [ORAS client](https://github.com/notaryproject/oras/tree/prototype-2) libraries, enabling the pushing, linking, discovering and pulling of artifacts, using the [OCI.artifact.manifest](https://github.com/opencontainers/artifacts/pull/29)
1. ORAS to support pushing by digest, enabling the ability to update an existing tag, after the updated image and signatures have been posted.
1. Update the [nv2 client](https://github.com/notaryproject/nv2/tree/prototype-2) and [docker-generate](https://github.com/notaryproject/docker-generate/tree/prototype-2) plug-in, supporting the oci.artifact.manifest and linked manifest discovery.
1. Incorporate any [nv2](https://github.com/notaryproject/nv2/tree/prototype-2) client feedback items related to signing and verifying.
1. With ORAS support, users should be able to link Notary v2 signatures or other files, including SBoM simulated files. ORAS push, would push the single file reference as a tar file, eliminating the `org.opencontainers.image.title` annotation. See [oras issue # 114](https://github.com/deislabs/oras/issues/114)
1. Update the [docker-generate](https://github.com/notaryproject/docker-generate/tree/prototype-2) plug-in to support pushing the image by digest, the linked signature, then pushing the tag update.
1. Implement an OPA/Gatekeeper policy that confirms signature validation, before deploying an image to k8s.

The prototype should not:

1. Attempt to complete the CNCF distribution implementation with de-duping or garbage collection. Linked storage layout will be completed, enabling CNCF distribution implementors to complete their specific implementations.
2. Focus on private or public key acquisition. Key management is not part of this prototype.

## Target Experience

The experience should focus on the e2e tooling, including the use of the docker cli.

[Demo Script](https://github.com/notaryproject/nv2/blob/prototype-2/docs/nv2/demo-script-detailed.md)

## Deliverables

1. Instance of cncf distribution to locally run. The prototype would be checked into: [notaryproject/distribution/prototype-2](https://github.com/notaryproject/distribution/tree/prototype-2).
2. A built image, users can run locally pushed to: [docker.io/notaryv2/registry:nv2-prototype-2](https://hub.docker.com/u/notaryv2)
3. A docker plug-in to simulate the docker build, sign, push experience
4. An OPA/Gatekeeper policy for validating signatures, prior to image deployment
