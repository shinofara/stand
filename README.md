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
zip_name: "zipfile's prefix"
life_cycle: 12
```

Unit of `life_cycle` is the number of files .

# License

MIT: [LICENSE](LICENSE)
