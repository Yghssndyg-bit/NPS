import socket
def send_hex_strings_to_server(host, port, hex_strings):    
    try:        
        # 创建 TCP 连接        
        conn = socket.socket(socket.AF_INET, socket.SOCK_STREAM)        
        conn.connect((host, port))        
        print(f"[*] Connected to {host}:{port}")                
        for hex_string in hex_strings:            
            # 将十六进制字符串转换为字节数据           
            data = bytes.fromhex(hex_string)            
            conn.sendall(data)            
            print(f"[*] Sent: {hex_string} (hex) -> {data} (bytes)")                
        # 接收远程主机返回的信息        
        response = conn.recv(1024)        
        print(f"[+] Received: {response.hex()} (hex) -> {response} (bytes)")            
    except socket.error as e:        
        print(f"Socket error: {e}")    
    except Exception as e:        
        print(f"An error occurred: {e}")    
    finally:        
        conn.close()
# 远程主机的 IP 地址和端口
remote_host = '192.168.26.154'   # 请根据需要更改
remote_port = 8024  # 请根据需要更改
# 要发送的十六进制字符串
hex_strings_to_send = ['545354', '06000000302e32362e30', '07000000302e32362e3130']  # 分别表示 "TST", "1", "2" 的十六进制
# 调用函数
send_hex_strings_to_server(remote_host, remote_port, hex_strings_to_send)