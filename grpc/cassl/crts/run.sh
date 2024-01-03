# 先去编辑openssl.cnf 文件
# 生成ca证书密钥
openssl genrsa -out ca.key 2048
# 生成ca csr请求文件
openssl req -new -key ca.key -out ca.csr
# 生成ca公钥
openssl x509 -req -days 3650 -in ca.csr -signkey ca.key -out ca.pem


# 给server颁发证书
openssl genrsa -out server.key 2048
openssl req -new -nodes -key ./server.key -out ./server.csr -config openssl.cnf -extensions 'v3_req'
openssl x509 -req -in ./server.csr -out ./server.pem -CA ca.pem -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions 'v3_req'

# 给client颁发证书
openssl genrsa -out client.key 2048
openssl req -new -nodes -key ./client.key -out ./client.csr -config openssl.cnf -extensions 'v3_req'
openssl x509 -req -in ./client.csr -out ./client.pem -CA ca.pem -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions 'v3_req'
