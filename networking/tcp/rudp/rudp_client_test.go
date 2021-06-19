package rudp

import (
	"testing"
)

func TestClientE2EDirectCommunication(t *testing.T) {
	receiver := RUDPClient{SockPort: 8000}
	sender := RUDPClient{SockPort: 0, SendToPort: 8000}
	defer receiver.Close()
	defer sender.Close()

	if receiver.fd != 0 {
		t.Errorf("Expected fd of client to be empty after initialization")
	}
	if sender.fd != 0 {
		t.Errorf("Expected fd of client to be empty after initialization")
	}
	err := receiver.OpenSocket()
	if err != nil {
		t.Errorf("Expected OpenSocket to not return error but got %s", err)
	}
	if receiver.fd < 1 {
		t.Errorf("Expected fd of receiver to be set to positive integer after opening socket")
	}
	err = sender.OpenSocket()
	if err != nil {
		t.Errorf("Expected OpenSocket to not return error but got %s", err)
	}
	if sender.fd < 1 {
		t.Errorf("Expected fd of sender to be set to positive integer after opening socket")
	}
	err = receiver.BindPort()
	if err != nil {
		t.Errorf("Expected BindPort to not return error but got %s", err)
	}
	err = sender.BindPort()
	if err != nil {
		t.Errorf("Expected BindPort to not return error but got %s", err)
	}

	// Test sending bytes back and forth
	b := "hello my name is reliable udp"
	err = sender.Send([]byte(b))
	if err != nil {
		t.Errorf("Expected Send to not return error but got %s", err)
	}
	rsp, n, err := receiver.Receive()
	if err != nil {
		t.Errorf("Expected Receive to not return error but got %s", err)
	}
	if n != len(b) {
		t.Errorf("Expected number of bytes to equal sent data but got %d", n)
	}
	if string(rsp[:n]) != b {
		t.Errorf("Expected data received to equal data sent but got %s", rsp)
	}
}

func TestRUDPChecksum(t *testing.T) {
	lrc := CalcChecksum([]byte("hey"))
	lrc2 := CalcChecksum([]byte("heythere"))
	lrc3 := CalcChecksum([]byte("heythere"))
	if lrc == lrc2 {
		t.Errorf("Expected checksums to be different but they were the same: %d <> %d", lrc, lrc2)
	}
	if lrc3 != lrc2 {
		t.Errorf("Expected checksums to be same but they were the same: %d <> %d", lrc3, lrc2)
	}
}
