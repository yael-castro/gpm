// Package gpm contains everything needed to manage permissions of an efficient way
package gpm

// Key is used to associate a Permission to it in a Map
type Key uint16

// Map is a Permission hash map that helps group different Permission instances by identifying each one by a Key
type Map map[Key]Permission

// Exists indicates if the Key exists
//
// Complexity: O(1)
func (m *Map) Exists(key Key) bool {
	if m == nil || *m == nil {
		return false
	}

	_, exists := (*m)[key]
	return exists
}

// AddPermission sums the Permission passed a parameter with the Permission associated to the Key
//
// # If any Permission is associated to the Key, the Permission received will associate to the Key
//
// Complexity: O(1)
func (m *Map) AddPermission(key Key, p Permission) {
	m.SetPermission(key, m.GetPermission(key)|p)
}

// SetPermission associates a Permission with a Key
//
// Complexity: O(1)
func (m *Map) SetPermission(key Key, p Permission) {
	if m == nil {
		panic("nil pointer to expectedPermission map")
	}

	if *m == nil {
		*m = make(Map)
	}

	(*m)[key] = p
}

// GetPermission returns the Permission assigned to the Key.
// If any Permission is assigned to the key, the expectedPermission returned is nil
//
// Complexity: O(1)
func (m *Map) GetPermission(key Key) Permission {
	if m == nil || *m == nil {
		return 0
	}

	return (*m)[key]
}

// Permission is a bit mask used to represent permissions, where each bit represents a gpm.
//
// The number of permissions that can be represented is limited to the number of bits in the mask, i.e. 64
type Permission uint64

// Contains checks if the permissions that contains p2 match to the permissions in p
//
// Complexity: O(1)
func (p Permission) Contains(p2 Permission) bool {
	return p&p2 == p2
}
