properties ([disableConcurrentBuilds()])
pipeline {
    agent {
        label 'master'
    }
//  triggers { pollSCM('* * * * *') }
    stages {
        stage ('docker build') {
            steps {
                echo "=============build=============="
                
                   sh 'docker build -t simple-go-app:latest .'
                
            }
        }   
          stage ('docker run') {
            steps {
                echo "=============run=============="
                
                   sh 'docker run -d --name simple-go-1 -p 8081:8080 simple-go-app '
                
            }
        }
          stage ('docker test') {
            steps {
                echo "=============test=============="
                script {
                   def curlResult = sh(script: 'curl -I http://127.0.0.1:8081', returnStatus: true)
                    echo "curl Result: ${curlResult}"
                echo "!!!!!!!!!!!!!success!!!!!!!!!!!!!!!!"
                }
            }
        }
 
      stage ('copy my app') {
            steps {
                echo "=============copy my app=============="
                sh 'docker cp simple-go-1:/app/myapp .'
            }
        }     
          stage ('check file') {
            steps {
                echo "=============check file=============="
                sh 'if [ -e "." ]; then echo "Успех"; else echo "Файл не найден"; fi'
            }
        } 
        stage ('copy local file') {
            steps {
                echo "=============copy local file=============="
                sh 'cp myapp /home/slavik'
            }
        } 
        
                         stage ('docker kill') {
            steps {
                echo "=============kill=============="
                sh 'docker kill simple-go-1'
            }
        } 
        stage ('docker delete') {
            steps {
                echo "=============docker delete=============="
                sh 'docker rm simple-go-1'
            }
        } 
    }


 
}
