package lean

import (
	"testing"
)

func TestSignature(t *testing.T) {
	println("begin get signature")
	client := NewClient("FFnN2hso42Wego3pWq4X5qlu", "UtOCzqb67d3sN12Kts4URwy8", "DyJegPlemooo4X1tg94gQkw1")
	agent := client.GetObjectById("test", "57e4fd355bbb50005d499f3e")
	agent.ts = 1453014943466
	keySign := agent.getSignature()
	if keySign != "d5bcbb897e19b2f6633c716dfdfaf9be,1453014943466" {
		t.Error("signature error:" + keySign)
	}

	agent = agent.UseMasterKey()
	masterSign := agent.getSignature()
	if masterSign != "e074720658078c898aa0d4b1b82bdf4b,1453014943466,master" {
		t.Error("master signature error " + masterSign)
	}
	println("end get signature")

}
