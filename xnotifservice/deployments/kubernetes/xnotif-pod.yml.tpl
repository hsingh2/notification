---
apiVersion: v1
kind: Pod
metadata:
  namespace: "{{ kubernetes_namespace }}"
  name: xnotif
  labels:
    app: xnotif
    group: platformms
spec:
  restartPolicy: Never
  containers:
    - name: xnotif
      image: {{ xnotif_image }}:{{ xnotif_version }}
      command:
        - /usr/bin/xnotif
        - --profile
        - production
        - migrate
      resources:
        requests:
          cpu: "{{ 1*deployment_mode_env[deployment_mode|lower]['replica_count']['xnotif'] }}"
      env:
        - name: SPRING_CLOUD_CONSUL_HOST
          value: "consul.service.consul"
        - name: SPRING_CLOUD_CONSUL_PORT
          value: "8500"
        - name: SPRING_CLOUD_CONSUL_SCHEME
          value: "{{ vault_scheme }}"
        - name: SPRING_CLOUD_CONSUL_CONFIG_ACLTOKEN
          valueFrom:
            secretKeyRef:
              name: msxconsul
              key: token
        - name: SPRING_CLOUD_VAULT_HOST
          value: "vault.service.consul"
        - name: SPRING_CLOUD_VAULT_PORT
          value: "8200"
        - name: SPRING_CLOUD_VAULT_SCHEME
          value: "{{ vault_scheme }}"
        - name: SPRING_CLOUD_VAULT_TOKEN
          valueFrom:
            secretKeyRef:
              name: msxvault-idm
              key: token
