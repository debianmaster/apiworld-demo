# -*- mode: Python -*-
allow_k8s_contexts('k3d-dev')
k8s_yaml('golang/deploy/deployment.yaml')
k8s_yaml('golang/deploy/service.yaml')

local_resource('make-golang-api','make golang-api',['./golang/server.go'])

docker_build('quay.io/coolcompany/apiworld-golang-api','./golang',
    dockerfile='golang/Dockerfile.dev',
    live_update=[
        sync('golang/bin/golang-api','/app'),
        run('/app')
    ]
)







