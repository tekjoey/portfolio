# A stupid simple unsecure webserver that makes a lot of assumptions and cuts a lot of corners.
# Seriously, this is not a production application. I'm just using it to learn about sockets, etc.
# Inspired by Coding a Web Server in 25 Lines - Computerphile -> https://www.youtube.com/watch?v=7GBlCinu9yg

# This project has been created to be a git submodule of the Python/projects directory. I've done this for a few reasons...
# 1) to learn about submodules
# 2) To learn about pushing to seperate remotes
# 3) to learn about using Atlasian products. This project has been pushed to a Bitbucket instance, and I (sorta) use Jira et all to track the tasks for this project.

import socket, pathlib, argparse

base_path = pathlib.Path(__file__).parent / 'www'

binary_file_types = ['ico', 'png', 'jpg', 'jpeg']

response_404 = """
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>HTTPY 404</title>
</head>
<body>
    404 Page Not Found
</body>
</html>"""

def main():
    parser = argparse.ArgumentParser(description='A terrible web server written in Python')
    parser.add_argument('-p', '--port', metavar='port', type=int, help='The port to serve the webserver on. Default to 8080', default=8080)
    parser.add_argument('-i', '--interface', metavar='interface', help="The interface for web server to listen on. Default is to bind to all interfaces.", default='')
    args = vars(parser.parse_args())
    
    listen_port = args['port']
    listen_address = args['interface']
    
    host = (listen_address, listen_port)
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.bind(host)
        s.listen(5)
        print(f"Listening on port {listen_port}")
        print(f"Listening on {'all interfaces' if listen_address == '' else f'interface {listen_address}'}")
        while True:
            try:
                conn, addr = s.accept()
                with conn:
                    print('Connected by', addr)
                    data = conn.recv(1024).decode("utf-8")
                    headers = parse_headers(data)
                    if headers['verb'] == "GET":
                        response_headders = "HTTP/1.1 200 OK\r\n"
                        location = "index.html" if headers["location"] == '/' else headers["location"][1:]
                        location_file_type = location.split('.')[-1]
                        print(f"Location is {location}")
                        print(f"trying to open {location}")
                        
                        match location_file_type:
                            case binary if binary in binary_file_types:
                                print("Location is binary file type")
                                try:
                                    with open(f"{base_path}/{location}", 'rb') as f:
                                        print(f"Successfully opened {location}")
                                        response_content = f.read()
                                except FileNotFoundError:
                                    print(f"Failed to open {location}, using 404 default instead")
                                    response_headders = 'HTTP/1.1 404  Not Found'
                                    response_content = response_404
                                conn.send(response_headders.encode("utf-8"))
                                conn.send('\r\n'.encode("utf-8")) # to separate headers from body
                                conn.send(response_content)
                            case _:
                                print("Location is NOT a binary file type")
                                try:
                                    with open(f"{base_path}/{location}", 'r') as f:
                                        response_content = f.read()
                                except FileNotFoundError:
                                    print(f"Failed to open {location}, using 404 default instead")
                                    response_headders = 'HTTP/1.1 404  Not Found'
                                    response_content = response_404
                                conn.send(response_headders.encode("utf-8"))
                                conn.send('\r\n'.encode("utf-8")) # to separate headers from body
                                conn.send(response_content.encode("utf-8"))
            except KeyboardInterrupt:
                ConnectionAbortedError
                s.close()
                break
                

def parse_headers(data: str) -> dict:
    headers = {}
    list_version = data.split("\r\n")
    verb, location, http = list_version[0].split(" ")
    headers['verb'] = verb
    headers['location'] = location
    headers["http version"] = http
    for option in list_version[1:]:
        if option == '':
            continue
        ls = option.split(": ")
        headers[ls[0]] = ls[1]

        
    # headers[field] = option
    return headers

if __name__ == "__main__":
    main()