package sql

import "testing"

// go test v -cgo -run TestInit
func TestInit(t *testing.T) {

	SetEngine()
	y := new(Ytdlp)
	y.ID = 1
	y.Status = "success"
	result := y.UpdateStatusById()
	t.Logf("1.%v\n2.%v\n", result.Statement, result.Error)
}
