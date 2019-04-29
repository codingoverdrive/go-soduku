# go-soduku
Soduku Solver in Go

A package comprising different solvers for Soduku puzzles.

# Solvers

The current implemented solvers are:
* Naked Singles
* Hidden Singles

# Running the Solver

Checkout the project from github

```
git clone https://github.com/codingoverdrive/go-soduku.git
cd cmd/solver
go run main.go
```

To solve a different puzzle, modify the board two dimensional array in cmd/solver/main.go

The 9x9 array represents a soduku board