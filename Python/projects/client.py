## A very simple TCP client

import socket

with socket.socket() as s:
    s.connect(("127.0.0.1", 8080))
    s.sendall(b'Hello world')
    