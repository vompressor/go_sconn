# secure socket library

## ecdh
Elliptic-curve Diffie–Hellman, ECDH implementation   
function   
 - generate ecdsa pub
 - generate ecdsa pub to .pem
 - load pem ro ecdsa pub
 - generate shared key with local private key and remote public key
 
## key_exchanger
key exchange use ecdh   
function
 - key exchange
 - master side
 - slave side

## protocol
lib for making byte based protocol easy

## sconn
Implementation of encrypted communication with a symmetric key   
functions
 - net.Conn upgrade to secured conn (sconn.SConn)
 - encrypt write
 - decrypt read

## snet
TODO::   

A library that encrypts, self-signs, and verifies data   
Integrity and confidentiality guaranteed