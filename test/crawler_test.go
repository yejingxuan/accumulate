package test

import (
	"fmt"
	"testing"
	"time"
)

func TestDate(t *testing.T) {

	unix := time.Now().UnixNano() / 1e6
	fmt.Println(unix)
}
