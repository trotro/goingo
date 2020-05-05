pipeline {
  agent { docker { image 'golang' } }
  stages {
//    stage('Code checkout') {
//      steps {
//        git 'https://github.com/trotro/goingo.git'
//      }
//    }
    stage('go version') {
      steps {
        sh 'go version'
      }
    }
    stage('Format & vet') {
      steps {
        sh 'go fmt'
        sh 'go vet'
      }
    }
    stage('SonarQube analysis') {
      steps {
        def scannerHome = tool 'SonarQube Scanner 4.3.0.2102';
        withSonarQubeEnv('sonarqube_lab') {
          sh "${scannerHome}/bin/sonar-scanner"
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
