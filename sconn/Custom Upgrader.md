# Custom Upgrader
package block, stream, aead have an `Upgrade` function   
This function creates a SConn by inputted the socket, cipher, and other necessary values.   

### block.Upgrade
```
// conn - Socket to be used for encrypted communication
// cip - block cipher
// eg - block encription mode getter e.g. cipher.NewCBCEncrypter
// dg - block decryption mode getter e.g. cipher.NewCBCDecrypter
func Upgrade(conn net.Conn, cip cipher.Block, eg, dg BlockModeGetter) sconn.SConn
```
### stream.Upgrade
```
// conn - Socket to be used for encrypted communication
// cip - stream cippher
Upgrade(c net.Conn, cip cipher.Stream) sconn.SConn
```
### aead.Upgrade
```
// conn - Socket to be used for encrypted communication
// cip - AEAD cipher
// init - Some value to be used to generate the nonce
//        Each host must set the same init
//        e.g. hashed key
func Upgrade(conn net.Conn, cip cipher.AEAD, init []byte) *AEADSConn
```

Depending on the type of encryption you need to use that feature.   
## sconn.ConnUpgrader
The function that will actually be used is the type 'sconn.Connupgrader' function.   
This function creates an encrypted socket by inputted a socket and key.   
Libs such as `key_exchanger` use this function.   

```
// conn - Socket to be used for encrypted communication
// key - Key to create a cipher
type ConnUpgrader func(conn net.Conn, key []byte) (SConn, error)
```
Create a cipher from the input key and call the 'Upgrade' function.   
   
e.g. aes_cbc mode   
```
func Upgrade(conn net.Conn, key []byte) (sconn.SConn, error) {
	cip, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return block.Upgrade(conn, cip, cipher.NewCBCEncrypter, cipher.NewCBCDecrypter), nil
}
```