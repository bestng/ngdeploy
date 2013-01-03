package mongodb

import "testing"

func TestSave(t *testing.T) {
	user := User{"mj", "manjia", "127.0.0.1", "mjtest1"}
	err := user.Save()
	if err != nil {
		t.Error(err)
	}
}
