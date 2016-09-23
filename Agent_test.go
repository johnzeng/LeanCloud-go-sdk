package lean

import (
	"testing"
)

func TestSignature(t *testing.T) {
	client := NewClient("L1rboIrylg7wJklCPV8v6TCO-gzGzoHsz", "", "")
	client.GetObjectById("test", "57e4fd355bbb50005d499f3e")

}
