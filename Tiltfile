# -*- mode: Python -*-
allow_k8s_contexts('k3d-dev')
k8s_yaml('golang/deploy/deployment.yaml')
k8s_yaml('golang/deploy/service.yaml')

docker_build('quay.io/chak/golang-api','./golang',
    dockerfile='golang/Dockerfile.dev',
    only=['bin/golang-api'],
    entrypoint='/app',
    live_update=[
        sync('golang/bin/golang-api','/app')
    ]
)


local_resource('make-golang-api','make golang-api',['./golang/server.go'])
#trigger_mode=TRIGGER_MODE_MANUAL








