# Infogen
Infogen is an information generator. This is used for generating an email alias based on your own domain, a name and an address.

This currently uses Rapid API for the name and address but I would like to move away from an API at some point

## Install
```yaml
$ git clone https://github.com/NotNoss/infogen
$ cd infogen
$ go build
$ sudo mv infogen /usr/local/bin/
```

## Usage

### Email
```yaml
$ infogen
or
$ infogen -m
or
$ infogen --mail
```

### Name
```yaml
$ infogen -n
or
$ infogen --name
```

### Address
```yaml
$ infogen -a
or
$ infogen --address
```

### Config
Generate a config file with one of the below commands
```yaml
$ infogen -c
or
$ infogen --config
```
