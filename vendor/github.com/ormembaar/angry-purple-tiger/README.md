# angry-purple-tiger

Animal-based hash digests for humans... in Go.

## Overview

Angry Purple Tiger generates animal-based hash digests meant to be memorable
and human-readable. Angry Purple Tiger is apt for anthropomorphizing project
names, crypto addresses, UUIDs, or any complex string of characters that need
to be displayed in a user interface.

## Install

```
go get "github.com/ormembaar/angry-purple-tiger"
```

## Example

```go
import (
	"fmt"

	apt "github.com/ormembaar/angry-purple-tiger"
)

data := []byte("112CuoXo7WCcp6GGwDNBo6H5nKXGH45UNJ39iEefdv2mwmnwdFt8")
animalName := apt.Sum(data)
fmt.Println(animalName) // "feisty-glass-dalmatian"
```
