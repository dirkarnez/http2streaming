http2streaming
==============
### HTTPs
- Run in `PortableGit\git-bash.exe`
  ```bash
  openssl req -new -newkey rsa:2048 -days 365 -nodes -x509 -keyout server.key -out server.crt
  # or
  openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt
  # or
  openssl req -x509 -newkey rsa:2048 -nodes -sha256 -subj '/CN=localhost' -keyout localhost-privkey.pem -out localhost-cert.pem
  ```

### Reference
- [ReadableStream - Web APIs | MDN](https://developer.mozilla.org/en-US/docs/Web/API/ReadableStream)
- [davedoesdev/browser-http2-duplex: Full-duplex stream emulation over HTTP/2](https://github.com/davedoesdev/browser-http2-duplex)
- [HTTP/2 Streaming in Golang - Codemio - A Software Developer's Blog](https://www.codemio.com/2018/03/http2-streaming-golang.html)
- [**herrberk/go-http2-streaming: Example of HTTP/2 Streaming written in Golang**](https://github.com/herrberk/go-http2-streaming)
