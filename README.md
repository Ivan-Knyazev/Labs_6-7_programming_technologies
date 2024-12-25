# Testing + ORM

## --- Run app
Run the desired application from root folder of project:
```bash
go run <type-of-patterns>/cmd/<pattren-name>/main.go
```

## Description
Output of run this programs are in the same place in the file `output.txt`

Client code are located in `cmd` directory.

The implementations of the patterns are located in the `internal` folder.

## Example
For example, creational patterns are located in `creational` folder.

Client code for `singleton` is located in `creational/cmd/singleton/main.go`.

Implementation code of `singleton` is located in `creational/internal/singleton/singleton.go`.

Output is located in `creational/cmd/singleton/output.txt`

For run this app from root folder of project:
```bash
go run creational/cmd/singleton/main.go
```

<hr>
It was created on 25.12.24.