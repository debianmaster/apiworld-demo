# -*- mode: Python -*-
allow_k8s_contexts('k3d-dev')
load('ext://restart_process', 'docker_build_with_restart')

k8s_yaml('golang/deploy/deployment.yaml')
k8s_yaml('golang/deploy/service.yaml')

docker_build_with_restart('quay.io/chak/golang-api','./golang',
    dockerfile='golang/Dockerfile.dev',
    only=['bin/golang-api'],
    entrypoint='/app',
    live_update=[
        sync('golang/bin/golang-api','/app')
    ]
)
local_resource('make-golang-api','make golang-api',['./golang/server.go'])



k8s_yaml('nodejs/deploy/deployment.yaml')
k8s_yaml('nodejs/deploy/service.yaml')

docker_build('quay.io/chak/nodejs-api','./nodejs',
    dockerfile='nodejs/Dockerfile.dev',  
    live_update=[
        sync('nodejs', '/app'),
    ]
)


k8s_yaml('react-ui/deploy/deployment.yaml')
k8s_yaml('react-ui/deploy/service.yaml')


docker_build('quay.io/chak/react-ui','./react-ui',
    dockerfile='react-ui/Dockerfile'
)


k8s_yaml('common/deploy/ingress.yaml')