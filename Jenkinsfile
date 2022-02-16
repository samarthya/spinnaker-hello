podTemplate(yaml: '''
    apiVersion: v1
    kind: Pod
    spec:
      volumes:
      - name: kaniko-secret
        secret:
          secretName: regcred
          items:
          - key: .dockerconfigjson
            path: config.json
      containers:
      - name: helm
        image: alpine/helm:latest
        command:
        - cat
        tty: true
      - name: golang
        image: golang:latest
        command:
        - cat
        tty: true
      - name: kaniko
        image: gcr.io/kaniko-project/executor:debug
        command:
        - cat
        tty: true
        volumeMounts:
          - name: kaniko-secret
            mountPath: /kaniko/.docker
''') {
  node(POD_LABEL) {
    // Define the image name
    def imageName = "bhanuni/spinnaker-hellow"
    def buildNumber = env.BUILD_NUMBER
    def branchName = 'main'
    def chartValue = 'hellow/values.yaml'


    stage('Get a Golang project') {
      git url: 'https://github.com/samarthya/spinnaker-hello.git', branch: branchName, credentialsId: 'github-samarthya'
      container('golang'){
        stage('Test the project') {
          sh '''
            go test
          '''
        }
        stage('Build a Go project') {
          sh '''
            echo hello from $POD_CONTAINER
            mkdir -p /go/src/github.com/spinnaker-hellow
            go build -o bin/spinnaker-hellow .
          '''
        }
        stage('New Values') {
          script {
            def data = readYaml file: chartValue
            data.image.tag = buildNumber
            sh "echo $data"
            sh "rm $chartValue"
            writeYaml file: chartValue, data: data
          }
        }
      }
    }

    if (currentBuild.currentResult == 'SUCCESS') {
      stage('New Chart') {
        container('helm') {
          sh """
          helm version
          helm package hellow
          """
        }
      }
      
      stage('docker image build') {
        container('kaniko'){
          stage('build image') {
            if (branchName == 'main') {
              sh '''
              echo Hello from $POD_CONTAINER
              '''
              sh "/kaniko/executor --context `pwd`  --dockerfile `pwd`/Dockerfile --destination=${imageName}:${buildNumber} --destination=${imageName}:latest"
            }
          }
        }
      }
    }
  }
}

