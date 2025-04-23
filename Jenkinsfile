pipeline {
  agent any

  stages {
    stage('Build Docker Image') {
      steps {
        script {
          dockerImage = docker.build("jcladerabe/miapp:latest")
        }
      }
    }
    stage('Push Docker Image') {
      steps {
        withDockerRegistry([credentialsId: 'dockerhub-cred-id', url: '']) {
          dockerImage.push()
        }
      }
    }
  }
}
