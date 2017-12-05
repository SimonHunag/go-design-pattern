package singleton

import "testing"

func TestGetInstance(t *testing.T) {

	manger := GetInstance()
	manger.name();
}
