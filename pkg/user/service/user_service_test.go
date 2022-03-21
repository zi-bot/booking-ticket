package service

import (
	"fmt"
	"testing"
)

func TestGenerateHash(t *testing.T) {
	fmt.Println(hashPassword("admin2"))
}
