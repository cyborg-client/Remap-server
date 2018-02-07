import socket


# implement HTTP
# def Main():
#     host = "127.0.0.1"
#     port = 5000

#     mySocket = socket.socket()
#     mySocket.bind((host,port))

#     mySocket.listen(1)
#     conn, addr = mySocket.accept()
#     print ("Connection from: " + str(addr))
#     while True:
#             data = conn.recv(1024).decode()
#             if not data:
#                     break
#             print ("from connected  user: " + str(data))

#             data = str(data).upper()
#             print ("sending: " + str(data))
#             conn.send(data.encode())

#     conn.close()


# without HTTP
def main():
    host = '127.0.0.1'
    port = 8001

    # connect to server
    sock = socket.socket()
    sock.connect((host, port))

    message = input("--> ")

    while message != "q":
        sock.send(message.encode('utf-8'))
        data = sock.recv(1024).decode('utf-8')
        print("Recieved from server:", data)
        message = input("--> ")

    sock.close()


main()