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

## **Table of contents**
### 1. [**About**](#about)
### 2. [**Usage**](#usage)
>#### 2.1. [**How does DevKit work?**](#how-does-devkit-work)
### 3. [**Features**](#features)
### 4. [**Documentation**](#documentation)
### 5. [**Contributing**](#contributing)
### 6. [**License**](#license)
### 7. [**Community**](#community)

## **About**
This repository has all the reused Horusec codes in one place. 

## **Usage**

### **How does DevKit work?**
DevKit is the repository where there are some abstractions Horusec's team use to simplify development and testing. For example:
-  If you want a code to help you make an HTTP request or if you need a database, you can find it here. 

It is reusable codes from other projects, like CLI, Engine, Platform, and Operator. See below DevKit's structure:

## **Features**
### **Entities**

We keep all entities shared by multiple microservices here. For example, analysis and vulnerability structs, which are used from the CLI to the web services.

### **Enums**

Enums refers to all the constants shared between the services. For example, the vulnerability severity is a constant.

### **Service**

You will find here some abstractions from the libraries we use. The difference between service and utils is that here they need an instance or a connection to some requirement. For example, the abstraction of the RabbitMQ library for Go, which simplifies testing and development.

### **Utils**

The utils refer to an abstraction in which instantiation or connection is not necessary. You just need to import and use it,  for example, the abstraction from the Logrus library was adapted to make it more comfortable to use.


## **Documentation**

For more information about Horusec, please check out the [**documentation**](https://horusec.io/docs/).


## **Contributing**

If you want to contribute to this repository, access our [**Contributing Guide**](https://github.com/ZupIT/horusec-devkit/blob/main/CONTRIBUTING.md). 

### **Developer Certificate of Origin - DCO**

 This is a security layer for the project and for the developers. It is mandatory.
 
 Follow one of these two methods to add DCO to your commits:
 
**1. Command line**
 Follow the steps: 
 **Step 1:** Configure your local git environment adding the same name and e-mail configured at your GitHub account. It helps to sign commits manually during reviews and suggestions.

 ```
git config --global user.name “Name”
git config --global user.email “email@domain.com.br”
```
**Step 2:** Add the Signed-off-by line with the `'-s'` flag in the git commit command:

```
$ git commit -s -m "This is my commit message"
```

**2. GitHub website**
You can also manually sign your commits during GitHub reviews and suggestions, follow the steps below: 

**Step 1:** When the commit changes box opens, manually type or paste your signature in the comment box, see the example:

```
Signed-off-by: Name < e-mail address >
```

For this method, your name and e-mail must be the same registered on your GitHub account.

## **License**
[**Apache License 2.0**](https://github.com/ZupIT/horusec-devkit/blob/main/LICENSE).

## **Community**
Do you have any question about Horusec? Let's chat in our [**forum**](https://forum.zup.com.br/).


This project exists thanks to all the contributors. You rock! ❤️🚀

