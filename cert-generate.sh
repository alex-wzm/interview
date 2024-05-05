#!/bin/bash
cd certs
sudo apt-get update
sudo apt-get install openssl
openssl genpkey -algorithm RSA -out server-key.pem
openssl req -new -key server-key.pem -out server-csr.pem
openssl x509 -req -days 90 -in server-csr.pem -signkey server-key.pem -out server-crt.pem
openssl x509 -in server-crt.pem -text -noout
