'''
name:小浩ge
date:2018-9-29
email:xxx@qq.com
readme:this is a dict project for practice
'''

from socket import *
import os
import sys
import time
import pymysql
import signal

#定义服务器需要的地址端口
HOST = '127.0.0.1'
PORT = 8000
ADDR = (HOST,PORT)

#流程控制
def main():
    #创建tcp套接字
    s = socket()
    s.setsockopt(SOL_SOCKET,SO_REUSEADDR,1)
    s.bind(ADDR)
    s.listen(5)

    #创建数据库连接
    db = pymysql.connect('localhost','root', '123456','dict_project')
    #忽略子进程信号
    signal.signal(signal.SIGCHLD,signal.SIG_IGN)
    #流程控制
    while True:
        try:
            c,addr = s.accept()
            print("Connect from",addr)
        except KeyboardInterrupt:
            s.close()
            sys.exit('服务端已退出....')
        except Exception as e:
            print(e)
            continue
        #创建子进程
        pid = os.fork()
        if pid == 0:
            # s.close() c.close()
            do_child(c,db)
        else:
            continue

def do_child(c,db):
    while True:
        data = c.recv(128).decode()
        print(c.getpeername(),":",data)
        if data[0] == 'R':
            do_register(c,db,data)#注册，检验是否存在该用户
        elif data[0] == 'L':
            do_login(c,db,data)#登录界面操作
        elif data[0] == 'Q':
            do_query(c,db,data)#查询单词操作
        elif data[0] == 'H':
            do_hist(c,db,data)#查询历史记录
        elif data[0] == 'E' or (not data):
            c.close()
            sys.exit(0)

def do_register(c,db,data):
    print('do_register,注册操作')
    l = data.split(" ")
    name = l[1]
    passewd = l[2]
    cursor = db.cursor()
    sql = "select * from user where u_name='%s';"%name
    cursor.execute(sql)
    res = cursor.fetchone()

    if res != None:
        c.send(b'EXISTS')
        return
    sql = "insert into user (u_name,u_pd) \
             values ('%s','%s');"%(name,passewd)
    try:
        cursor.execute(sql)
        db.commit()
        c.send(b'OK')
    except:
        db.rollback()
        c.send(b'FALL')
    else:
        print("%s注册成功"%name)

def do_login(c,db,data):
    print('do_register,登录操作')
    l = data.split(" ")
    name = l[1]
    passewd = l[2]
    cursor = db.cursor()
    sql = "select* from user where u_name='%s' and u_pd='%s';"\
            %(name,passewd)
    cursor.execute(sql)
    res = cursor.fetchone()
    if res != None:
        print('%s登录成功'%name)
        c.send(b'OK')
    else:
        print('登录失败')
        c.send(b'FALL')

def do_query(c,db,data):
    print('do_query 查询操作')
    l = data.split(" ")
    name = l[1]
    wd = l[2]
    cursor = db.cursor()
    sql = "select trans from dict where word='%s';"%wd
    cursor.execute(sql)
    res = cursor.fetchone()
    if res == None:
        c.send(b'FALL')
        return
    elif res == "##":
        return
    else :
        c.send(b'OK')
        time.sleep(0.1)
        c.send(res[0].encode())
        sql = "insert into hist (name,record) values ('%s','%s'); "%(name,wd)
        cursor.execute(sql)
        db.commit()
        return

def do_hist(c,db,data):
    print('do_hist')
    l = data.split(' ')
    name = l[1]
    cursor = db.cursor()
    sql = "select record from hist where name='%s';"%name
    cursor.execute(sql)
    result = cursor.fetchall()
    if result == None:
        c.sent(b'FALL')
        return
    else:
        c.send(b'OK')

        for i in result:
            time.sleep(0.1)
            msg = "%s"%(i[0])
            c.send(msg.encode())
    time.sleep(0.1)
    c.send(b'##')

if __name__ == '__main__':
    main()

