# <code>nuxctl<sup>&beta;</sup></code> - NuageX CLI control
[NuageX](https://nuagex.io) has been built to let new customers, partners and everyone else play with the best in class SDN solution by offering on-demand pre-defined, tested, and simple Nuage VSP environment.

But NuageX can be of a great help to the experienced Nuage Networks users as well! Advanced templates that are exposed to the verified partners and existing customers makes it possible to create complex labs for feature-testing, technical trainings, solution validation and more...

One of the tools that should radically simplify NuageX usage is <code>nuxctl<sup>&beta;</sup></code> - the CLI client to the NuageX API that offers the following major features:

* Define Labs configuration as-a-code and deploy them with a single click.
* Dump existing NuageX Lab into the configuration file that can be used by `nuxctl` to instantaneous recreation of that lab.

## Why nuxctl?
### Speed up the lab configuration & deployment time
The NuageX Web UI allows users to launch the existing templates and augment them by adding new servers, networks, services. If we are talking about deploying a lab from an existing template, then Web UI is a perfect choice to move forward. But things get a bit slow when a user wants to add multiple components to the base template.

Since for every, say, _server_ you need to type in the _flavor_, _name_, select the _interfaces_ and their _IP addresses_, the process quickly becomes monotonous.

With `nuxctl` the lab definition is a YAML file, so it can be written once, read many, effectively eliminating the need to type in things that are not changed.

### Lab definition as a code
By defining a Lab in a YAML structured file the Lab configuration becomes a piece of code, and here are the benefits of it:

* version controlled
* highly customizable and editable with a text editor of your choice
* shareable, testable and maintainable
* automation ready

### Redeploy instantly, anytime!
Its impossible to get a prolongation for a NuageX Lab. This is a necessity. And its not a deal breaker anymore! 

Remember, your Lab configuration is now stored in a human-readable file and can be re-instantiated anytime by `nuxctl`, while your overlay configuration can be defined and launched by various automation tools Nuage has to offer ([vspk-ansible](https://github.com/nuagenetworks/vspk-ansible), [vspk-python](https://github.com/nuagenetworks/vspk-python), [CATS](http://cats-docs.nuageteam.net) to name a few).

You no longer need to request a Lab prolongation, it is easily **redeployable**.

## Download
The CLI utility is distributed in a binary format (arch: amd64) for the following platforms:

* [linux](https://s3.eu-west-2.amazonaws.com/nuxctl/binaries/0.2.1/linux/nuxctl)
* [os x](https://s3.eu-west-2.amazonaws.com/nuxctl/binaries/0.2.1/darwin/nuxctl)
* [windows](https://s3.eu-west-2.amazonaws.com/nuxctl/binaries/0.2.1/windows/nuxctl.exe)

You can download the utility and make it ready to use in one line:
```bash
# example for os x
# the path /usr/local/bin should be in your $PATH if you want the binary to be runnable from the arbitrary current working directory.
cd /usr/local/bin && curl -O https://s3.eu-west-2.amazonaws.com/nuxctl/binaries/0.2.1/darwin/nuxctl && chmod a+x nuxctl
```

## Usage
`nuxctl` supports the following commands:
* **create-lab** - Create NuageX lab (environment) out of the lab configuration YAML file.
* **dump-lab** - Dump existing NuageX lab (environment) configuration in a YAML file.
* **list-templates** - Display NuageX Lab templates.
* **version** - Print the version number of nuxctl

Prior to jumping into each commands syntax and usage examples its important to mention that only an authenticated user can interact with NuageX API. This means that each command requires a proper authentication to happen.

### Authentication
`nuxctl` can receive credentials in the following forms:

1. Credentials are stored in a YAML file (aka _credentials file_)
2. Credentials are exported as the environment variables `NUX_USERNAME` and `NUX_PASSWORD`.

#### Credentials file
Credentials file should have the following key-value pairs:
```yaml
---
# example credentials file user_creds.yml
username: "nuagex"
password: "passw0rd"
```

A user should create such a file with theirs credentials and supply it to the appropriate commands using the `-c` flag. If the file is called `user_creds.yml` then the `-c` flag can be omitted, since `nuxctl` will assume this name by default.

```bash
# credentials file is passed explicitly
nuxctl create-lab -c my_creds.yml -l mylab.yml

# credentials file is omitted, assumed that user_creds.yml file is located in the same dir as nuxctl
nuxctl create-lab -l mylab.yml
```

#### Credentials via environment variables
Credentials from env vars will be read in case of `-c` flag missing and no default credentials file present (user_creds.yml) in the directory where nuxctl binary is located. In that case `nuxctl` will attempt to get the values of `NUX_USERNAME` and `NUX_PASSWORD` env. vars.

### Lab definition file
Another key concept used in various `nuxctl` commands is the _Lab definition/configuration file_. This file declares the NuageX Lab configuration expressed in YAML markup and is the source of information for the `create-lab` command.

Take a look at the following example:

```yml
name: nuxctl-example
expires: "2018-08-28T17:13:11.278Z"
template: "5980ee745a38da00012d158d"

sshKeys:
  - name: cats
    key: "ssh-rsa AAAAB3Nza [OMITTED FOR BREVITY] gzfTp"

####################################
######    S E R V I C E S
####################################
services:
  - name: vsd
    type: public
    port: 8443
    urlScheme: https
    protocols:
      - tcp
    destination:
      port: 8443
      address: "10.0.0.2"

####################################
######    N E T W O R K S
####################################
networks:
  - name: nsg1-lan1
    cidr: "172.17.1.0/24"
    dns: "10.0.0.1"
    dhcp: false
  - name: wan-dummy
    cidr: "10.10.10.0/24"
    dhcp: false

####################################
######    S E R V E R S
####################################
servers:
  - name: nsg1
    flavor: "m1.small"
    image: "nux_nsg_5.2.3"
    interfaces:
      - network: private
        address: "10.0.0.11"
        index: 0
      - network: wan-dummy
        address: "10.10.10.2"
        index: 1
      - network: nsg1-lan1
        address: "172.17.1.2"
        index: 2
```

Lets go over some notable blocks here:

* `expires`: expiration date, must be no longer than 30 days since the date of the deployment.
* `template`: ID of the _base template_. You can get the list of available templates and theirs ID by using `nuxctl list-templates` command.
* `sshKeys`: put a name and the public key that you want to use with this lab. Multiple keys can be provided.

> **IMPORTANT:** If you see that your lab is queued for creation by `nuxctl`, but nothing is seen in the web ui, then first check your key, make sure that it is pasted without errors or truncation.

#### Services
In the Services block the port-forwarding rules are defined. These rules are configured in the jumpbox vm of your Lab.

#### Networks
This block holds Networks definitions that will be created for your Lab. Bear in mind, that the default `private: 10.0.0.0/24` network is always created in the background. **DO NOT INCLUDE** the default private network in your YAML file since it will cause the deployment to fail.

#### Servers
Servers block contains configuration for the VMs you want to add. Make sure to correctly refer to the network names, IP addresses and indexes of the network interfaces. The rest can be retrieved from the Web UI.

### create-lab
To create (aka deploy) a NuageX Lab from the Lab definition file use the `create-lab` command:
```
Usage:
  nuxctl create-lab [flags]

Flags:
  -c, --credentials          Path to the user credentials file (default "user_creds.yml")
  -l, --lab-configuration    Path to the Lab configuration file (default "lab.yml")
```

Note the default names used for credentials and lab-configuration file names. A successful execution of a `create-lab` command looks as follows. Credentials file and lab configuration file are passed explicitly, if they would have been omitted, default names would have been assumed by nuxctl.

```
bash-3.2$ ./nuxctl create-lab -c demo_user_creds.yml -l demo_lab.yml 
Loading user credentials from 'demo_user_creds.yml' file
Logging 'hellt' user in...
User 'hellt' logged in...
Loading lab configuration from 'demo_lab.yml' file
Lab ID 5b990011e8b2618598c92f8c has been successfully queued for creation! Request ID c0297bea-edf9-42e4-a303-9af563ce089f
```

Mark the Lab and Request IDs reported in the last line of the output, if you see that Labs failed to create or did not appear at all, reach out to nuxctl maintainer providing these two identifiers for troubleshooting.

### dump-lab
In case a user wants to save a running NuageX Lab (that has been created via Web UI) in a configuration file the `dump-lab` command can be used:

```
Usage:
  nuxctl dump-lab [flags]

Flags:
  -c, --credentials    Path to the user credentials file. (default "user_creds.yml")
  -f, --file           Path to the local YAML file that will receive lab configuration. (default "dumplab.yml")
  -i, --lab-id         Lab ID. Seen as the variable portion in the lab hostname.
```

Again, credentials should be supplied along with the file name to save the Lab configuration. The Lab itself is referenced by its ID, which can be seen in the hostname portion of a lab:

> ![pic 1](https://gitlab.com/rdodin/pics/wikis/uploads/24527d2b40858051d82ebf1edc1d2e67/image.png)

For reference' sake lets dump the lab created previously in a file named `dump-example.yml`:

```
$ ./nuxctl dump-lab -c demo_user_creds.yml -i 5b95785ee8b2619d3bc92d5a -f dump-example.yml
Loading user credentials from 'demo_user_creds.yml' file
Logging 'hellt' user in...
User 'hellt' logged in...
Retrieving NuageX Lab configuration...
Parsing Lab configuration...
Writing Lab configuration to 'dump-example.yml' file...
Lab configuration has been successfully written to 'dump-example.yml' file!
```

Note, that the lab configuration retrieved with the `dump-lab` command will have the _private_ network commented out. This is done on purpose, to let you reuse the dumped configuration without manually commenting out this block.

Another intervention this command makes is that the original _template ID_ that was used to build a lab will appear in the YAML file as the _Empty Template ID (5980ee745a38da00012d158d)_. This is a necessary change to make the dumped configuration be immediately reusable by the `create-lab` command.

### list-templates
NuageX provides a long list of Lab Templates ranging from simple Nuage VNS/VCS installations to the complex partner integrations. Every NuageX Lab should be based on some Template.

`nuxctl` expects to see a `Template ID` in the Lab definition file to understand which template to use. At the moment of this writing, Templates listed by this URL https://experience.nuagenetworks.net/app/templates do not expose the Template ID, therefore a separate command has been added to help with extracting the Template ID by the Template name.

```
Usage:
  nuxctl list-templates [flags]

Flags:
  -c, --credentials   Path to the user credentials file. (default "user_creds.yml")
```

And its easy to get the sorted by a name list of Templates:
```
$ ./nuxctl list-templates -c demo_user_creds.yml 
Loading user credentials from 'demo_user_creds.yml' file
Logging 'hellt' user in...
User 'hellt' logged in...
Retrieving NuageX Lab Templates...

ID                        Name
------------------------  ------------------------
5980ee745a38da00012d158d  Empty template "For Customization"
5a5e8fa6ff408b00010a3f7b  Nuage Networks 3.2R6 - Liberty Fusion Layer Infinity
5a5e9159ff408b00010a3f98  Nuage Networks 4.0R3 - Generic Hardware VTEP Demo
OMITTED FOR BREVITY
```

## Changelog
Changelog is [here](CHANGELOG.md)

## Author
Roman Dodin - roman.dodin@nokia.com