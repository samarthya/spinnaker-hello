helm install haproxy haproxytech/kubernetes-ingress \
    --set controller.kind=DaemonSet \
    --set controller.daemonset.useHostPort=true

