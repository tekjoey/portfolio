# A stupid simple unsecure webserver that makes a lot of assumptions and cuts a lot of corners.
# Seriously, this is not a production application. I'm just using it to learn about sockets, etc.
# Inspired by Coding a Web Server in 25 Lines - Computerphile -> https://www.youtube.com/watch?v=7GBlCinu9yg

import socket

def main():
    listen_address = ""
    listen_port = 8092
    counter = 0
    host = (listen_address, listen_port)
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.bind(host)
        s.listen(0)
        while True:
            conn, addr = s.accept()
            with conn:
                print('Connected by', addr)
                data = conn.recv(1024).decode("utf-8")
                headers = parse_headers(data)
                print(data)
                
                if headers['verb'] == "GET":
                    print('get')
                    text = f"Hello world {counter}"
                    conn.sendall(text.encode("utf-8"))
                    counter += 1

def parse_headers(data: str) -> dict:
    headers = {}
    list_version = data.split("\r\n")
    verb, location, http = list_version[0].split(" ")
    headers['verb'] = verb
    headers['location'] = location
    headers["http version"] = http
    print(list_version)
    for option in list_version[1:]:
        if option == '':
            continue
        ls = option.split(": ")
        headers[ls[0]] = ls[1]

        
        # headers[field] = option
        print(headers)
    return headers

if __name__ == "__main__":
    main()