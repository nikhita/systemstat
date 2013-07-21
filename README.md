# systemstat

[Documentation online](http://godoc.org/bitbucket.org/bertimus9/systemstat)

**systemstat** is a package written in Go generated automatically by `gobi`.

It is a package that allows you to add system statistics to your go program;
currently polls the linux kernel for CPU usage, free/used memory and swap
sizes, and uptime for your go process as well as the linux system you're
running it on, as well as the system load. Can be used to make a crippled
version of top that monitors the current go process and ignores other processes
and the number of users with ttys.

## Install (with GOPATH set on your machine)
----------

* Step 1: Get the `systemstat` package

```
go get bitbucket.org/bertimus9/systemstat
```

* Step 2 (Optional): Run tests

```
$ go test -v bitbucket.org/bertimus9/systemstat
```

##Usage
----------
```
// read in file ../foo.go
```

##License
----------
systemstat is MIT licensed.
