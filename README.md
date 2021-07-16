<p align="center" margin="20 0"><a href="https://horusec.io/">
    <img src="assets/horusec_logo.png" alt="logo_header" width="65%" style="max-width:100%;"/></a></p>

<p align="center">
    <a href="https://github.com/ZupIT/horusec-devkit/pulse" alt="activity">
        <img src="https://img.shields.io/github/commit-activity/m/ZupIT/horusec-devkit"/></a>
    <a href="https://github.com/ZupIT/horusec-devkit/graphs/contributors" alt="contributors">
        <img src="https://img.shields.io/github/contributors/ZupIT/horusec-devkit"/></a>
    <a href="https://github.com/ZupIT/horusec-devkit/actions/workflows/lint.yml" alt="lint">
        <img src="https://img.shields.io/github/workflow/status/ZupIT/horusec-devkit/Lint?label=lint"/></a>
    <a href="https://github.com/ZupIT/horusec-devkit/actions/workflows/tests.yml" alt="test">
        <img src="https://img.shields.io/github/workflow/status/ZupIT/horusec-devkit/Test?label=test"/></a>
    <a href="https://github.com/ZupIT/horusec-devkit/actions/workflows/security.yml" alt="security">
        <img src="https://img.shields.io/github/workflow/status/ZupIT/horusec-devkit/Security?label=security"/></a>
    <a href="https://github.com/ZupIT/horusec-devkit/actions/workflows/coverage.yml" alt="coverage">
        <img src="https://img.shields.io/github/workflow/status/ZupIT/horusec-devkit/Coverage?label=coverage"/></a>
    <a href="https://opensource.org/licenses/Apache-2.0" alt="license">
        <img src="https://img.shields.io/badge/license-Apache%202-blue"/></a>
</p>

# Horusec Development Kit

This repository has the idea of centralizing all the reused horusec codes in one place. 
It is also where are some abstractions that we use in order to simplify development and testing.

## Structure

### Entities

All entities that are shared by multiple microservices are kept here.
An example is our analysis and vulnerability structs, which are used from the cli to the web services.

### Enums

It refers to all the constants shared between the services.
An example is the vulnerability severity constant.

### Service

Here you will find abstractions from libraries that we use, very similar to the next topic.
The big difference between service and utils, is that here they all need an instance or
connection due to some requirement.
An example is the abstraction of rabbit mq library for go, that simplifies testing and development.

### Utils

To finish, the utils refers to an abstraction which instantiation or connection is not necessary. 
Just import and use it, simply and quickly.
An example is our abstraction from the logrus library, that we adapted to make it more comfortable to use.

## Communication

We have a few channels for contact, feel free to reach out to us at:

- [GitHub Issues](https://github.com/ZupIT/horusec-devkit/issues)
- [Zup Open Source Forum](https://forum.zup.com.br)

## Contributing

Feel free to use, recommend improvements, or contribute to new implementations.

If this is our first repository that you visit, or would like to know more about Horusec, 
check out some of our other projects.

- [Horusec CLI](https://github.com/ZupIT/horusec)
- [Horusec Platform](https://github.com/ZupIT/horusec-platform)
- [Horusec Engine](https://github.com/ZupIT/horusec-engine)
- [Horusec Operator](https://github.com/ZupIT/horusec-operator)
- [Horusec Admin](https://github.com/ZupIT/horusec-admin)
- [Horusec VsCode](https://github.com/ZupIT/horusec-vscode-plugin)

This project exists thanks to all the contributors. You rock! ❤️🚀
