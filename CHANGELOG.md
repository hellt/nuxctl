## 0.3.1

* Accepting non secure connections to NuageX API is now switched off to allow nuxctl to run on WSL.

## 0.3.0

* `dump-lab` now comments out the entities that are found in the template referenced in this lab. This makes it possible to dump a lab and immediately use the produced file in the `create-lab` command.
* added command `dump-template`. This commands dumps the template definition to a YAML file.

## 0.2.1

* fixed debug printout of env vars

## 0.2.0

* added ability to pass credentials via env vars. Check [README.md](README.md)
* `dump-lab` command substitutes template ID to the ID of the empty template to allow the file be suitable for later deployment without changes.
* Lab ID and Request ID are emitted by `create-lab` command.
* improved error handling

## 0.1.0

* added CLI framework to support multiple commands
* added `create-lab`, `dump-lab`, `list-templates`, `version` commands

## 0.0.1
Initial version.