package main

import (
	"testing"
)

func BenchmarkGetUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetUser()
		//t.Run("fw", func(t *testing.T) {
		//	result, err := GetUser()
		//	if err != nil {
		//		t.Error(err)
		//	} else {
		//		t.Log(result)
		//	}
		//})
	}
}

func TestGetUser(t *testing.T) {
	t.Run("tt.name", func(t *testing.T) {
		got, err := GetUser()
		if err != nil {
			t.Errorf("GetUser() error = %v", err)
			return
		} else {
			t.Log(got)
		}
	})
}
