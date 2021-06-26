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
                sh "ls"
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
}
