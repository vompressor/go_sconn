package asym_test

import (
	"crypto/elliptic"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/vompressor/go_sconn/asym/asym_ecdsa"
	"github.com/vompressor/go_sconn/asym/asym_ed25519"
	"github.com/vompressor/go_sconn/asym/asym_rsa"
)

var test_dir = "./test_asym"

func TestGenRSAPairPEM(t *testing.T) {
	os.MkdirAll(filepath.Join(test_dir, "rsa"), os.ModePerm)
	a, err := asym_rsa.New(4096)
	if err != nil {
		t.Fatal(err.Error())
	}

	p, _ := a.PubPEM()
	k, err := a.PrivPEM()
	if err != nil {
		t.Fatal(err.Error())
	}

	ioutil.WriteFile(filepath.Join(test_dir, "rsa", "pub.pem"), p, 0600)
	ioutil.WriteFile(filepath.Join(test_dir, "rsa", "priv.pem"), k, 0600)
}

func TestLoadRSAPairPEM(t *testing.T) {
	pubPath := filepath.Join(test_dir, "rsa", "pub.pem")
	privPath := filepath.Join(test_dir, "rsa", "priv.pem")

	pubPEM, _ := ioutil.ReadFile(pubPath)
	privPEM, _ := ioutil.ReadFile(privPath)

	priv, _ := asym_rsa.LoadPEM(privPEM)
	pub, _ := asym_rsa.LoadPubPEM(pubPEM)

	cipt, err := pub.Encrypt([]byte("test"))
	if err != nil {
		t.Fatal(err.Error())
	}

	plaint, err := priv.Decrypt(cipt)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("%s", plaint)
}

func TestGenECDSAPairPEM(t *testing.T) {
	os.MkdirAll(filepath.Join(test_dir, "ecdsa"), os.ModePerm)
	pubPath := filepath.Join(test_dir, "ecdsa", "pub.pem")
	privPath := filepath.Join(test_dir, "ecdsa", "priv.pem")

	a, err := asym_ecdsa.New(elliptic.P256())
	if err != nil {
		t.Fatal(err.Error())
	}

	pubPEM, _ := a.PubPEM()
	privPEM, _ := a.PrivPEM()

	ioutil.WriteFile(pubPath, pubPEM, 0600)
	ioutil.WriteFile(privPath, privPEM, 0600)

}

func TestLoadECDSAPairPEM(t *testing.T) {
	pubPath := filepath.Join(test_dir, "ecdsa", "pub.pem")
	privPath := filepath.Join(test_dir, "ecdsa", "priv.pem")

	pubPEM, _ := ioutil.ReadFile(pubPath)
	privPEM, _ := ioutil.ReadFile(privPath)

	priv, _ := asym_ecdsa.LoadPEM(privPEM)
	pub, _ := asym_ecdsa.LoadPubPEM(pubPEM)

	sig, err := priv.Sign([]byte("hello"))
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(pub.Verify([]byte("hello"), sig))
}

func TestGenED25519PairPEM(t *testing.T) {
	os.MkdirAll(filepath.Join(test_dir, "ed25519"), os.ModePerm)
	pubPath := filepath.Join(test_dir, "ed25519", "pub.pem")
	privPath := filepath.Join(test_dir, "ed25519", "priv.pem")

	a, err := asym_ed25519.New()

	if err != nil {
		t.Fatal(err.Error())
	}

	pubPEM, _ := a.PubPEM()
	privPEM, _ := a.PrivPEM()

	ioutil.WriteFile(pubPath, pubPEM, 0600)
	ioutil.WriteFile(privPath, privPEM, 0600)
}

func TestLoadED25519PairPEM(t *testing.T) {
	pubPath := filepath.Join(test_dir, "ed25519", "pub.pem")
	privPath := filepath.Join(test_dir, "ed25519", "priv.pem")

	pubPEM, _ := ioutil.ReadFile(pubPath)
	privPEM, _ := ioutil.ReadFile(privPath)

	priv, _ := asym_ed25519.LoadPEM(privPEM)
	pub, _ := asym_ed25519.LoadPubPEM(pubPEM)

	sig, err := priv.Sign([]byte("hello"))
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(pub.Verify([]byte("hello"), sig))
}