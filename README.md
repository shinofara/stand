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
target: "/path/to/target/dir"
output: "/path/to/output/dir"
life_cycle: 12
compression:
  prefix: sample
  format: zip # or tar
```

or

```
# /path/to/config.yml
- target: "/path/to/target/dir"
  output: "/path/to/output/dir"
  life_cycle: 12
  compression:
    prefix: sample
    format: zip # or tar
- target: "/path/to/target/dir"
  location: "s3"
  output: "/path/to/output/dir"
  life_cycle: 12
  s3:
    access_key_id: XXXXXXXXXXXXXXXXXXX
    secret_access_key: XXXXXXXXXXXXXXXXXXX
    region: "ap-northeast-1"
    bucket_name: "backupbacket"    
  compression:
    prefix: sample
    format: zip # or tar    
```

Unit of `life_cycle` is the number of files .

# Contribution

1. Fork it 
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create new Pull Request

# License

MIT: [LICENSE](LICENSE)
