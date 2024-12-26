# Testing + ORM
This mini-project is designed to demonstrate working with <i>ORM</i> and <i>Testing</i>

<b>Lab 6</b> - Testing services with `UNIT tests` and using `mock repository`

<b>Lab 7</b> - Test app for work with `ORM` (Object-Relational Mapping) in Go (with `GORM`)

## Run app
* Instruction for start test app on linux (Ubuntu)

1. Copy `.env.example` to `.env`:
    ```bash
    cp .env.example .env
    ```

2. Put vars into `.env`:
    ```bash
    nano .env
    ```

3. If you have not DB, run Postgres in Docker environment. Use one of this commands (`make` is required in first command):
    ```bash
    make run
    ```

    ```bash
    docker compose up -d
    ```

4. Comment out the comment lines with `-2-` and `-3-` for consistent testing of the implemented operations.

5. Run the test application from root folder of project:
    ```bash
    go run ./cmd/main.go
    ```


## Description
Output of run this program is in the file `cmd/output.txt`

Client code are located in `cmd` directory.

The implementation of this application is located in the `internal` folder and have this structure:
```bash
.
├── cmd
│   ├── main.go
│   └── output.txt
├── docker-compose.yaml
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── database
│   │   └── database.go
│   ├── models
│   │   └── productModels.go
│   ├── repositories
│   │   ├── categoryRepository.go
│   │   └── productRepository.go
│   ├── services
│   │   ├── categoryService.go
│   │   └── productService.go
│   ├── tests
│   │   └── tests.go
│   └── utils
│       └── serializer.go
├── Makefile
└── README.md

10 directories, 16 files
```

## Results

Tests for services shows this result:
```
ok  	orm-tests/internal/services	0.010s	coverage: 85.7% of statements
```

<hr>
It was created on 25.12.24