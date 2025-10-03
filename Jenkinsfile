pipeline {
    agent any

    environment {
        REGISTRY = "docker.io"
        REGISTRY_CREDENTIALS = "docker-hub" // Jenkins credentials ID (Docker Hub account)
        IMAGE_NAME = "hdhuong1/go-app"
        
        // Server credentials (SSH private key)
        STAGING_SERVER = "staging-server"
        PROD_SERVER    = "prod-server"

        // Deploy paths
        STAGING_PATH = "/home/ubuntu/staging-app"
        PROD_PATH    = "/home/ubuntu/prod-app"
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
                    sh "export PATH=$PATH:/usr/local/bin && docker build -t $IMAGE_NAME:$tag ."
                    env.IMAGE_TAG = tag
                }
            }
        }

        stage('Push to Registry') {
            steps {
                withCredentials([usernamePassword(credentialsId: "${REGISTRY_CREDENTIALS}", usernameVariable: 'USER', passwordVariable: 'PASS')]) {
                    sh "echo $PASS | docker login -u $USER --password-stdin $REGISTRY"
                    sh "docker push $IMAGE_NAME:${IMAGE_TAG}"
                }
            }
        }

        // stage('Deploy Staging') {
        //     when {
        //         branch 'develop'
        //     }
        //     steps {
        //         sshagent(credentials: ["${STAGING_SERVER}"]) {
        //             sh """
        //             ssh -o StrictHostKeyChecking=no ubuntu@staging-server-ip '
        //                 cd ${STAGING_PATH} &&
        //                 docker pull $IMAGE_NAME:${IMAGE_TAG} &&
        //                 docker compose down &&
        //                 IMAGE_TAG=${IMAGE_TAG} docker compose up -d --build
        //             '
        //             """
        //         }
        //     }
        // }

        // stage('Deploy Production') {
        //     when {
        //         branch 'main'
        //     }
        //     steps {
        //         sshagent(credentials: ["${PROD_SERVER}"]) {
        //             sh """
        //             ssh -o StrictHostKeyChecking=no ubuntu@prod-server-ip '
        //                 cd ${PROD_PATH} &&
        //                 docker pull $IMAGE_NAME:${IMAGE_TAG} &&
        //                 docker compose down &&
        //                 IMAGE_TAG=${IMAGE_TAG} docker compose up -d --build
        //             '
        //             """
        //         }
        //     }
        // }
    }
}