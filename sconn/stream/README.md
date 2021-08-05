# stream
SConn used stream cipher.   
Encrypt/decrypt communication with stream ciphers.  

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