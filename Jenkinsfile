podTemplate(yaml: '''
    apiVersion: v1
    kind: Pod
    spec:
      volumes:
      - name: dockersock
        hostPath:
          path: /var/run/docker.sock
      containers:
      - name: golang
        image: golang:latest
        command:
        - sleep
        args:
        - 99d
      - name: docker
        image: docker:latest
        volumeMounts:
          - name: dockersock
            mountPath: /var/run/docker.sock
        command:
        - sleep
        args:
        - 99d
        
''') {
  node(POD_LABEL) {
    // Define the image name
    def imageName = "samarthya/spinnaker"
    def buildNumber = env.BUILD_NUMBER

    stage('Get a Golang project') {
      git url: 'https://github.com/samarthya/spinnaker-hello.git', branch: 'main', credentialsId: 'github-samarthya'
      container('golang'){
        stage('Build a Go project') {
          sh '''
            mkdir -p /go/src/github.com/spinnaker-hellow
            go version
          '''
        }
      }
    }

    stage('docker image build') {
      container('docker'){
        stage('build image') {
          script {
            dockerImage = docker.build imageName
          }
        }
      }
    }

    stage('docker image push') {
      container('docker'){
        stage('push image') {
          script {
            docker.withRegistry('', 'docker-samarthya') {
              dockerImage.push("${buildNumber}")
              dockerImage.push('latest')
            }
          }
        }
      }
    }

  }
}

