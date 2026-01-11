package main

import "testing"

func Test_hello(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		wantErr bool
	}{
		{
			name:    "It should print hello without error",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := hello()
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("hello() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("hello() succeeded unexpectedly")
			}
		})
	}
}
