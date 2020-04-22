pipeline {
  agent { docker { image 'golang' } }
  stages {
    stage('version') {
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
    stage('Test & benchmarks') {
      steps {
        sh 'go test -v --bench . --benchmem --cover'
      }
    }
  }
}
