<h1 align="center"> requisition</h1>

<p align="center">
  <a href="https://github.com/neggles/requisition/releases" rel="nofollow">
    <img alt="GitHub release (latest SemVer including pre-releases)" src="https://img.shields.io/github/v/release/neggles/requisition?include_prereleases">
  </a>

  <a href="https://github.com/neggles/requisition/actions/workflows/release.yaml" rel="nofollow">
    <img src="https://github.com/neggles/requisition/actions/workflows/release.yaml/badge.svg" alt="goreleaser" style="max-width:100%;">
  </a>

  <a href="https://pkg.go.dev/github.com/neggles/requisition" rel="nofollow">
    <img src="https://pkg.go.dev/badge/github.com/neggles/requisition.svg" alt="Go reference" style="max-width:100%;">
  </a>

  <a href="https://github.com/gojp/goreportcard/blob/master/LICENSE" rel="nofollow">
    <img src="https://img.shields.io/badge/license-GPLv3-blue.svg" alt="License GPLv3" style="max-width:100%;">
  </a>

  <br/>

  <a href="https://codecov.io/gh/neggles/requisition" >
    <img src="https://codecov.io/gh/neggles/requisition/branch/main/graph/badge.svg?token=CLP6KW4QLK"/>
  </a>

  <a href="https://github.com/neggles/requisition/actions/workflows/codeql.yaml" rel="nofollow">
    <img src="https://github.com/neggles/requisition/actions/workflows/codeql.yaml/badge.svg" alt="codeql" style="max-width:100%;">
  </a>

  <a href="https://goreportcard.com/report/github.com/neggles/requisition" rel="nofollow">
    <img src="https://goreportcard.com/badge/github.com/neggles/requisition" alt="Go report card" style="max-width:100%;">
  </a>
</p>
<br/>

a multithreaded http downloader


## Install

You can install the pre-compiled binary (in several ways), use Docker or compile from source (when on OSS).

Below you can find the steps for each of them.

<details>
  <summary><h3>homebrew tap</h3></summary>

```bash
brew install neggles/tap/requisition
```

</details>

<details>
  <summary><h3>apt</h3></summary>

```bash
echo 'deb [trusted=yes] https://apt.fury.io/neggles/ /' | sudo tee /etc/apt/sources.list.d/neggles.list
sudo apt update
sudo apt install requisition
```

</details>

<details>
  <summary><h3>yum</h3></summary>

```bash
echo '[neggles]
name=Gemfury neggles repository
baseurl=https://yum.fury.io/neggles/
enabled=1
gpgcheck=0' | sudo tee /etc/yum.repos.d/neggles.repo
sudo yum install goreleaser
```

</details>

<details>
  <summary><h3>deb, rpm and apk packages</h3></summary>
Download the .deb, .rpm or .apk packages from the [release page](https://github.com/neggles/requisition/releases) and install them with the appropriate tools.
</details>

<details>
  <summary><h3>go install</h3></summary>

```bash
go install github.com/neggles/requisition@latest
```

</details>

<details>
  <summary><h3>from the GitHub releases</h3></summary>

Download the pre-compiled binaries from the [release page](https://github.com/neggles/requisition/releases) page and copy them to the desired location.

```bash
$ VERSION=v1.0.0
$ OS=Linux
$ ARCH=x86_64
$ TAR_FILE=requisition_${OS}_${ARCH}.tar.gz
$ wget https://github.com/neggles/requisition/releases/download/${VERSION}/${TAR_FILE}
$ sudo tar xvf ${TAR_FILE} requisition -C /usr/local/bin
$ rm -f ${TAR_FILE}
```

</details>

<details>
  <summary><h3>manually</h3></summary>

```bash
$ git clone github.com/neggles/requisition
$ cd requisition
$ go generate ./...
$ go install
```

</details>

## Contribute to this repository

This project adheres to the Contributor Covenant [code of conduct](https://github.com/neggles/requisition/blob/main/.github/CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. We appreciate your contribution. Please refer to our [contributing](https://github.com/neggles/requisition/blob/main/.github/CONTRIBUTING.md) guidelines for further information.
