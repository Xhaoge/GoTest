import os
from time import sleep

i = os.fork()

if i < 0:
    print('failed')
elif i ==0:
    sleep(3)
    print('son process exit',os.getpid())
    os.exit()
else:
    pid,status = os.wait()
    print(pid,status)
    print(os.WEXITSTATUS(status))#获取子进程退出状态
    while True:
        pass