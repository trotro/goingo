pipeline {
  agent { docker { image 'golang' } }
  stages {
    stage('link repo into $GOPATH') {
      steps {
        sh "echo ${env.WORKSPACE}"
        sh "echo ${env.JOB_NAME}"
        sh "echo ${env.BRANCH_NAME}"
        //sh "ln -s ${env.WORKSPACE}/${env.JOB_NAME}_${env.BRANCH_NAME} /go/src/${env.JOB_NAME}"
        sh "ls -l $GOPATH/"
      }
    }
    stage('go version') {
      steps {
        sh 'go version'
      }
    }
    stage('Format & vet') {
      steps {
        //sh "cd $GOPATH/${env.JOB_NAME}"
        sh 'go fmt .'
        sh 'go vet .'
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
