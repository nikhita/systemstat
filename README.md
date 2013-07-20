# serverstat

[Documentation online](http://godoc.org/bitbucket.org/bertimus9/serverstat)

**serverstat** is a package written in Go generated automatically by `gobi`. Happy hacking!

## Install (with GOPATH set on your machine)
----------

* Step 1: Get the `serverstat` package

```
go get bitbucket.org/bertimus9/serverstat
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
  "bitbucket.org/bertimus9/serverstat"
)

func main() {
  serverstatExample, err := serverstat.New(1, "gobi")
  if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

  serverstatExample.SetId(serverstatExample.Id() + 1)
  serverstatExample.SetName(serverstatExample.Name() + " is great")

  fmt.Println(serverstatExample.Id(), serverstatExample.Name())
  // Output: 2 gobi is great
}
```

##License
----------
serverstat is MIT licensed.