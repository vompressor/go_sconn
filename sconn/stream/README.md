# stream
SConn used stream cipher.   
Encrypt/decrypt communication with stream ciphers.  
 - simple
 - fast
 - No packet length verification
 - No data integrity verification
 - No buffer
 - Vulnerable to tampering

 - chacha20
 - aes-ctr

##
Encription Write
```
 +-------+      +--------+      +--------+
 | plain |----->| cipher |----->| socket |
 +-------+      +--------+      +--------+

 protocol
 +-- ... bytes ---+
 |   cipher text  |
 +----------------+
```
Decription Read
```
 +--------+      +--------+      +-------+
 | socket |----->| cipher |----->| plain |
 +--------+      +--------+      +-------+
 
```