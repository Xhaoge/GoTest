#!/usr/bin/python3
#coding=utf-8

from socket import *
import sys
import getpass

#创建网络连接
def main():
    if len(sys.argv)<3:
        print('argv is error')
        return
    HOST = sys.argv[1]
    PORT = int(sys.argv[2])
    ADDR = (HOST,PORT)
    #创建tcp套接字
    s = socket()
    try:
        s.connect(ADDR)
    except Exception as e:
        print(e)
        return
    while True:
        print('''
              ========请先选择操作======== 
              --1.注册   2.登录   3.退出-- 
              ==========================
              ''')
        try:
            cc = int(input('请先选择操作>>'))
        except Exception as e:
            print(e)
            continue
        if cc not in [1,2,3]:
            print('请选择正确选项:')
            sys.stdin.flush() #清除标准输入
            continue

        elif cc == 1:
            d = do_register(s)
            if d == 0:
                print('注册成功...')
                #进入登录页面
            elif d == 1:
                print('该用户已存在，请重新输入...')
                continue
            else:
                print('注册失败...')

        elif cc == 2:
            name = do_login(s)
            if name:
                print('登录成功....')
                do_select(s,name)#进入查询界面
            else:
                print('用户名或密码错误...')
                continue

        elif cc == 3:
            s.send(b'E')
            sys.exit('谢谢使用,欢迎下次使用...')

def do_register(s):
    while True:
        name = input('User:')
        password = getpass.getpass()
        password1 = getpass.getpass('Again:')
        if password != password1:
            print('俩次密码不一致，请重新输入...')
            continue
        if " " in name or ' ' in password:
            print('账号和密码中不允许存在空格...')
            continue
        msg = 'R {} {}'.format(name,password)
        s.send(msg.encode())
        data = s.recv(1024).decode()
        if data == 'OK':
            return 0
        elif data == 'EXISTS':
            return 1
        else:
            return 2
def do_login(s):
    name = input('请输入账号：')
    passwd = getpass.getpass()
    msg = 'L {} {}'.format(name,passwd)
    s.send(msg.encode())
    data = s.recv(128).decode()
    if data == 'OK':
        return name
    else:
        return

def do_query(s,name):
    while True:
        word = input('请输入要查询的单词:')
        if word == '##':
            break
        msg = 'Q {} {}'.format(name,word)
        s.send(msg.encode())
        data = s.recv(128).decode()
        if data == 'OK':
            data = s.recv(2048).decode()
            print(data)
        else:
            print('没有查到该单词...')

def do_hist(s,name):
    msg = 'H {}'.format(name)
    s.send(msg.encode())
    data = s.recv(128).decode()
    if data == 'OK':
        while True:
            data = s.recv(1028).decode()
            if data == '##':
                break
            print(data)
    else:
        print("没有查到该单词历史记录....")

def do_select(s,name):
    while True:
        print('''
            ==========查询界面==========
            1.查词    2.历史记录   3.退出
            ===========================
            ''')
        try:
            ss = int(input('请输入操作选项>>'))
        except Exception as e:
            print(e)
            continue
        if ss not in [1,2,3]:
            sys.stdin.flush()
            print('请输入正确操作>>')
            continue

        elif ss == 1:
            do_query(s,name)
            continue
        elif ss == 2:
            do_hist(s,name)
            continue
        elif ss == 3:
            break


if __name__ == "__main__":
    main()