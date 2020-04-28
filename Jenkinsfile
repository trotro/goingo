pipeline {
  agent { docker { image 'golang' } }
  stages {
    stage('Code checkout') {
      steps {
        git url: 'https://github.com/trotro/goingo.git'
      }
    }
    stage('go version') {
      steps {
        sh 'go version'
      }
    }
    stage('Format & vet') {
      steps {
        sh 'go fmt -w'
        sh 'go vet'
      }
    }
    stage('SonarQube analysis') {
      steps {
        withSonarQubeEnv('sonarqube_lab') {
          
        }
      }
    }
    stage('Test & benchmarks') {
      steps {
        sh 'go test -v --bench . --benchmem --cover'
      }
    }
  }
}
