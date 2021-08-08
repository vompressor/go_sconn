# secure socket library
A library for secretly exchanging messages between two hosts.   
A library for creating lightweight and fast cryptographic sockets.

## ecdh
Elliptic-curve Diffie–Hellman, ECDH implementation   
 - generate ecdsa pub
 - generate ecdsa pub to .pem
 - load pem to ecdsa pub
 - generate shared key with local private key and remote public key
 
## key_exchanger
Key exchange use ecdh   
 - generate pub/priv pair
 - send local pub to remote
 - generate shared key from local priv and remote pub
 - listen, if accepted, create sconn use exchanged shared key
 - dial, if dialed, create sconn use exchanged shared key

## protocol
Lib for making byte based protocol easy   
 - define protocol header
 - write protocol to io.Writer
 - read protocol from io.Reader
 - read protocol body as much as the length set in the header

## sconn
Implementation of encrypted communication with a symmetric key   
functions
 - net.Conn upgrade to secured conn (sconn.SConn)
 - encrypt write
 - decrypt read
There are three kinds of implementations.   
 - Block cipher used
   - aes cbc mode
 - stream cipher used
   - aes ctr mode
   - chacha20
 - AEAD cipher used
   - aes gcm mode
   - chacha20-poly1305

# TODO::   
 1. High-capacity data transmitter with guaranteed integrity and confidentiality