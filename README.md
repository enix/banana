# The banana project

## Introduction

Banana is an enterprise-grade, fully secured, backuping system.

It has been developed as an alternative to [backuppc](http://backuppc.sourceforge.net) which is quite powerful ... but may put your content at risk.

The Banana project aims to backup thousands of nodes, without requiring direct (centralized) access to them.
All backups are ciphered prior to their push to a storage backend.
All credentials or keys are managed through the usage of [vault](https://www.vaultproject.io/) from HashiCorp.

We took extra hours to make sure that both the installation and usage of banana is as simple as it can be from an ergonomy standpoint. Please feel free to suggest any complex area we might have missed.

The current status of the project is *alpha*.

## Features

* Privilege separation
	* Storage only sees encrypted backups
	* Central components do not have access to storage and do not have access to any credentials
	* vault do not have access to storage, neither to nodes
	* Nodes can reach storage, vault and central components, the opposite is **not** true
* Minimal security risks on nodes
	* Nodes push their status (excluding any credential) to a centralized component
	* Nodes get temporary backup ciphering key from *vault*
	* Nodes encrypts their backup prior to pushing them.
* All Backups are monitored, including alerting, through central component
* A Web UI provides consolidated information about all backup jobs and status
* Support for various backup types through a plugin based implementation
	* FileSystem
	* Databases (Mysql, Etcd, ...)
* Support for multiple storage backends: S3, Swift, NFS, Samba, ...
* Encryption-at-rest on storage backends

## Project composition

Banana is componed of :
* *bananaui*: A centralised "watch tower"
* *bananactl*: An agent to be launched on any node to backup
* *bananadm*: An administrative configuration helper

Banana depends on other components in order to work :
* At least one storage backend
* A running vault instance (although bananadm can set it up for you)

## Installation

See the [install guide](https://gitlab.enix.io/products/banana/blob/master/docs/INSTALLATION.md) to install banana.

## Upgrade

See the [upgrade guide](https://gitlab.enix.io/products/banana/blob/master/docs/UPGRADE.md) to upgrade banana.

## Plugins

See the [plugin reference](https://gitlab.enix.io/products/banana/blob/master/docs/PLUGINS.md) to install or implement plugins.
