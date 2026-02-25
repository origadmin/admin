// Package providers implements the functions, types, and contracts for the module.
package providers

import (
	"testing"
)

func TestProvideHasher(t *testing.T) {
	tests := []struct {
		name      string
		oldString string
		password  string
		wantErr   bool
	}{
		{
			name:      "Testing ProvideHasher v1",
			oldString: "$bcrypt$v1$c:10$243261243130244e44744e3144515a53446a63372e32626a752e354e75576c6a594a6c495544326f4346466e373763576a4d49325157446b372f4879$4a36457446474d376b35674d5a4c7376",
			password:  "admin123",
			wantErr:   false,
		},
		{
			name:      "Testing ProvideHasher v2",
			oldString: "$bcrypt$v1$c:100244e44744e3144515a53446a63372e32626a752e354e75576c6a594a6c495544326f4346466e373763576a4d49325157446b372f4879a36457446474d376b35674d5a4c7376",
			password:  "testpassword",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasher, err := ProvideHasher()
			if err != nil {
				t.Fatalf("ProvideHasher() error = %v", err)
				return
			}
			err1 := hasher.Verify(tt.oldString, tt.password)
			if (err1 != nil) != tt.wantErr {
				t.Errorf("Hasher.Verify() error = %v, wantErr %v", err1, tt.wantErr)
			}
		})
	}
}
