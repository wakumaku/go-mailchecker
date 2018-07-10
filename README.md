# Mailchecker [![Build Status](https://travis-ci.org/wakumaku/go-mailchecker.svg?branch=master)](https://travis-ci.org/wakumaku/go-mailchecker) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/9b66f7d42dcb413bbf96f8f4d1471020)](https://www.codacy.com/app/wakumaku/go-mailchecker?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=wakumaku/go-mailchecker&amp;utm_campaign=Badge_Grade) [![Code Coverage](https://scrutinizer-ci.com/g/wakumaku/go-mailchecker/badges/coverage.png?b=master)](https://scrutinizer-ci.com/g/wakumaku/go-mailchecker/?branch=master) [![GoDoc](https://godoc.org/github.com/wakumaku/go-mailchecker?status.svg)](https://godoc.org/github.com/wakumaku/go-mailchecker)

### Disposable email address (DEA) detection
*Super inspired by: https://github.com/FGRibreau/mailchecker

```
go get github.com/wakumaku/go-mailchecker
```

```
mailchecker.IsValid("an_email_here@provider.tld")
```

Makefile:
* `make gen` Generate/Update blacklist
* `make test` Runs tests

*Note: Due I'm not able to include a golang platform into the original mailchecker (https://github.com/FGRibreau/mailchecker) because I don't know how to ride off function/vars collisions, I tried to find my way doing this thing.
