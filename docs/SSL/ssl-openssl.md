```
# 1. Extract the .key file from the .pfx file:
openssl pkcs12 -in pfx-filename.pfx -nocerts -out key-filename.key

#2. Decrypt the .key file:
openssl rsa -in key-filename.key -out key-filename-decrypted.key

# 3. Extract the .crt file from .pfx file:
openssl pkcs12 -in pfx-filename.pfx -clcerts -nokeys -out crt-filename.crt

# 4. Create a secret in your Kubernetes cluster:
kubectl create secret tls new-ssl-cert --cert crt-filename.crt --key key-filename-decrypted.key

# 5. Verify that your new secret exists in your clusters namespace:
kubectl get secret -n new-ssl-cert
```

