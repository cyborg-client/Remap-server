import socket


def main():
    host = '127.0.0.1'
    port = 8001

    sock = socket.socket()
    sock.bind((host, port)) # binding a socket to a machine

    sock.listen(1) # one connection at a time
    client, address = sock.accept() # accept a connection
    print("Connection from:", str(address))

    while True:
        # recieve 1024 bytes at a time
        recv_data = client.recv(1024).decode('utf-8')

        if not recv_data:
            break

        print("Data from client:", recv_data)

        # send back
        client.send(recv_data.upper().encode('utf-8'))

    client.close()


main()
