podTemplate(yaml: readFile(file: 'node.yaml')) {
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

