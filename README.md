1. It should scan all the pods [DONE]
2. List all the pods in the running state and display them in the application output [REMOVED]
3. List all the non-running pods in the application output [REMOVED]
4. Send notification of all the pods that are not in the running state. [DONE]
5. Generate a red symbol for all of them. [DONE]

# How to begin
1. In a cluster, create a secret named `slack-credentials` in the following way
```
kubectl create secret generic slack-credentials --from-literal=workspace-id=<your-workspace-id> --from-literal=slack-token=<your-slack-token>
```
2. Run the below command to start receiving alerts in the channel
```
make deploy
```
3. To test, build or play with the code run
```
make create
```
4. Now run the binary created in the `bin` folder via
```
./bin/aub

# or

make run
```
5. Code is written in such way that it can be executed from inside as well as outside the k8s cluster as well.