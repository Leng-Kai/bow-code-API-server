pipeline {
    
    agent any
    
    environment {
        GITHUB_REPO_URL = "https://github.com/Leng-Kai/bow-code-API-server"
    }
    
    stages {
        
    	stage('Clone') {
            steps {
                echo 'Cloning..'
                sh "rm -rf ./*"
                sh "git clone $GITHUB_REPO_URL"
                sh "ls"
            }
        }
    
        stage('Build') {
            steps {
                echo 'Building..'
                dir("bow-code-API-server") {
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
    }
    
    post {
        
        always {
            echo "Finished"
        }
        
        success {
            discordSend(
                description: "success",
                link: currentBuild.absoluteUrl,
                result: currentBuild.currentResult,
                successful: currentBuild.resultIsBetterOrEqualTo('SUCCESS'),
                title: currentBuild.fullDisplayName,
                webhookURL: "${env.DISCORD_WEBHOOK_URL}"
            )
        }
        
        failure {
            discordSend(
                description: "failed",
                link: currentBuild.absoluteUrl,
                result: currentBuild.currentResult,
                successful: currentBuild.resultIsBetterOrEqualTo('SUCCESS'),
                title: currentBuild.fullDisplayName,
                webhookURL: "${env.DISCORD_WEBHOOK_URL}"
            )
        }
    }
}

