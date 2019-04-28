/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/clustersmgmt/v1

// AdminCredentials represents the values of the 'admin_credentials' type.
//
// Temporary administrator credentials generated during the installation of the
// cluster.
type AdminCredentials struct {
	user     *string
	password *string
}

// User returns the value of the 'user' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Cluster administrator user name.
func (o *AdminCredentials) User() string {
	if o != nil && o.user != nil {
		return *o.user
	}
	return ""
}

// GetUser returns the value of the 'user' attribute and
// a flag indicating if the attribute has a value.
//
// Cluster administrator user name.
func (o *AdminCredentials) GetUser() (value string, ok bool) {
	ok = o != nil && o.user != nil
	if ok {
		value = *o.user
	}
	return
}

// Password returns the value of the 'password' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Cluster administrator password.
func (o *AdminCredentials) Password() string {
	if o != nil && o.password != nil {
		return *o.password
	}
	return ""
}

// GetPassword returns the value of the 'password' attribute and
// a flag indicating if the attribute has a value.
//
// Cluster administrator password.
func (o *AdminCredentials) GetPassword() (value string, ok bool) {
	ok = o != nil && o.password != nil
	if ok {
		value = *o.password
	}
	return
}

// AdminCredentialsList is a list of values of the 'admin_credentials' type.
type AdminCredentialsList struct {
	items []*AdminCredentials
}

// Len returns the length of the list.
func (l *AdminCredentialsList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *AdminCredentialsList) Slice() []*AdminCredentials {
	var slice []*AdminCredentials
	if l == nil {
		slice = make([]*AdminCredentials, 0)
	} else {
		slice = make([]*AdminCredentials, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *AdminCredentialsList) Each(f func(item *AdminCredentials) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *AdminCredentialsList) Range(f func(index int, item *AdminCredentials) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}