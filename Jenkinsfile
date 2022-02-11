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
      container('kaniko'){
        stage('build image') {
          sh '''
          executor --help
          ls -als
          '''
          sh "/kaniko/executor --context `pwd`  --dockerfile `pwd`/Dockerfile --destination=${imageName}:${buildNumber} --destination=${imageName}:latest"
        }
      }
    }
  }
}

