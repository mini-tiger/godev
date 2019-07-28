
cd /home/go/src/godev/mymodels/grpc/CutstomTLS/perm
openssl genrsa -out server.key 2048

openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650

Country Name (2 letter code) [XX]:
State or Province Name (full name) []:
Locality Name (eg, city) [Default City]:
Organization Name (eg, company) [Default Company Ltd]:
Organizational Unit Name (eg, section) []:
Common Name (eg, your name or your server's hostname) []:tj-test todo 这里是 认证时候的名字
Email Address []: