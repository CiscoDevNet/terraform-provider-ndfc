# How to Contribute

Thanks for your interest in contributing to this Terraform provider! Here are a few general guidelines on contributing and
reporting bugs that we ask you to review. Following these guidelines helps to communicate that you respect the time of
the contributors managing and developing this open source project. In return, they should reciprocate that respect in
addressing your issue, assessing changes, and helping you finalize your pull requests. In that spirit of mutual respect,
we endeavor to review incoming issues and pull requests within 10 days, and will close any lingering issues or pull
requests after 60 days of inactivity.

Please note that all of your interactions in the project are subject to our [Code of Conduct](/CODE_OF_CONDUCT.md). This
includes creation of issues or pull requests, commenting on issues or pull requests, and extends to all interactions in
any real-time space e.g., Slack, Discord, etc.

## Table Of Contents

- [Reporting Issues](#reporting-issues)
- [Development](#development)
  - [Building the Provider](#building-the-provider)
  - [Acceptance Tests](#acceptance-tests)
- [Sending Pull Requests](#sending-pull-requests)
- [Other Ways to Contribute](#other-ways-to-contribute)

## Reporting Issues

Before reporting a new issue, please ensure that the issue was not already reported or fixed by searching through our
[issues list](https://github.com/netascode/terraform-provider-ndfc/issues).

When creating a new issue, please be sure to include a **title and clear description**, as much relevant information as
possible, and, if possible, a test case.

**If you discover a security bug, please do not report it through GitHub. Instead, please see security procedures in
[SECURITY.md](/SECURITY.md).**

## Development

### Building the Provider

1. Clone the repository
2. Enter the repository directory
3. Build the provider using the Go `install` command:

```shell
go install
```

### Acceptance Tests

In order to run the full suite of Acceptance tests, run `run_accept_tests.sh`. Make sure that the test configuration is updated with the NDFC details

Note: Acceptance tests needs a real NDFC instance and it creates real resources

```shell
run_accept_tests.sh 
```

## Sending Pull Requests

At this time we are not accepting Pull Requests from outside for new resources and data sources as we are in initial stages of development.
We will selectively accept PRs related to adding new tests, minor bug fixes, documentation etc

## Other Ways to Contribute

We welcome anyone that wants to contribute to this Terraform provider to triage and reply to open issues to help troubleshoot
and fix existing bugs. Here is what you can do:

- Help ensure that existing issues follows the recommendations from the _[Reporting Issues](#reporting-issues)_ section,
  providing feedback to the issue's author on what might be missing.
- Review existing pull requests, and testing patches against real infrastructure.
- Write a test, or add a missing test case to an existing test.

Thanks again for your interest on contributing to this Terraform provider!

:heart:
