# <code>nuxctl<sup>&alpha;</sup></code> - NuageX CLI control
[NuageX](https://nuagex.io) has been built to let our customers and partners taste the best in class SDN solution by offering pre-defined, tested, and simple templates.

But NuageX can be (and already is) of a great help to all of us, the engineers. Currently the NuageX team is working hard to see what features can be unleashed and what additional tools we have to build to let our internal users satisfy their needs and boost the overall performance.

One of the tools that should radically simplify NuageX usage in complex labs is <code>nuxctl<sup>&alpha;</sup></code> - the CLI client to the NuageX API that allows end-users to define Labs as-a-code and launch them with a single click.

## Why `nuxctl`?
### Speed up the deployment time
The NuageX Web UI allows users to launch the existing templates and augment them by adding new servers, networks, services. If we are talking about deploying a lab from an existing template, then Web UI is a perfect choice to move forward. But things get a bit slowly when a user wants to add multiple components to a _base template_.

Since for every, say, _server_ you need to type in the flavor, name, select the interfaces and their IP addresses, the process becomes quite routine if you need to repeat it more than once.

With `nuxctl` the lab definition is a YAML file, so it can be written once, read many, effectively eliminating the need to type in things that are not changed.

### Lab definition as a code
By defining a Lab in a YAML structured file we are moving the NuageX labs into the buzzy area of `X-as-a-code`. Effectively, your Lab is now a piece of code, and here are the benefits of it:

* version controlled
* highly customizable and editable with a text editor of your choice
* shareable, testable and maintainable
* automation ready

### Redeploy instantly, anytime!
Its impossible to get a prolongation for a NuageX Lab. _Sad, but true_.

But imagine that your Lab structure defined in a file and launched by `nuxctl`, while your overlay configuration is defined in another file and handled by [CATS](http://cats-docs.nuageteam.net). You see where are we going?.. You no longer need to request to prolong a Lab, its easily **redeployable**.

## Usage
Finally, how to use the alpha version of `nuxctl`?

The CLI utility is distributed as a single binary for the following platforms:

* linux
* os x
* windows

You can download the binaries [here](https://nokia-my.sharepoint.com/:f:/p/roman_dodin/ErXXISWPmkpAo3Er3Qls5ksBie4OosAyP45Dt6GCVOv44g?e=BQyY1i).

The utility supports the following operation:
* login a user
* create a lab out of a definition file for a logged in user

The basic usage example is as follows:

```bash
# -u <credentials_file_path>
# -l <lab_file_path>
nuxctl -u my_credentials.yml -l vns_lab.yml

# output
Loading user credentials from 'user_creds.yml' file
Logging 'hellt' user in...
User 'hellt' logged in...
Loading lab configuration from 'lab.yml' file
Sending request to create a lab...
Lab has been successfully queued for creation!
```

Here the two files are used by `nuxctl`, one reads credentials from a file, the other expects to find a lab definition file.

If `nuxctl` is executed without the `-u` and `-l` flags supplied, then the defaults files path will be assumed:

* default path for user credentials file: `./user_creds.yml`
* default path for lab file: `./lab.yml`

## Credentials file
Credentials file accepts username and password in the plain text format:
```yml
---
username: "nuagex"
password: "passw0rd"
```

## Lab definition file
Lab definition file is a bit more complex, since it deals with the feature-rich NuageX API.
Here is a simple Lab, expressed in YAML that is understandable to engineers and `nuxctl`:

```yml
name: nuxctl-example
reason: "Nux'em all"
expires: "2018-08-28T17:13:11.278Z"
template: 5980ee745a38da00012d158d

sshKeys:
  - name: cats
    key: "ssh-rsa AAAAB3Nza[OMITTED FOR BREVITY]gzfTp"

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
        address: "255.255.255.1"
        index: 1
      - network: nsg1-lan1
        address: "172.17.1.1"
        index: 2
```

Lets go over some notable blocks here. In the very beginning you need to supply the basic Lab information. Important fields there are:

* `expires`: expiration date, must be no longer than 30 days since the date of the deployment.
* `template`: ID of the _base template_. You can get this ID by checking the JSON response via the Developers Console. For instance, [this is the list](https://pastebin.com/5ZpKQcfZ) of ID-Name pairs for the templates available by the date of 2018-08-11.
* `sshKeys`: put a name and the public key that you want to use with this lab. Multiple keys can be provided.

> **IMPORTANT:** the public key should have been added to your NuageX profile before you use it in the lab definition file `nuxctl`. If you see that your lab is queued for creation by `nuxctl`, but nothing is seen in the web ui, then first check your key, make sure that it is pasted in full.

### Services
In the Services block you can define the port-forwarding rules that will be configured within your lab.

### Networks
Here you define what networks you want to create. Bear in mind, that the default `private: 10.0.0.0/24` network is always created in the background.

### Servers
Servers block is self explanatory. Make sure to correctly refer to the network names and indexes of the network interfaces. The rest of the information can be grepped from the Web UI.

Now if you want to see some real use case where `nuxctl` is at its best - consider [the following Lab definition](https://pastebin.com/BgDrY7N2) I use lately. Its literally impossible to recreate that many components via Web UI.

## Disclaimer
`nuxctl` is young and fast-coded, that is why its been marked with &alpha; (alpha). This means that the functionality is limited, bugs might appear in unexpected places and code is not hardened.

But this is alright, its just a start and I want to kindly ask you to report any bugs seen directly to me - roman.dodin@nokia.com. You can also use this email to suggest improvements and additional features.

## Author
Roman Dodin - roman.dodin@nokia.com