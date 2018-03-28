// Copyright (c) 2014 The grhsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package grhjson_test

import (
	"testing"

	"github.com/grhsuite/grhd/grhjson"
)

// TestErrorCodeStringer tests the stringized output for the ErrorCode type.
func TestErrorCodeStringer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   grhjson.ErrorCode
		want string
	}{
		{grhjson.ErrDuplicateMethod, "ErrDuplicateMethod"},
		{grhjson.ErrInvalidUsageFlags, "ErrInvalidUsageFlags"},
		{grhjson.ErrInvalidType, "ErrInvalidType"},
		{grhjson.ErrEmbeddedType, "ErrEmbeddedType"},
		{grhjson.ErrUnexportedField, "ErrUnexportedField"},
		{grhjson.ErrUnsupportedFieldType, "ErrUnsupportedFieldType"},
		{grhjson.ErrNonOptionalField, "ErrNonOptionalField"},
		{grhjson.ErrNonOptionalDefault, "ErrNonOptionalDefault"},
		{grhjson.ErrMismatchedDefault, "ErrMismatchedDefault"},
		{grhjson.ErrUnregisteredMethod, "ErrUnregisteredMethod"},
		{grhjson.ErrNumParams, "ErrNumParams"},
		{grhjson.ErrMissingDescription, "ErrMissingDescription"},
		{0xffff, "Unknown ErrorCode (65535)"},
	}

	// Detect additional error codes that don't have the stringer added.
	if len(tests)-1 != int(grhjson.TstNumErrorCodes) {
		t.Errorf("It appears an error code was added without adding an " +
			"associated stringer test")
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.String()
		if result != test.want {
			t.Errorf("String #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}

// TestError tests the error output for the Error type.
func TestError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		in   grhjson.Error
		want string
	}{
		{
			grhjson.Error{Description: "some error"},
			"some error",
		},
		{
			grhjson.Error{Description: "human-readable error"},
			"human-readable error",
		},
	}

	t.Logf("Running %d tests", len(tests))
	for i, test := range tests {
		result := test.in.Error()
		if result != test.want {
			t.Errorf("Error #%d\n got: %s want: %s", i, result,
				test.want)
			continue
		}
	}
}
