pipeline {
    
    agent any
    
    environment {
        GITHUB_REPO_URL = "https://github.com/Leng-Kai/bow-code-API-server"
        DEPLOY_URL = "http://localhost:8089/"
    }
    
    stages {
        
    	stage('Clone') {
            steps {
                echo 'Cloning..'
                sh "rm -rf ./bow-code-API-server"
                sh "git clone $GITHUB_REPO_URL"
                sh "ls"
            }
        }
    
        stage('Build') {
            steps {
                echo 'Building..'
                dir("bow-code-API-server") {
                    sh "git checkout dev"
                    sh "docker-compose up --force-recreate --build -d"
                    sh "docker image prune -f"
                }
            }
        }
            
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
            
        stage('Deploy') {
            steps {
                echo 'Deploying..'
            }
        }

        stage('Healthy Check') {
            steps {
                sleep(time:3)
                sh "curl $DEPLOY_URL"
            }
        }
    }
    
    post {
        
        always {
            echo "Finished"
        }
        
        success {
            discordSend(
                description: "dev - success",
                link: currentBuild.absoluteUrl,
                result: currentBuild.currentResult,
                successful: currentBuild.resultIsBetterOrEqualTo('SUCCESS'),
                title: currentBuild.fullDisplayName,
                webhookURL: "${env.DISCORD_WEBHOOK_URL}"
            )
        }
        
        failure {
            discordSend(
                description: "dev - failed",
                link: currentBuild.absoluteUrl,
                result: currentBuild.currentResult,
                successful: currentBuild.resultIsBetterOrEqualTo('SUCCESS'),
                title: currentBuild.fullDisplayName,
                webhookURL: "${env.DISCORD_WEBHOOK_URL}"
            )
        }
    }
}

