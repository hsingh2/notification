---
apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: vms
  name: xnotif
spec:
  replicas: {{ deployment_mode_env[deployment_mode|lower]['replica_count']['xnotif'] }}
  selector:
    matchLabels:
      app: xnotif
      group: platformms
      consul-gossip: allow
  template:
    metadata:
      name: xnotif
      labels:
        app: xnotif
        group: platformms
        consul-gossip: allow
      annotations:
        fluentbit.io/parser: gomsx
    spec:
      affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - weight: 100
            podAffinityTerm:
              labelSelector:
                matchExpressions:
                - key: app
                  operator: In
                  values:
                  - xnotif
              topologyKey: kubernetes.io/hostname
{% if cloud == 'aws' %}
          - weight: 100
            podAffinityTerm:
              labelSelector:
                matchExpressions:
                - key: app
                  operator: In
                  values:
                  - xnotif
              topologyKey: topology.kubernetes.io/zone
{% endif %}
      containers:
        - name: consul
          image: {{ consul_image }}:{{ consul_version }}
          command:
            - consul
            - agent
            - -bind=0.0.0.0
            - -client=0.0.0.0
            - -datacenter={{ consul_dc }}
            - -retry-join=consul.service.consul
            - -data-dir=/consul/data
            - -config-dir=/consul/config
          volumeMounts:
            - mountPath: /consul/config
              name: phi
        - name: xnotif
          image: {{ xnotif_image }}:{{ xnotif_version }}
          command:
            - "/usr/bin/xnotif"
            - --profile
            - production
          livenessProbe:
            httpGet:
              path: /notification/admin/alive
              port: 9213
            initialDelaySeconds: 300
            periodSeconds: 30
          readinessProbe:
            httpGet:
              path: /notification/admin/health
              port: 9213
            initialDelaySeconds: 60
            periodSeconds: 30
          resources:
            requests:
              memory: "64Mi"
              cpu: "500m"
            limits:
              memory: "256Mi"
              cpu: "2000m"
          env:
            - name: PROFILE
              value: production
            - name: SPRING_CLOUD_CONSUL_HOST
              value: "localhost"
            - name: SPRING_CLOUD_CONSUL_SCHEME
              value: "{{ vault_scheme }}"
            - name: SPRING_CLOUD_CONSUL_PORT
              value: "8500"
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
                  name: msxvault
                  key: token
            - name: SPRING_REDIS_SENTINEL_ENABLE
              value: "true"
          ports:
            - containerPort: 9213
          volumeMounts:
            - mountPath: /keystore
              name: keystore
            - mountPath: /etc/ssl/certs/ca-certificates.crt
              name: certs
      volumes:
        - name: phi
          configMap:
            name: phi
        - name: keystore
          hostPath:
            path: /data/vms/keystore/
        - name: certs
          hostPath:
            path: /etc/ssl/certs/ca-bundle.crt
