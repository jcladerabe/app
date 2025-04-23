pipeline {
  agent any

  environment {
    DOCKER_IMAGE = "jcladerabe/miapp:latest"
  }

  stages {
    stage('Clonar repo') {
      steps {
        git branch: 'main', url: 'git@github.com:jcladerabe/app.git'
      }
    }

    stage('Construir imagen Docker') {
      steps {
        script {
          dockerImage = docker.build("${DOCKER_IMAGE}")
        }
      }
    }

    stage('Subir imagen Docker') {
      steps {
        script {
          docker.withRegistry('', 'dockerhub-cred-id') {
            dockerImage.push()
          }
        }
      }
    }

    stage('Desplegar en Kubernetes') {
      steps {
        sh 'kubectl apply -f deployment.yaml'
      }
    }
  }
}
