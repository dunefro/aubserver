# Introduction
1. This tool send slack notification of failed pods
2. Receive notification of the failing pods with their status and the reason of failure on slack.
3. Simple to use. Create one k8s secret and then run the `make deploy` command to quickstart with.
# How to begin
1. In a cluster, create a secret named `slack-credentials` in the following way
```
kubectl create secret generic slack-credentials --from-literal=workspace-id=<your-workspace-id> --from-literal=slack-token=<your-slack-token>
```
2. Run the below command to start receiving alerts in the channel
```
make deploy
```
3. To test, build or play with the code run locally
```
make create
```
4. TO build an image for the same change the image name in the Makefile and `kubernetes/deploy.yaml` and then run
```
make build
make deploy
```
5. Now run the binary created in the `bin` folder via
```
./bin/aub

# or

make run
```
5. Code is written in such way that it can be executed from inside as well as outside the k8s cluster as well.