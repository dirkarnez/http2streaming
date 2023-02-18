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
