pipeline{
  agent {
    kubernetes {
      yamlFile 'node.yaml'
    }
  }
  environment {
    imageName = "bhanuni/spinnaker-hellow"
    buildNumber = "${env.BUILD_NUMBER}"
  }
  stages {
    stage('Get a Golang project') {
      steps {
        git url: 'https://github.com/samarthya/spinnaker-hello.git', branch: 'main', credentialsId: 'github-samarthya'
        container('kaniko') {
          sh '''
          /kaniko/executor --context `pwd`  --dockerfile `pwd`/Dockerfile --destination=${imageName}:${buildNumber} --destination=${imageName}:latest
          '''
        }
      }
    }
  }
}