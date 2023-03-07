mport socket
import time

def start_client():
    host = "127.0.0.1" # địa chỉ IP của máy chủ
    port = 4444 # cổng kết nối
    retry_delay = 5 # thời gian chờ giữa các lần kết nối lại

    while True:
        try:
            client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
            client.connect((host, port))
            print("[*] Connected to {}:{}".format(host, port))

            while True:
                command = input("Enter command: ")
                if not command:
                    continue
                client.sendall(command.encode('utf-8'))

                data = client.recv(1024).decode('utf-8')
                if not data:
                    print("[*] Connection closed by server")
                    break
                print(data)
        except Exception as e:
            print(str(e))
        finally:
            client.close()
            print("[*] Connection closed. Retrying in {} seconds...".format(retry_delay))
            time.sleep(retry_delay)

if __name__ == '__main__':
    start_client()