# Mail Alias
mail alias is a simple cli written in go for generating a random email alias. I am creating this in an attempt at inproving my opsec as I am bad at coming up with things on my own.

Mail Alias queries a Dictionary API to create the alias.

Currently the project is to only auto generate the aliases. Maybe it will be something else one day but that is unlikely.

This currently only works on Linux but if someone else wants to submit a PR for another OS, I can test and possibly merge it. I only use linux and I am creating this for my own specific use case.

## Usage
```yaml
$ git clone https://github.com/NotNoss/mail-alias
$ cd mail-alias
$ go build
$ sudo mv mail-alias /usr/local/bin/
$ mail-alias
```
