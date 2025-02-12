CONTAINER CREATION

docker build -t adarshinipriyaa/go-web-app:v1 .

CONTAINER RUNNING IN THE LOCAL HOST

docker run -p 8080:8080 -it adarshinipriyaa/go-web-app:v1

PUSHING THE IMAGE TO THE DOCKER REPOSITORY

docker push adarshinipriyaa/go-web-app:v1

==================================================================
k8s

CREATE A DEPLOYMET, SERVICE, INGRESS FILE
take the refernce from k8s official documentation

INSTALL KUBECTL, EKSCTL, AWS CLI

aws configure
eksctl create cluster --name demo-cluster --region ap-south-1

eksctl create cluster --name adarshini --region ap-south-1 --version 1.31 #version

eksctl delete cluster --name demo-cluster --region ap-south-1 #for deleting

eksctl delete cluster --name adarshini --region ap-south-1

kubectl apply -f .\deployment.yml

kubectl apply -f .\service.yml

kubectl apply -f .\ingress.yml

kubectl get all

kubectl edit svc go-web-app

change cluserIP --> NodePort

kubectl get svc

kubectl get nodes -o wide

port forward

kubectl port-forward svc/go-web-app 8080:80

http://localhost:8080/

curl http://localhost:8080

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/cloud/deploy.yaml #INGRESS controller

kubectl get ing

nslookup ELB --> nslookup a2518d583b7294335a555194c851a6cf-1089494526.ap-south-1.elb.amazonaws.com

---

Server: UnKnown
Address: fe80::1
Non-authoritative answer:
Name: a2518d583b7294335a555194c851a6cf-1089494526.ap-south-1.elb.amazonaws.com
Address: 13.232.217.71

===================================================================================

C:\Windows\System32\drivers\etc\hosts #add the address in this path
13.232.217.71 go-web-app.local

===================================================================================
Installing helm

choco install kubernetes-helm

cd helm

helm create go-web-app-chart

C:\Users\Adarshini\Desktop\Dev+ops\go-http-webpage\helm> helm install go-web-app go-web-app-chart
kubectl edit deploy go-web-app
helm uninstall go-web-app #uninstall

===================================================================================
