# Advent of Code

This repo contains all the advent of code solutions in various languages. 

# Initializing solution code for different languages
Create directory for year and day.
```shell
mkdir 2015
cd 2015
mkdir day01
cd day01
```

## Go


```shell
mkdir go
cd go
go mod init 2015/day01
```
Create main.go
```go
package main

func main() {
    //...
}
```
Create main_test.go
```go
package main

import (
    "testing"
)

func TestStub(t *testing.T) {
    sum := 1 + 1 // change this to simulate failing test
    if sum != 2 {
        t.Fatalf(`Got %d, want %d`, sum, 2)
    }
}
```

## Deno

```shell
mkdir deno
cd deno
deno init
```

## Rescript

```shell
npm create rescript-app@latest
<name project rescript>
cd rescript
rm -rf .git # remove git folder
```

## Elm

```shell
mkdir elm
cd elm
elm init
```
Copy over elm gitignore from another elm folder.