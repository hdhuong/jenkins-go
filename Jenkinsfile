pipeline {
    agent any

    environment {
        REGISTRY = "docker.io"
        REGISTRY_CREDENTIALS = "docker-hub"
        IMAGE_NAME = "hdhuong1/go-app"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build Image') {
            steps {
                script {
                    def tag = "${env.BRANCH_NAME}-${env.BUILD_NUMBER}"
                    sh "/Applications/OrbStack.app/Contents/MacOS/bin/docker build -t $IMAGE_NAME:$tag ."
                    env.IMAGE_TAG = tag
                }
            }
        }

        stage('Push to Registry') {
            steps {
                withDockerRegistry([credentialsId: REGISTRY_CREDENTIALS, url: "https://$REGISTRY"]) {
                    sh "/Applications/OrbStack.app/Contents/MacOS/bin/docker push $IMAGE_NAME:${IMAGE_TAG}"
                }
            }
        }
    }
}