from socket import *
import os

s = socket()

sock_file = './sock_file'
if os.path.exists(sock_file):
    os.remove()

sockfd = socket(AF_UNIX,SOCK_STREAM)

sockfd.bind(sock_file)
sockfd.listen(3)

while True:
    c,addr = sockfd.accept()
    while True:
        data = c.recv(1024)
        if data:
            print(data.decode())
            c.send(b'receive')
        else:
            break
    c.close()
sockfd.close()
