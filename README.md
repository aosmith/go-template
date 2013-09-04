go-template
===========

Project Config
--------------

You must do one of the following:

Add an alias to include the current directory in gopath:
```
alias go='GOPATH=$GOPATH:`pwd` go'
```

Alternatively you may add the project directory to your gopath.  This may cause namespace conflicts down the line.

Running A Project
-----------------

```
go run app.go
```
