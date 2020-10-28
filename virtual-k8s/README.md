helm install vk8s1 virtualcluster --repo https://charts.devspace.sh/ \
  --namespace vk8s1 \
  --values values.yaml \
  --create-namespace \
  --wait

kubectl -n vk8s1 exec vk8s1-0 -c syncer -- cat /root/.kube/config > vk8s1.yaml

k port-forward po/vk8s1-0 -n vk8s1 8443:8443 &

kubectl run --image=debianmaster/nodejs-welcome welcome --kubeconfig=vk8s1.yaml

kubectl get pods,ing --kubeconfig=vk8s1.yaml

kubectl get pods -A  --kubeconfig=vk8s1.yaml