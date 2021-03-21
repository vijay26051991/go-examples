# go-examples

Go Language:

    -  Appeared in 2009.
    -  Go 1.0 was released in 2012.
    -  Conceived at Google.
    -  Go is simple, minimal and programatic..
        - simple - Designed for simplicity.
        - Features of Go is minimal..
        - 25 keywords
        - Progmatic - It focus on practical side.
        - Design comes with more practical side.
        - Focuses on solving on the problems.
    -  Statically typed Language
        - Explicitly typed.
        - Type inference.
        - Performance is good.
        - C, C++, Java and GO.
    -  Dynamically typed Language
        - Type is set at Runtime..
        - Performance is not good.
        - Javascript, Groovy and Ruby.

    - Compiles at native machine code.
    - Does not need virtual machine and JIT complier.
    - Garbage Collected Language. --> Remove unused memory.
    - Concurrency is built-in feature.

Package Identifiers:
----------------------
    - Main : An executable code.
    - any other package name : shared library.

Loops:
-------
https://yourbasic.org/golang/for-loop/

Collections:
----------------
    - Slice
    - Map
    - Arrays

Slices:
-------
https://blog.golang.org/slices

**package main** --> executable source code.

**package package_name** --> library

**fmt** - Standard library package.

**func main()** --> Entrypoint to the go


Commands
--------------
    - go build : Build an executable
    - go clean : cleans executable.
    - go run : compiles and runs the source code.
    - go mod init : records Dependency management
    - GOMODCACHE : The place where all third party code to be placed in.
    - go get - Gets the package with latest version
    - go mod tidy - It looks up the import statement and get it
                    sync up and remove unneeded packages.
    - go env GOMODCACHE -- Returns the path.
    - go mod vendor -> It creates directory vendor which contains a text file having all the modules to be supported for builds and tests.
    - go mod init <module_name>  -> Will initialize a module and keeps all the dependencies Basically, it is a dependency management.
    - go mod verify -> It checks the dependencies of the main module available in the module cache and have not been modified since they were downloaded.
    - gofmt -w main.go --> formats the source code.
    - go fmt /Users/vsubra278/personal_repo/Workspace/go/training
    - go test ./... -> traverse all directories and test all go classes.
    - go test ./... -> traverse all directories and build all go classes.

Errors
-------------
    - Errors are values because it can be stored as variable.

User Defined types:
--------------------
    - Struct - Concrete data type.
        * struct does not include behaviors.
    - Interface 
    - No inheritance(type embedding to implement composition).
    - No classes.

Go Routines:
------------------
    - goroutine function is a function which runs the function concurrently.
    - Channels are great signalling function..

    Send : <- Channel operator
    Recieve : ch <- ""  / send operation
    val := <- ch  / receive operation

Channel:
-------------
    Channel used to send and receive message witnin go routines and also it is used for signalling.

    Types   
    ---------
        Channel has two types
        Unbuffered channel - doesn't have capacity.
        Buffer channel - It has capacity.

```
Sample code for creating channel
    chInterrupt := make(chan os.Signal,1)
    func gotInterrupted() bool {
    select {
        case <-chInterrupt:
            return true
        default:
            return false
    }
```

HTTP programming
---------------
    - net/http :  Standard library
    - Servermux(struct) : router
    - Handler(interface) : (responsewriter, http.request)
    - handle or handlefunc to create handler function
    
HTTP tests:
-------------
    - "net/http/httptest" : Test library to execute the tests.
    - ResponseRecorder(One way to implement http tests). ResponseRecorder is an implementation of http.ResponseWriter that records its mutations for later inspection in tests.
    - A Server is an HTTP server listening on a system-chosen port on the local loopback interface, for use in end-to-end HTTP tests. Server is a test server just for running unit tests
      ginkgo  
    - Server(One way to implement http tests).

GORILLAMUX
-------------
    - Third package library used for implementing http programming.


Reference links
--------------
    https://shijuvar.medium.com/
    https://github.com/go-delve/delve
    https://www.youtube.com/watch?v=oV9rvDllKEg(Concurrency is not parallism)
    https://github.com/golang/mock






