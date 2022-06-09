# Kyverno Automate Performance Test

Kyverno is a Kubernetes native policy engine that secures and automates Kubernetes configurations. 
This project automates scalability tests for Kyverno on large Kubernetes clusters with several namespaces and resources


## Test scenario
This test scenario will create loads of Kubernetes objects (Pod, Namespace, Deployment, Cronjob, ConfigMap, Secret) based on user defined scale

Scales mapping:
```
  xs: 100 total resource
  small: 500 total resource
  medium: 1000 total resource
  large: 2000 total resource 
  xl: 3000 total resource
  xxl: 4000 total resource
```


## Automation Phases

  Phase 1: Collecting Kyverno usage started (it's run concurrently with the next phases)

  Phase 2: Creating resources based on scales provided (steps up)

  Phase 3: Idle process (to ensure the workloads has been successfully created and ready)

  Phase 4: Deleting a half of resources (steps down)

  Phase 5: Idle process (to get performance behaviour after steps up & down process)

  Phase 6: Report Generated
  

## Getting Started

```
git clone https://github.com/husnialhamdani/kyvernop.git
cd kyvernop
go build .
```

### Setup Environment
to create Kubernetes cluster with Kyverno base policies
```
bash setup.sh
``` 

to install additional policy you pass the url's to the script, for example:
```
bash setup.sh https://raw.githubusercontent.com/kyverno/policies/main/best-practices/require_probes/require_probes.yaml
``` 


### Start automation
create resources based on pre-defined scales:
```
./kyvernop execute --scale medium
``` 
create custom resources quantity based on user defined:
```
./kyvernop execute --number 5000
``` 

Cleanup
```
./kyvernop cleanup -size 500
```

## Anomaly Detection

Isolation Forest is an algorithm that detects anomalies by taking a subset of data and constructing many isolation trees out of it.

An isolation tree is constructed by randomly selecting a feature and randomly selecting a value from that feature. A forest is constructed by aggregating all the isolation trees.

We pass the the Kyverno usage as input data and this algorithm will provide a prediction, The isolation forest assigns 0 to the anomalous data and 1 to the normal data and finally it plot the anomalies predicted by Isolation forest.


## Report

After the automation has completed, the tools will automatically generate a report based on Kyverno performance behaviour during the test and using the algorithm mentioned above.

![alt text](https://github.com/husnialhamdani/kyvernop/blob/main/report.png?raw=true)

## Kyverno Consideration on large cluster

[Scales]  [Expected behaviour]

### Recommendations setup
#### Kyverno Resource Limit config:

Kyverno limit config:

less than 4000 total resource: use default limit
more than 4000 total resource: use 768Mi (2x the default)

#if the workloads more than 2000
- Memory limit
- kyverno instance 

#production cluster
- kyverno HA mode (3 or 5 replicas for minimum)

* total resources : combination of static (ConfigMap & Secret) and workloads (Pod, Cronjob, Deployment)