# gh-requested-prs

`gh-requested-prs` is a GitHub CLI extension that opens all pull requests where you have been requested for review in your browser.

## Installation

1. Make sure you have [GitHub CLI](https://cli.github.com/) installed and configured.

2. Download and install the `gh-requested-prs` extension.

```shell
$ gh extension install KazukiHayase/gh-requested-prs
```

## Usage

To open all pull requests where you have been requested for review, simply run the following command:

```shell
gh requested-prs
```

This will retrieve the list of pull requests from GitHub GraphQL API, filter them based on your review request, and open each pull request URL in your default browser.

## Configuration

No additional configuration is required. The extension uses the default GitHub CLI configuration for authentication and API access.
