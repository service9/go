package test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

//go get github.com/satori/go.uuid
//36位 唯一标识
func TestGenerateUUid(t *testing.T) {
	s := uuid.NewV4().String()
	println(s)
}
