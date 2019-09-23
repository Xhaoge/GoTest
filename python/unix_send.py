from socket import *

#确保通信俩端用的是同一个套接字文件'
sock_file = './sock_file'
sockfd = socket(AF_UNIX,SOCK_STREAM)

#连接另一端
sockfd.connect(sock_file)

while True:
    msg = input('>>')
    if msg:
        sockfd.send(smg.encode())
        print(sockfd.recv(1024))
    else:
        break
sockfd.close()