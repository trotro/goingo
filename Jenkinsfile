pipeline {
  agent none
  stages {
    stage('SonarQube analysis') {
      agent any
      steps {
        script {
          scannerHome = tool 'sonarqubescanner';
        }
        withSonarQubeEnv('sonarqube_lab') {
          sh "${scannerHome}/bin/sonar-scanner"
        }
      }
    }
//    stage('Quality Gate') {
//      agent any
//      steps {
//        timeout(time: 5, unit: 'MINUTES') {
          // Parameter indicates whether to set pipeline to UNSTABLE if Quality Gate fails
          // true = set pipeline to UNSTABLE, false = don't
//          waitForQualityGate abortPipeline: true
//        }
//      }
//    }
    stage('go test & benchmarks') {
      agent any
//      agent { docker { image 'golang' } }
//      steps {
//        sh 'go version'
//        sh 'go test -v --bench . --benchmem --cover'
//      }
      steps {
        docker.image('golang').inside {
          git 'https://github.com/trotro/goingo.git'
          sh 'go version'
          sh 'go test -v --bench . --benchmem --cover'
        }
      }
    }
  }
}
