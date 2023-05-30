package gpm

import (
	"strconv"
	"testing"
)

func TestMap_GetPermission(t *testing.T) {
	cases := [...]struct {
		key                Key
		expectedPermission Permission
	}{
		// Test case: success
		{
			key:                0,
			expectedPermission: 0b100,
		},
		// Test case: permission does not exist
		{
			key:                1_000,
			expectedPermission: 0,
		},
	}

	pm := &Map{
		0: 0b100,
		1: 0b001,
	}

	for i, v := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			permission := pm.GetPermission(v.key)
			if permission != v.expectedPermission {
				t.Fatalf("expected %b not %b", v.expectedPermission, permission)
			}
		})
	}
}

func TestMap_SetPermission(t *testing.T) {
	cases := [...]struct {
		pm         *Map
		permission Permission
	}{
		// Test case: nil map
		{
			pm: func() *Map {
				mp := Map(nil)
				return &mp
			}(),
			permission: 0b101,
		},
		// Test case: empty map
		{
			pm:         &Map{},
			permission: 0b111,
		},
	}

	for i, v := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			v.pm.SetPermission(0, v.permission)

			if !v.pm.GetPermission(0).Contains(v.permission) {
				t.Fatalf("gpm '%b' was not set correctly", v.permission)
			}
		})
	}
}

func TestPermission_Contains(t *testing.T) {
	cases := [...]struct {
		permission     Permission
		expectedResult bool
	}{
		// Test case: FALSE
		{
			permission:     0b010,
			expectedResult: false,
		},
		// Test case: TRUE
		{
			permission:     0b100,
			expectedResult: true,
		},
	}

	// Object to test
	permission := Permission(0b101)

	for i, v := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := permission.Contains(v.permission)
			if result != v.expectedResult {
				t.Fatalf("expected %v not %v", v.expectedResult, result)
			}

			t.Logf(`|%-10b|%-7b|%v`, permission, v.permission, result)
		})
	}
}
