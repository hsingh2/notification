---
apiVersion: policy/v1beta1
kind: PodDisruptionBudget
metadata:
  labels:
    app: xnotif
    group: platformms
  name: xnotif-pdb
  namespace: {{ kubernetes_namespace }}
spec:
  maxUnavailable: 1
  selector:
    matchLabels:
      app: xnotif
