import socket
import os, os.path
import time
import sys
from collections import deque

server_socket_file = "/tmp/test_unix_socket.server"
client_socket_file = "/tmp/test_unix_socket.client"

def run_server():
    if os.path.exists(server_socket_file):
        os.unlink(server_socket_file)

    server = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
    server.bind(server_socket_file)
    server.listen(1000)
    print("Staring server accept loop...")
    while True:
        conn, addr = server.accept()
        try:
            msg = conn.recv(10)
            print(f"Server received data {msg}")
            conn.sendall(msg)
        except BrokenPipeError as e:
            print(f"broken pipe. Client {addr} probably dropped the connection. ")
        finally:
            conn.close()

def run_client(msg: str):
    if os.path.exists(client_socket_file):
        os.unlink(client_socket_file)

    sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
    sock.connect(server_socket_file)
    try:
        print(f"Send message: {msg}")
        sock.sendall(msg.encode("utf-8"))
        response = sock.recv(10)
        print(f"client received response {response}")
    finally:
        sock.close()

if __name__ == "__main__":
    assert len(sys.argv) > 1
    mode = sys.argv[1]
    if mode == "server":
        run_server()
    elif mode == "client":
        assert len(sys.argv) > 2
        run_client(sys.argv[2])
    else:
        raise ValueError(f"Wrong mode {sys.argv[1]}")
