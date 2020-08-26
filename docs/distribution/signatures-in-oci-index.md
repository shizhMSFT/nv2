# OCI Distribution

We introduce new REST APIs in the registry to support storing signature objects together with the target artifacts, and retrieving them for verification.

## Table of contents
1. [Sample APIs for retrieving signatures](#Sample-APIs-for-retrieving-signatures)
2. [Sample APIs for storing signatures](#Sample-APIs-for-storing-signatures)
3. [Implementation](#Implementation)

## Sample APIs for retrieving signatures

The list of signatures for an artifact can be retrieved from the registry.

### _Get a list of paginated signatures by tag_

### Request
`GET http://localhost:5000/v2/hello-world/manifests/v1.0/signatures/`

### Response

```json
{
    "tag": "v1.0",
    "digest": "sha256:80559bf80b44ce6be8234e6ff90a1ac34acbeb826903b02cfa0da11c82cbc042",
    "@nextLink": "{opaqueUrl}",
    "signatures": [
        "sha256:2235d2d22ae5ef400769fa51c84717264cd1520ac8d93dc071374c1be49cc77c",
        "sha256:3335d2d22ae5ef400769fa51c84717264cd1520ac8d93dc071374c1be49cc88d"
    ]
}
```

### _Get a list of paginated signatures by digest_

### Request
`GET http://localhost:5000/v2/hello-world/manifests/sha256:80559bf80b44ce6be8234e6ff90a1ac34acbeb826903b02cfa0da11c82cbc042/signatures/`

### Response

```json
{
    "digest": "sha256:80559bf80b44ce6be8234e6ff90a1ac34acbeb826903b02cfa0da11c82cbc042",
    "@nextLink": "{opaqueUrl}",
    "signatures": [
        "sha256:2235d2d22ae5ef400769fa51c84717264cd1520ac8d93dc071374c1be49cc77c",
        "sha256:3335d2d22ae5ef400769fa51c84717264cd1520ac8d93dc071374c1be49cc88d"
    ]
}
```

## Sample APIs for storing signatures

### _Put an OCI index linking a signature to a collection of manifests by tag_

### Request
`PUT https://localhost:5000/v2/hello-world/manifests/signature-1`
```json
{
  "schemaVersion": 2.1,
  "mediaType": "application/vnd.oci.image.index.v2+json",
  "config": {
    "mediaType": "application/vnd.cncf.notary.config.v2+jwt",
    "digest": "sha256:2235d2d22ae5ef400769fa51c84717264cd1520ac8d93dc071374c1be49cc77c",
    "size": 1906
  },
  "manifests": [
    {
      "mediaType": "application/vnd.oci.image.manifest.v1+json",
      "digest": "sha256:80559bf80b44ce6be8234e6ff90a1ac34acbeb826903b02cfa0da11c82cbc042",
      "size": 7023,
      "platform": {
        "architecture": "ppc64le",
        "os": "linux"
      }
    }
  ]
}
```

### _Put an OCI index linking a signature to a collection of manifests by digest_

### Request
`PUT https://localhost:5000/v2/hello-world/manifests/sha256:90659bf80b44ce6be8234e6ff90a1ac34acbeb826903b02cfa0da11c82cbc042`
```json
{
  "schemaVersion": 2.1,
  "mediaType": "application/vnd.oci.image.index.v2+json",
  "config": {
    "mediaType": "application/vnd.cncf.notary.config.v2+jwt",
    "digest": "sha256:2235d2d22ae5ef400769fa51c84717264cd1520ac8d93dc071374c1be49cc77c",
    "size": 1906
  },
  "manifests": [
    {
      "mediaType": "application/vnd.oci.image.manifest.v1+json",
      "digest": "sha256:80559bf80b44ce6be8234e6ff90a1ac34acbeb826903b02cfa0da11c82cbc042",
      "size": 7023,
      "platform": {
        "architecture": "ppc64le",
        "os": "linux"
      }
    }
  ]
}
```

## Implementation
For illustration, let's say that an image already exists in the registry:
- repository: `hello-world`
- digest: `sha256:80559bf80b44ce6be8234e6ff90a1ac34acbeb826903b02cfa0da11c82cbc042`
- tag: `v1.0`

The storage layout is as follows:

```
<root>
└── v2
    └── repositories
        └── hello-world
            └── _manifests
                └── revisions
                    └── sha256
                        └── 80559bf80b44ce6be8234e6ff90a1ac34acbeb826903b02cfa0da11c82cbc042
                            └── link
```

Now we push the OCI index with the signature reference in its config. The client would first upload the signature blob as a layer, followed by the index. The digest of the index is `sha256:90659bf80b44ce6be8234e6ff90a1ac34acbeb826903b02cfa0da11c82cbc042`.

```json
{
  "schemaVersion": 2.1,
  "mediaType": "application/vnd.oci.image.index.v2+json",
  "config": {
    "mediaType": "application/vnd.cncf.notary.config.v2+jwt",
    "digest": "sha256:2235d2d22ae5ef400769fa51c84717264cd1520ac8d93dc071374c1be49cc77c",
    "size": 1906
  },
  "manifests": [
    {
      "mediaType": "application/vnd.oci.image.manifest.v1+json",
      "digest": "sha256:80559bf80b44ce6be8234e6ff90a1ac34acbeb826903b02cfa0da11c82cbc042",
      "size": 7023,
      "platform": {
        "architecture": "ppc64le",
        "os": "linux"
      }
    }
  ]
}
```

At this point, the manifests storage layout will look as follows:

```
<root>
└── v2
    └── repositories
        └── hello-world
            └── _manifests
                └── revisions
                    └── sha256
                        ├── 80559bf80b44ce6be8234e6ff90a1ac34acbeb826903b02cfa0da11c82cbc042
                        │   └── link
                        │   └── signatures
                        │       └── sha256
                        │           └── 2235d2d22ae5ef400769fa51c84717264cd1520ac8d93dc071374c1be49cc77c
                        └── 90659bf80b44ce6be8234e6ff90a1ac34acbeb826903b02cfa0da11c82cbc042
                            └── link
```

Let's add another signature for the manifest `80559bf80b44ce6be8234e6ff90a1ac34acbeb826903b02cfa0da11c82cbc042`, where the digest of the index is `sha256:007170c33ebc4a74a0a554c86ac2b28ddf3454a5ad9cf90ea8cea9f9e75a153b`.

```json
{
  "schemaVersion": 2.1,
  "mediaType": "application/vnd.oci.image.index.v2+json",
  "config": {
    "mediaType": "application/vnd.cncf.notary.config.v2+jwt",
    "digest": "sha256:3335d2d22ae5ef400769fa51c84717264cd1520ac8d93dc071374c1be49cc88d",
    "size": 1906
  },
  "manifests": [
    {
      "mediaType": "application/vnd.oci.image.manifest.v1+json",
      "digest": "sha256:80559bf80b44ce6be8234e6ff90a1ac34acbeb826903b02cfa0da11c82cbc042",
      "size": 7023,
      "platform": {
        "architecture": "ppc64le",
        "os": "linux"
      }
    }
  ]
}
```

```
<root>
└── v2
    └── repositories
        └── hello-world
            └── _manifests
                └── revisions
                    └── sha256
                        ├── 80559bf80b44ce6be8234e6ff90a1ac34acbeb826903b02cfa0da11c82cbc042
                        │   └── link
                        │   └── signatures
                        │       └── sha256
                        │           ├── 2235d2d22ae5ef400769fa51c84717264cd1520ac8d93dc071374c1be49cc77c
                        │           │   └── link
                        │           └── 3335d2d22ae5ef400769fa51c84717264cd1520ac8d93dc071374c1be49cc88d
                        │               └── link
                        ├── 90659bf80b44ce6be8234e6ff90a1ac34acbeb826903b02cfa0da11c82cbc042
                        │   └── link
                        └── 007170c33ebc4a74a0a554c86ac2b28ddf3454a5ad9cf90ea8cea9f9e75a153b
                            └── link
```

Summarising various operations:
- Retrieving signatures for a particular manifest means enumerating all the links under its signatures folders.
- Storing a signature for a collection of manifests means adding a link file under the signatures folder for each manifest. The registry would verify basic  syntactic correctness, such as format of the JWT, but not trust.
- An index that references a signature in its config can be deleted. In this case, the link file to the signature should be removed from each referenced manifest's signatures folder.