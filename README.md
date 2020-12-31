# Minichat

For those who are accustomed to the terminal interface, **minichat** has a certain effect. The message is transmitted through the established udp connection, and then the message encryption module will be added. Need to be used with the support of intranet penetration.

## Obtain

After cloning from Github, you can use the go compiler to compile.

```bash
git clone https://github.com/fooage/minichat.git
# Move to the directory and compile it.
go build
```

- Tips: An error may occur if the terminal is not UTF-8 encoded!
- Tips: Please `ping` your target address before use minichat.

Use the <kbd>Ctrl</kbd> + <kbd>C</kbd> quit this, or use the <kbd>Cmd</kbd> + <kbd>C</kbd> on mac.

## Parameter

### Main structure

Before introducing the parameters, it is important to understand the structure of the handler.

```go
type Handler struct {
	// The message receive channel.
	Buf chan []byte
	// Key for AES encryption.
	AesKey []byte
	// Which is client's listen address and connection with the other client.
	LocalAddr  string
	RemoteAddr string
}
```

### Startup parameters

As indicated by the above structure, there are two parameters should be set. The first two parameters are must have.

- `h`

Set the host address and listening port, the default is `127.0.0.1:10000`.

- `t`

Set the target address and listening port, the default is `127.0.0.1:11000`.

- `k`

Set the key for encryption the data, the default is `8SMEE7ieNjSWVFqq`, and **key's length must be `128bit`, `192bit` or `256bit`**.

## Expection

This is the original version, there may be some mistakes, hope the ace coder can point it out.
