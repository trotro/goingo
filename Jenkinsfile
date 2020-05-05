pipeline {
  agent { docker { image 'golang' } }
  stages {
    stage('go version') {
      steps {
        sh 'go version'
      }
    }
    stage('SonarQube analysis') {
      steps {
        script {
          scannerHome = tool 'sonarqube_local_scanner';
        }
        withSonarQubeEnv('SonarQube') {
          sh "${scannerHome}/bin/sonar-scanner -Dsonar.host.url=http://localhost:9000"
          //sh "/home/jenkins/tools/hudson.plugins.sonar.SonarRunnerInstallation/sonarqubescanner/bin/sonar-scanner -Dsonar.host.url=http://192.168.0.14:9000 -Dsonar.projectName=meanstackapp -Dsonar.projectVersion=1.0 -Dsonar.projectKey=meanstack:app -Dsonar.sources=. -Dsonar.projectBaseDir=/home/jenkins/workspace/sonarqube_test_pipeline"
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
