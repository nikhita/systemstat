# systemstat

[Documentation online](http://godoc.org/bitbucket.org/bertimus9/systemstat)

**systemstat** is a package written in Go generated automatically by `gobi`. Happy hacking!

## Install (with GOPATH set on your machine)
----------

* Step 1: Get the `systemstat` package

```
go get bitbucket.org/bertimus9/systemstat
```

* Step 2 (Optional): Run tests

```
$ go test -v ./...
```

##Usage
----------
```
package main

import (
  "fmt"
  "os"
  "bitbucket.org/bertimus9/systemstat"
)

func main() {
  systemstatExample, err := systemstat.New(1, "gobi")
  if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

  systemstatExample.SetId(systemstatExample.Id() + 1)
  systemstatExample.SetName(systemstatExample.Name() + " is great")

  fmt.Println(systemstatExample.Id(), systemstatExample.Name())
  // Output: 2 gobi is great
}
```

##License
----------
systemstat is MIT licensed.
