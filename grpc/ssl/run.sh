protoc --go_out=. simple.proto
protoc --go-grpc_out=. simple.proto

# 先去编辑openssl.cnf 文件
# 生成server密钥
openssl genrsa -out server.key 2048
# 生成csr请求文件
openssl req -new -nodes -key ./server.key -out ./server.csr -config openssl.cnf -extensions 'v3_req'
# 生成server公钥
openssl x509 -req -in ./server.csr -out ./server.pem -CAcreateserial -extfile ./openssl.cnf -extensions 'v3_req' -signkey ./server.key
