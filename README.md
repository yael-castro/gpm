# GPM: Go Permission Masks 

![banner](docs/banner.svg)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg)](https://pkg.go.dev/github.com/yael-castro/gpm)


#### Overview
GPM is a Go library created to manage many permissions on different systems with constant complexity
of the `O(1)` algorithm and optimizing memory usage.

#### Theory

The key to `O(1)` complexity when managing permissions are bitmasks and hash maps.

###### Bitmasks

Basically, a bitmask is an integer in base 2.<br>
So a bitmask represents a permission for each bit as show below.

```
UNSIGNED INTEGER
________________
BASE 10 | BASE 2
________|_______
9       | 1001

PERMISSIONS           
_______________________________
CREATE | READ | UPDATE | DELETE
_______|______|________|_______
1      | 0    | 0      | 1
```
It is important to note that the number of permissions that can be represented is limited by the number of bits.

###### Hash maps
As seen in the previous point, there is a limit to the permissions that can be
represented by a bitmask.

The idea to break this limit is to group the bitmasks by key, as shown below.

```
KEY | BITMASK
____|________
0   | 1 0 0 1
1   | 1 1 0 0
```
And the ideal data structure for this purpose is the hash map.

#### Bitwise operators
The bitwise operators can help to create permission constants

`<< (LEFT SHIFT)`

This operator can be used to create permissions.<br>
For example.

```
UNSIGNED INTEGER | BITMASK | OPERATION | RESULTING BITMASK
_________________|_________|___________|__________________
1                |  0 1    | 0 1 << 1  | 1 0
```

`| (OR)`

This operator can be used to merge two or more permissions.<br>
For example.

```
UNSIGNED INTEGER | BITMASK | OPERATION  | RESULTING BITMASK 
_________________|_________|____________|__________________
1                |  0 1    | 0 1 OR 1 0 | 1 1
```

#### Installation

```shell
go get github.com/yael-castro/gpm@latest
```

#### How to use
###### Create the keys
Create the keys to group the permissions
```go
const (
    WriteKey gpm.Key = iota
    ReadKey
)
```

###### Create the permissions
Create the constants for the permissions for each key

```go
// Permission constants for the WriteKey
const (
    WriteName gpm.Permission = 1 << iota
    WriteLastName
)

// Permission constants for the ReadKey
const (
    ReadName gpm.Permission = 1 << iota
    ReadLastName
)
```

###### Create permission map
```go
pm := gpm.Map{
    WriteKey: WriteName,
    ReadKey:  ReadName | ReadLastName,
}
```

###### Validate the permissions
This is one of the many ways to validate permissions.
```go
switch {
    case !pm.GetPermission(WriteKey).Contains(WriteName | WriteLastName):
        err = fmt.Errorf("you do not have the correct permissions for permission group '%d'", WriteKey)

    case !pm.GetPermission(ReadKey).Contains(ReadName):
        err = fmt.Errorf("you do not have the correct permissions for permission group '%d'", ReadKey)
}
```