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
          //sh "/home/jenkins/tools/hudson.plugins.sonar.SonarRunnerInstallation/sonarqubescanner/bin/sonar-scanner -Dsonar.host.url=http://192.168.0.14:9000 -Dsonar.projectName=meanstackapp -Dsonar.projectVersion=1.0 -Dsonar.projectKey=meanstack:app -Dsonar.sources=. -Dsonar.projectBaseDir=/home/jenkins/workspace/sonarqube_test_pipeline"
        }
      }
    }
    stage('Quality Gate') {
      agent any
      steps {
        timeout(time: 1, unit: 'HOURS') {
          // Parameter indicates whether to set pipeline to UNSTABLE if Quality Gate fails
          // true = set pipeline to UNSTABLE, false = don't
          waitForQualityGate abortPipeline: true
        }
      }
    }
    stage('go test & benchmarks') {
      agent { docker { image 'golang' } }
      steps {
        sh 'go version'
        sh 'go test -v --bench . --benchmem --cover'
      }
    }
  }
}
