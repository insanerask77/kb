```yaml
# Local
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: {{ .Chart.Name }}-ingress
  namespace: {{ $NAMESPACE }}
  annotations:
    kubernetes.io/ingress.class: nginx
    nginx.ingress.kubernetes.io/rewrite-target: /$1
    nginx.ingress.kubernetes.io/proxy-body-size: "100m"
    # nginx.ingress.kubernetes.io/client_max_body_size: "0"
    nginx.ingress.kubernetes.io/client-body-buffer-size: "1m"
    nginx.ingress.kubernetes.io/proxy-buffering: "on"
    nginx.ingress.kubernetes.io/proxy-buffers-number: "4"
    nginx.ingress.kubernetes.io/proxy-buffer-size: "1m"
    nginx.ingress.kubernetes.io/proxy-max-temp-file-size: "1024m"
    nginx.ingress.kubernetes.io/proxy-next-upstream: timeout
    nginx.ingress.kubernetes.io/proxy-read-timeout: '1003'
spec:
  tls:
    - hosts:
        - {{ include "common.ingress.url" . }}
      secretName: {{ .Values.ingress.secrets }}
  rules:
    - host: {{ include "common.ingress.url" . }}
      http:
        paths:
          - path: {{ .Values.ingress.path | toString }}
            pathType: {{ .Values.ingress.pathType }}
            backend:
              service:
                name: {{ include "helper.names.fullname" . }}
                port:
                  number: {{ .Values.global.net.ports.frontend }}
```

```yaml
# Local
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: {{ .Chart.Name }}-ingress
  namespace: {{ $NAMESPACE }}
  annotations:
    kubernetes.io/ingress.class: nginx
    nginx.ingress.kubernetes.io/rewrite-target: /$1
spec:
  rules:
    - host: backend.mybts.com
      http:
        paths:
          - path: /api/(.*)
            pathType: Prefix
            backend:
              service:
                name: <SERVICE NAME>
                port:
                  number: <SERVICE PORT | iota >
```

