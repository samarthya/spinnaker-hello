# Simple Rest End point

 Simply runs the docker image

 ```bash
 docker run --env HTTP_PORT=8181 -i -t --rm -p 8181:8181 spinnaker:v1.0
 ```

It uses Echo refer to [easy setup](https://github.com/labstack/echo) as below

```bash
package main

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "net/http"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", hello)

  // Start server
  e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}
```

## Console

 ![](./images/console.png)

 ## Ping

 ![](./images/ping.png)

 /* pipeline {
  environment {
    imagename = "samarthya/spinnaker-hellow"
    registryCredential = 'docker-samarthya'
    dockerImage = ''
  }
  
  agent {
    kubernetes {
      yaml '''
        apiVersion: v1
        kind: Pod
        spec:
          containers:
          - name: golang
            image: golang:latest
            command:
            - cat
            tty: true
      '''
    }
  }


  stages {
    stage('Clone & Build') {
      steps {
        container('goalng'){
          sh 'echo hello'
        }
          // echo 'git clone'
          // git([url: 'https://github.gwd.broadcom.net/ss670121/spinnaker-hellow', branch: 'main', credentialsId: 'pat-ss670121'])
          // sh '''
          // ls
          // go build -o bin/spinnaker-hellow .
          // '''
      }
    }

  }
}
*/



/*
pipeline {
    environment {
      imagename = "samarthya/spinnaker-hellow"
      registryCredential = 'docker-samarthya'
      dockerImage = ''
    }

    agent none
    
    stages {
      
      stage('Cloning Git') {
          agent { label 'GOLANG-AGENT' }
        steps {
          echo 'git clone'
          git([url: 'https://github.gwd.broadcom.net/ss670121/spinnaker-hellow', branch: 'main', credentialsId: 'pat-ss670121'])
        }
      }
      
      stage('Build Binary') {
          agent { label 'GOLANG-AGENT' }
        steps {
          echo 'go build'
          script{
            go version
            go build -o spinnaker-hellow ./
          }
        }
      }
      
      stage('Building image') {
          agent {
              label 'DOCKER-AGENT'
        }
        steps{
          script {
            dockerImage = docker.build imagename
          }
        }
      }
      
      stage('Deploy Image') {
          agent {
              label 'DOCKER-AGENT'
        }
        steps{
          script {
            docker.withRegistry( '', registryCredential ) {
              dockerImage.push("$BUILD_NUMBER")
              dockerImage.push('latest')
            }
          }
        }
      }
      
      stage('Remove Unused docker image') {
          agent {
              label 'DOCKER-AGENT'
        }
        steps{
          sh "docker rmi $imagename:$BUILD_NUMBER"
          sh "docker rmi $imagename:latest"
        }
      }
    }
  }
  */

  ```bash
  NAME: my-release
LAST DEPLOYED: Fri Feb 11 07:16:49 2022
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
CHART NAME: kafka
CHART VERSION: 15.1.0
APP VERSION: 3.1.0

** Please be patient while the chart is being deployed **

Kafka can be accessed by consumers via port 9092 on the following DNS name from within your cluster:

    my-release-kafka.default.svc.cluster.local

Each Kafka broker can be accessed by producers via port 9092 on the following DNS name(s) from within your cluster:

    my-release-kafka-0.my-release-kafka-headless.default.svc.cluster.local:9092

To create a pod that you can use as a Kafka client run the following commands:

    kubectl run my-release-kafka-client --restart='Never' --image docker.io/bitnami/kafka:3.1.0-debian-10-r14 --namespace default --command -- sleep infinity
    kubectl exec --tty -i my-release-kafka-client --namespace default -- bash

    PRODUCER:
        kafka-console-producer.sh \
            --broker-list my-release-kafka-0.my-release-kafka-headless.default.svc.cluster.local:9092 \
            --topic test

    CONSUMER:
        kafka-console-consumer.sh \
            --bootstrap-server my-release-kafka.default.svc.cluster.local:9092 \
            --topic test \
            --from-beginning
```