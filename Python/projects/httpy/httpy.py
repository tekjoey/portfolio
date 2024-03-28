# A stupid simple unsecure webserver that makes a lot of assumptions and cuts a lot of corners.
# Seriously, this is not a production application. I'm just using it to learn about sockets, etc.
# Inspired by Coding a Web Server in 25 Lines - Computerphile -> https://www.youtube.com/watch?v=7GBlCinu9yg
# SoliDeoGloria

import socket, pathlib, argparse, datetime

from urllib.parse import urlparse

base_path = pathlib.Path(__file__).parent / 'www'

# Define parsing characters as per RFC 2616 2.2
CRLF = '\r\n'       # Carriage return & line feed
LWS = f'{CRLF} '    # Linear white space

http_version = "HTTP/1.1"

binary_files = {
    '.ico': 'image/vnd.microsoft.icon',
    '.png': 'image/png',
    '.jpg': 'image/jpeg',
    '.jpeg': 'image/jpeg',
}

file_types = {
    '.html': 'text/html',
    '.css': 'text/css',
    '.js': 'text/javascript'
} | binary_files

response_404 = (
    f'{http_version} 404  Not Found{CRLF}',     #⬅️ Headers
                                                #⬇️ HTML Response Body; TODO, make this an actual HTML file
    """
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
)

def main():
    parser = argparse.ArgumentParser(description='A terrible web server written in Python')
    parser.add_argument('-p', '--port', metavar='port', type=int, help='The port to serve the webserver on. Default to 8080', default=8080)
    parser.add_argument('-i', '--interface', metavar='interface', help="The interface for web server to listen on. Default is to bind to all interfaces.", default='')
    args = vars(parser.parse_args())
    
    listen_port = args['port']
    listen_address = "192.168.100.10" #args['interface']
    
    host = (listen_address, listen_port)
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.bind(host)
        s.listen(5)
        print(f"Listening on port {listen_port}")
        print(f"Listening on {'all interfaces' if listen_address == '' else f'interface {listen_address}'}")
        print(s)
        while True:
            print("Accepting connections")
            try:
                print("begin try")
                conn, addr = s.accept()
                print("connection accepted!")
                with conn:
                    print('Connected by', addr)
                    data = conn.recv(1024).decode("utf-8")
                    receive_headers = parse_headers(data)
                    if receive_headers['verb'] == "GET":
                        response_headers, response_content = get(receive_headers)
                        conn.send(response_headers.encode("utf-8"))
                        print("sent headers")
                        conn.send(CRLF.encode("utf-8")) # to separate headers from body
                        print("sent CRLF")
                        conn.send(response_content)
                        print("Sent body")
            except KeyboardInterrupt:
                ConnectionAbortedError
                s.close()
                break
                

def parse_headers(data: str) -> dict:
    print(data)
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
    print(headers)
    return headers

def get(recieve_headers:dict):
    header_date = "Date: " + datetime.datetime.strftime(datetime.datetime.now(datetime.UTC), '%a, %d %b %Y %H:%M:%S GMT')
    header_server = "Server: HTTPY/0.1"
    
    # HTTP 1.1 requires servers to reject clients who do not use the Host header field
    if recieve_headers['http version'] == 'HTTP/1.1' and "Host" not in recieve_headers.keys():
        print(recieve_headers)
        return f'{http_version} 400 Bad Request{CRLF}{header_date}{CRLF}{header_server}{CRLF}Content-Type: text/html{CRLF}Content-Language: en-US{CRLF}', """
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>HTTPY 400</title>
</head>
<body>
    <h2>No Host: header received</h2>
    <p>HTTP 1.1 requests must include the Host: header.</p>
</body></html>
""".encode('utf-8')

    # HTTP1.1 requires servers to be able to manage full URLs, not just pathnames
    url = urlparse(recieve_headers['location'])
    location = "index.html" if url.path == '/' else url.path[1:]
    location_file_type = pathlib.Path(location).suffixes[-1]
    
    print(f"Location is {location}")
    print(f"Trying to open {location}")

    try:
        with open(f"{base_path}/{location}", 'rb') as f:
            print(f"Successfully opened {location}")
            response_content = f.read()
            content_size = pathlib.Path(base_path/ location).stat().st_size
    except FileNotFoundError:
        print(f"'{location}' not found, using 404 default instead")
        return response_404
    
    if location_file_type not in binary_files:
        response_content = response_content.decode('utf-8').encode('utf-8')

    header_status = f"{http_version} 200 OK"
    header_content_type = "Content-Type: " + file_types[location_file_type]
    #header_content_length = f"Content-Length: {str(content_size + 101)} B"
    header = f'{header_status}{CRLF}{header_date}{CRLF}{header_server}{CRLF}{header_content_type}{CRLF}Content-Language: en-US{CRLF}'#{header_content_length}'

    return header, response_content

if __name__ == "__main__":
    main()