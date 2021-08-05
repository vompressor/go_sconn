# key_exchanger
Generate shared key from key exchange.
And create `SConn` use shared key.
## net.Conn -> sconn.SConn
```
// Server
c, _ := net.Accept()
ac, err := key_exchanger.ServerSideUpgrade(c, chacha20poly1305_upgrader.Upgrade)

if err != nil {
    println(err.Error())
    os.Exit(1)
}
```

```
// Client
c, _ := net.Dial(...)
ac, err := key_exchanger.Upgrade(c, chacha20poly1305_upgrader.Upgrade)

if err != nil {
    println(err.Error())
    os.Exit(1)
}
```

## key_exchanger Listen/Dial
```
// Server
l, err := net.Listen(...)
if err != nil {
    println(err.Error())
    return
}
defer l.Close()

excl := key_exchanger.NewExcListener(l, aes_ctr_upgrader.Upgrade)

sc, err := excl.Accept()
if err != nil {
    println(err.Error())
    return
}
```
```
// Client
sc, err := key_exchanger.ExcDial(..., aes_ctr_upgrader.Upgrade)
if err != nil {
    t.Fatal(err.Error())
}
```