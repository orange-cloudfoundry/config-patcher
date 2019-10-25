# config-patcher

small utility to patch config file in different format by using yaml patch expression which can be found on this link: https://github.com/krishicks/yaml-patch .

Config-patcher support theses formats as config input:
- json
- yaml
- toml

## Installation 

**Config-patcher has been made to work during a bosh lifecycle and should be use through its boshrelease: https://github.com/orange-cloudfoundry/config-patcher-boshrelease**

You can use by cli, by using `go get https://github.com/orange-cloudfoundry/config-patcher`

```
Usage of ./config-patcher:
  -patch string
    	Set in glob format where to find rules for patching config files (default "/var/vcap/jobs/*/config-patcher/*.yml")
```

## Patch format

Create a yaml file in this format (which is accessible which -patch flag):

```yaml
- config_file: <config-file> # config file in input which will be patched
  config_type: <json | yaml | toml> # this is actually not mandatory but you could need to set explicitly type of your config file
  patches:
    - op: <add | remove | replace | move | copy | test>
      from: <source-path> # only valid for the 'move' and 'copy' operations
      path: <target-path> # always mandatory
      value: <any-yaml-structure> # only valid for 'add', 'replace' and 'test' operations
```