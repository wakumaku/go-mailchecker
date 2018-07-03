IMPORTANT: Due I'm not able to include a golang platform into the original mailchecker (https://github.com/FGRibreau/mailchecker) because I don't know how to ride off function/vars collisions, I tried to find my way doing this thing.

## Mailchecker
[![Build Status](https://travis-ci.org/wakumaku/mailchecker.svg?branch=master)](https://travis-ci.org/wakumaku/mailchecker)

Disposable email address (DEA) detection

```
go get github.com/wakumaku/go-mailchecker
```

```
mailchecker.IsValid("an_email_here@emailprovider.com"
```

Makefile:
* `make gen` Generate/Update blacklist
* `make test` Runs tests
