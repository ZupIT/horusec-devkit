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


# **Horusec Development Kit**

This repository has all the reused Horusec codes in one place. 

## **How does DevKit work?**
DevKit is the repository where there are some abstractions Horusec's team use to simplify development and testing.  For example:
-  If you want a code to help you make an HTTP request or if you need a database, you can find it here. 

It is reusable codes from other projects, like CLI, Engine, Platform, and Operator. See below DevKit's structure:

### **Entities**

We keep all entities shared by multiple microservices here. For example, analysis and vulnerability structs, which are used from the CLI to the web services.

### **Enums**

Enums refers to all the constants shared between the services. For example, the vulnerability severity constant.

### **Service**

You will find here some abstractions from the libraries we use. The difference between service and utils is that here they need an instance or a connection to some requirement.
For example, the abstraction of the RabbitMQ library for Go, which simplifies testing and development.

### **Utils**

The utils refer to an abstraction in which instantiation or connection is not necessary. 
Just import and use it, simply and quickly. For example, the abstraction from the Logrus library was adapted to make it more comfortable to use.


## **Documentation**

For more information about Horusec, please check out the [**documentation**](https://horusec.io/docs/).

## **Contributing**

If you want to contribute to this repository, access our [**Contributing Guide**](https://github.com/ZupIT/charlescd/blob/main/CONTRIBUTING.md). 
And if you want to know more about Horusec, check out some of our other projects:


- [**Horusec CLI**](https://github.com/ZupIT/horusec)
- [**Horusec Platform**](https://github.com/ZupIT/horusec-platform)
- [**Horusec Engine**](https://github.com/ZupIT/horusec-engine)
- [**Horusec Operator**](https://github.com/ZupIT/horusec-operator)
- [**Horusec Admin**](https://github.com/ZupIT/horusec-admin)
- [**Horusec VsCode**](https://github.com/ZupIT/horusec-vscode-plugin)

## **Community**
Feel free to reach out to us at:

- [**GitHub Issues**](https://github.com/ZupIT/horusec-devkit/issues)
- [**Zup Open Source Forum**](https://forum.zup.com.br)


This project exists thanks to all the contributors. You rock! ❤️🚀
