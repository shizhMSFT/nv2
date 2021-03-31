# Prototype 3 Scope DRAFT

With linked artifacts, including Notary v2 signatures and SBoMs persistance completed in Prototype 2, we must address tag locking and key management requirements. Prototype 3 will exist with an understanding of what we might do next, but isn't expected to produce a solid spec. It will validate what it would take to sign tags and manage key, and/or signature lifecycle requirements.

## Goals & Intent

Implement a tag signing solution for oci-distribution-spec based registries.

The prototype should:

1. Experiment with how tags may be signed, with notary v2 signatures.
2. Experiment how key and/or signature revocation/invalidation scenarios may be implemented, without requiring short-lived keys that must be continually updated.

The prototype should not:

1. TBD

## Target Experience

TBD:

## Deliverables

1. An experimental instance of cncf distribution, supporting tag signing. The prototype would be checked into: [notaryproject/distribution/prototype-3](https://github.com/notaryproject/distribution/tree/prototype-3).
2. A built image, users can run locally pushed to: [docker.io/notaryv2/registry:nv2-prototype-3](https://hub.docker.com/u/notaryv2)
3. A docker plug-in to simulate the docker build, sign, push experience
4. An OPA/Gatekeeper policy for validating signatures, prior to image deployment
