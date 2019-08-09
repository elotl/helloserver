pipeline {
  agent {
    kubernetes {
      defaultContainer 'jnlp'
      yaml """
apiVersion: v1
kind: Pod
metadata:
  annotations:
    kubernetes.io/target-runtime: kiyot
spec:
  containers:
  - name: golang
    image: golang:1.12
    command:
    - cat
    tty: true
"""
    }
  }
  stages {
    stage('Build helloserver') {
      steps {
        container('golang') {
          sh 'mkdir -p /go/src/github.com/elotl; ln -s `pwd` /go/src/github.com/elotl/helloserver; cd /go/src/github.com/elotl/helloserver && sleep 300 && go build'
        }
      }
    }
  }
}
