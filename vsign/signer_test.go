package vsign_test

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"fmt"
	"os"
	"testing"

	"github.com/vompressor/go_sconn/vsign"
)

func TestSign(t *testing.T) {
	k := sha256.Sum256([]byte("password"))

	bo, _ := aes.NewCipher(k[:])

	aead, err := cipher.NewGCM(bo)

	if err != nil {
		t.Fatal(err.Error())
	}

	vs := vsign.NewVSigner(aead, sha256.New())

	f, _ := os.Open("siner.go")

	x, err := vs.Seal(f)

	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Printf("%x\n", x)

	f.Seek(0, 0)
	b, err := vs.Open(f, x)
	if err != nil {
		t.Fatal(err.Error())
	}

	println(b)

}
