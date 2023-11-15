properties ([disableConcurrentBuilds()])
pipeline {
    agent {
        label 'master'
    }
 // triggers { pollSCM('* * * * *') }
    stages {

         stage ('check and delete old build_app1') {
            steps {
                echo "=============check and delete old build_app2=============="
                sh 'docker ps -q -f name=build_app2 && docker kill build_app1 && docker rm build_app2 || echo "Контейнер build_app2 не найден или не запущен."'
            }
        } 
        stage ('docker build') {
            steps {
                echo "=============build=============="
                
                   sh 'docker build -t simple-go-app2:latest .'
                
            }
        }   
          stage ('docker run') {
            steps {
                echo "=============run=============="
                
                   sh 'docker run -d --name simple-go-2 -p 8081:8080 simple-go-app2 '
                
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
                sh 'docker cp simple-go-2:/app/myapp .'
            }
        }     
  /*        stage ('check file') {
            steps {
                echo "=============check file=============="
                sh 'if [ -e "." ]; then echo "Успех"; else echo "Файл не найден"; fi'
            }
        } 
        stage ('copy local file') {
            steps {
                echo "=============copy local file=============="
                sh 'cp myapp /home/slavik/'
            }
        } 
        */
                         stage ('docker kill') {
            steps {
                echo "=============kill=============="
                sh 'docker kill simple-go-2'
            }
        } 
        stage ('docker delete') {
            steps {
                echo "=============docker delete=============="
                sh 'docker rm simple-go-2'
            }
        } 

          stage ('docker build app') {
            steps {
                echo "=============docker build app=============="
                sh 'docker build -t build-app2:latest -f Dockerfile_build .'
            }
        } 
        
        stage ('copy arhiv') {
            steps {
                echo "=============copy arhiv=============="
                sh '[ -e "/home/slavik/file/build-app-archive2.tar" ] && echo "Файл build-app-archive2.tar уже существует." || docker save -o /home/slavik/file/build-app-archive2.tar build-app2:latest'
            }
        } 

            stage ('docker run app') {
            steps {
                echo "=============docker run app=============="
                sh 'docker run -d --name build_app2 -p 8086:8080 build-app:latest'
            }
        } 
    }


 
}
