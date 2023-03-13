# Credentials Rotation in Kubernetes – Putting Together the Puzzle Pieces

[![Netlify Status](https://api.netlify.com/api/v1/badges/30179840-3ac4-4fcb-940a-7175be505f88/deploy-status)](https://app.netlify.com/sites/talk-credentials-rotation/deploys)

Take me to the [slides](https://talks.timebertt.dev/credentials-rotation/)!

## About

This is a talk by [@rfranzke](https://github.com/rfranzke) and [@timebertt](https://github.com/timebertt) at [Cloud Native Rejekts 2023](https://cloud-native.rejekts.io/) in Amsterdam ([event schedule](https://cfp.cloud-native.rejekts.io/cloud-native-rejekts-eu-amsterdam-2023/talk/YGEAZF/)).

### Abstract

Every single Kubernetes cluster brings a plethora of credentials: server certificates, client certificates, ServiceAccount tokens, static tokens, etcd encryption keys, etc. But how do you manage them in a secure way?
Security best practices suggest using short-lived credentials wherever possible and frequently rotating static credentials everywhere else. What does this look like in practice when managing an entire fleet of clusters?
This talk puts together the puzzle pieces and presents how one can leverage Kubernetes primitives to securely handle all involved credentials in practice. It summarizes learnings that both cluster administrators and application developers can adopt to provide minimal-ops and disruption-free credentials management in Kubernetes.

### Description

Given the many distributed components inside a Kubernetes cluster that are connecting to each other, hardening and securing their communication is not as straightforward as one might hope. As a consequence, not every software in the Kubernetes ecosystem is following the best practices for managing credentials.
This talk shall inspire the audience on how such best practices (short-lived credentials, auto-rotation) can be implemented to improve the overall security of the ecosystem.
Apart from demystifying credentials management and rotation procedures in general, the listeners get insights into the Kubernetes community's transition from static ServiceAccount token secrets to projected tokens (along with interesting pitfalls).

## Presenting and Editing the Slides

Slides are built in Markdown using [reveal.js](https://revealjs.com/), packaged with [webpack](https://webpack.js.org/), and deployed with [netlify](https://www.netlify.com/).

### Prerequisites

Install a recent `node` version. Preferably, the one specified in [`.node-version`](./.node-version).

```bash
brew install node
```

### Present Locally

Perform a production build and serve the slides from the `dist` folder:

```bash
NODE_ENV=production npm run build
npm run serve
```

Important: Set `NODE_ENV=production` to yield the same build outputs as in production deploys to netlify.
If you don't set it, the QR will link to a local IP instead of the canonical URL, for example.

### Edit Locally

Run a dev server with hot-reload and open the slides in the browser:

```bash
npm start
```

Alternatively, use the preconfigured `start` run configuration for JetBrains IDEs.

Now, start editing the [content](./content) files.
When saving, slides are automatically rebuilt and refreshed in the browser.

> Note, that `npm start` doesn't write the output to `dist`.

### Build Locally

Run a full build and write output files to `dist`:

```bash
npm run build
```

Now, output files can be inspected in the `dist` folder.
Also, the slides can be served locally from the `dist` folder (no hot-reload):

```bash
npm run serve
```

Using the above will output non-minimized files.
Set `NODE_ENV=production` to enable minimization as it is done in netflify builds:

```bash
NODE_ENV=production npm run build
```

## Netlify Deploys

Netlify builds and publishes new commits to the `master` branch on https://talk-credentials-rotation.netlify.app/.

https://github.com/timebertt/talks contains a [netlify proxy configuration](https://github.com/timebertt/talks/blob/master/netlify.toml) to make the slides available at https://talks.timebertt.dev/credentials-rotation/.

The netlify site is configured to publish deploy previews for pull requests to the `master` branch and for pushes to arbitrary other branches.
