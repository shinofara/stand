Stand
============

[![CircleCI](https://circleci.com/gh/shinofara/stand.svg?style=svg)](https://circleci.com/gh/shinofara/stand)

# Description

Stand is a tool to perform directory backup and backup's generation management  and restore.

# Installation

```
$ go install github.com/shinofara/stand
```

# Usage

```
$ stand -conf /path/to/config.yml
```

# Configration


```
# /path/to/config.yml
type: dir # or file
path: "/path/to/target/1"
compression:
  prefix: "prefix1"
  format: zip # or tar
storages:
  - type: local
    path: "/path/to/output/1"
    life_cycle: 1
```

or

```
# /path/to/config.yml
- type: file
  path: "/path/to/file/path"
  storages:
    - path: "/path/to/local/directory"
      life_cycle: 1
- type: dir
  path: "/path/to/compress/target/directory"
  compression:
    prefix: sample # prefix_YYYYMMDDHHMiss.zip
    format: zip # or tar
  storages:
      type: s3
      path: "/path/to/s3/path"
      life_cycle: 2
      s3:
        access_key_id: AKXXXXXXXXXXXX
        secret_access_key: yXxxxxxxXXXXXXXXXXXXXXXX
        region: "ap-northeast-1"
        bucket_name: "backup_bucket_name"
```

## S3 Config pattern

### Pattern set in YAML

```
- .
  other settings
  .
  .
  storages:
    - s3:
        access_key_id: AKXXXXXXXXXXXX
        secret_access_key: yXxxxxxxXXXXXXXXXXXXXXXX
        region: "ap-northeast-1"
        bucket_name: "backup_bucket_name"
```

### Pattern to be set in the environment variable

```
$ export AWS_ACCESS_KEY_ID=AKIXXXXXXXXXXXXXX
$ export AWS_SECRET_ACCESS_KEY=xxxxxxxxxxxxxxxxxxxxxxxxxxxxx
$ export AWS_DEFAULT_REGION=ap-northeast-1
```

Unit of `life_cycle` is the number of files .

# Policy of release versioning

v `Majors` . `Minors` . `Patches`

### Patches

Patch releases are defined as bug, performance, and security fixes. They do not change or add any public interfaces. They do not change the expected behavior of a given interface. They are meant to correct behavior when it doesn't match the documentation, without also introducing a change that makes seamless upgrades impossible.

### Minors

Minors are the addition and refinements of APIs or subsystems. They do not change APIs or introduce backwards compatible breaking changes. These are mostly completely additive releases.

### Majors

Majors contain changes in behavior that could potentially break code that worked in prior releases.

# Contribution

1. Fork it 
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create new Pull Request

# License

MIT: [LICENSE](LICENSE)
