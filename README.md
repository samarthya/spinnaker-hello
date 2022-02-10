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