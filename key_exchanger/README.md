# key_exchanger

## Handshake
```
 | Server |                  | Client |
  0. Generate key pair        0. Generate key pair
 
  1. Listen
                              1. Dial
  2. Accept Socket
                              2. Send client pub key    ( Client Hello )
  3. Recv client pub key <----+
 
  4. Send server pub key                                ( Server Hello )
  +-------------------------> 3. Recv server pubkey
  5. combine client pub, server priv
                              4. combine server pub, client priv
  6. Get shared key                                     ( hash key )
                              5. Get shared key         ( hash key )
```
## Protocol
```
 +-- 2bytes --+-- 2bytes --+-- 4bytes --+-- 4bytes --+--- 32 bytes ----+
 |    Type    |   Method   | Seq (Unuse)|  Body Len  | body (e.g. key) |
 +------------+------------+------------+------------+-----------------+
```
### Client Hello
```
protocol
 - Type: 0x08
 - Method: 0x01
 - Seq: Unuse
 - Body Len: 32
 - Body: [Client pub key]
```
### Server Hello
```
protocol
 - Type: 0x08
 - Method: 0x02
 - Seq: Unuse
 - Body Len: 32
 - Body: [Server pub key]
```

## Shared Key
```
func (kx *KeyExchanger) GenerateSharedKey(remotePub *DHPubKey, hasher hash.Hash) []byte {

    // Get a common value a, b
	a, b := remotePub.Pub.Curve.ScalarMult(remotePub.Pub.X, remotePub.Pub.Y, kx.priv.D.Bytes())

    // Hash a and b
	hasher.Write(a.Bytes())
	hasher.Write(b.Bytes())
	return hasher.Sum(nil)
}
```