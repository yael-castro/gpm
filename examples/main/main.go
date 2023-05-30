package main

import (
	"fmt"
	"github.com/yael-castro/gpm"
	"log"
)

// Step 1: Create the permission keys using iota
const (
	WriteKey gpm.Key = iota
	ReadKey
)

// Step 2: create the permissions associated with WriteKey using the bitwise operator << and iota
const (
	WriteName gpm.Permission = 1 << iota
	WriteLastName
)

// Step 3: create the permissions associated with ReadKey using the bitwise operator << and iota
const (
	ReadName gpm.Permission = 1 << iota
	ReadLastName
)

func main() {
	// Step 4: get the system user permissions (in this case it logic is simulated)
	pm := gpm.Map{
		WriteKey: WriteName,
		ReadKey:  ReadName | ReadLastName, // Always append the permissions using the bitwise operator |
	}

	// Step 5: validate with a O(1) complexity if the user has the right permissions to perform a protected action
	var err error

	switch {
	case !pm.GetPermission(WriteKey).Contains(WriteName | WriteLastName):
		err = fmt.Errorf("you do not have the correct permissions for permission group '%d'", WriteKey)

	case !pm.GetPermission(ReadKey).Contains(ReadName):
		err = fmt.Errorf("you do not have the correct permissions for permission group '%d'", ReadKey)
	}

	if err != nil {
		log.Fatal(err)
	}
}
