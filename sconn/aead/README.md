# aead
SConn used aead cipher.   
Encrypt/decrypt communication with aead ciphers.  
 - Packet length verification
 - Cipher text and additional data integrity verification

 - chacha20-poly1305
 - aes-gcm

##
Encription Write
```
 +-------+    +--------+    +--------+    +--------+
 | plain |--->| encode |--->| cipher |--->| socket |
 +-------+    +--------+    +--------+    +--------+

  protocol
 +-- 4bytes --+-- 4bytes --+---- 4bytes ----+-- ... bytes --+--------- ... bytes --------+
 | Msg Length |    Seq     | Additional Len |  cipher text  | additional data (optional) |
 +------------+------------+----------------+---------------+----------------------------+
```
Decription Read
```
 +--------+    +--------+    +--------+    +--------+    +-------+
 | socket |--->| cipher |--->| decode |--->| buffer |--->| plain |
 +--------+    +--------+    +--------+    +--------+    +-------+
```

## nonce
```
// nonce = sha256 ( init bytes + seq[4bytes] )

func makeNonce(seq uint32, init []byte, nonceLen int) []byte {
	hasher := sha256.New()
	var seqBytes [4]byte
	hasher.Write(init)

	binary.BigEndian.PutUint32(seqBytes[:], seq)

	hasher.Write(seqBytes[:])

	tNonce := hasher.Sum(nil)

	if len(tNonce) < nonceLen {
		padding := make([]byte, nonceLen-len(tNonce))
		tNonce = append(tNonce, padding...)
	}
	return tNonce[:nonceLen]
}
```