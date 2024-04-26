---
aliases: [README]
tags: []
title: README
linter-yaml-title-alias: README
date created: Wednesday, April 17th 2024, 4:11:40 pm
date modified: Thursday, April 18th 2024, 8:19:25 am
---

## Distributed Key Generation (DKG)

Package `dkg` is intended to contain implementations of distributed key generation (DKG) protocols.
Besides the DKG protocol we introduced in `github.com/di-dao/core/crypto/pkg/crypto/tecdsa/README.md`, currently
package `dkg` contains the following distributed key generation implementations.

-   Gennaro DKG: an adapted version [[overleaf]](https://www.overleaf.com/project/60915c0df1d6917f5cde6657) of
    DKG by Gennaro et al. [[GennaroDKG]](https://link.springer.com/content/pdf/10.1007/s00145-006-0347-3.pdf). (We call it
    GennaroDKG for convenience in the following context.)

-   FROST DKG: the distributed key generation protocol used in [FROST tSchnorr signature](https://tools.ietf.org/pdf/draft-komlo-frost-00.pdf). We also
    have its [pseudocode write-up](https://www.overleaf.com/read/nvmyjwsnbrwj). We call it FROST DKG in the following context.