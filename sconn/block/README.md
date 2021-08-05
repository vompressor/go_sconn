# block
SConn used block cipher with mode.   
Encrypt/decrypt communication with block ciphers.  

##
Encription Write
```
 +-------+    +-----------------+    +--------+    +-------------+    +--------+
 | plain |--->| plain + padding |--->| encode |--->| mode cipher |--->| socket |
 +-------+    +-----------------+    +--------+    +-------------+    +--------+

 protocol
 +-- 4bytes --+-- 16 bytes -- +--- 16 multiples bytes ---+
 | Msg Length |      iv       |   cipher text  | padding |
 +------------+---------------+----------------+---------+

```
Decription Read
```
 +--------+    +-------------+    +--------+    +-----------------+
 | socket |--->| mode cipher |--->| decode |--->| plain - padding |
 +--------+    +-------------+    +--------+    +-------+---------+
                                                        |
                              +-------+    +--------+   |
                              | plain |<---| buffer |<--+
                              +-------+    +--------+
```