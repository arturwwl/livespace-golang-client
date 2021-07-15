# Livespace golang client

### New client from config
```go
client, err := livespace_client.New("conf/cfg.ini")
if err != nil {
    t.Fatal(err)
}
//(...) some client actions
```

### TODO:
- [ ] custom errors
- [ ] other requests than just create/get client and create note