properties ([disableConcurrentBuilds()])
pipeline {
    agent {
        label 'master'
    }
 // triggers { pollSCM('* * * * *') }
    stages {

         stage ('check and stop old build_app1') {
            steps {
                echo "=============check and delete old build_app1=============="
                sh 'docker ps -q -f name=build_app1 && docker kill build_app1 || echo "Контейнер build_app1 не найден или не запущен."'
            }
        } 

           stage ('check and delete old container build_app1') {
            steps {
                echo "=============check and delete old container build_app1=============="
                sh 'docker ps -a -q -f name=build_app1 && docker rm build_app1 || echo "Контейнер build_app1 не найден или удален."'
            }
        } 
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
                sh 'docker kill simple-go-1'
            }
        } 
        stage ('docker delete') {
            steps {
                echo "=============docker delete=============="
                sh 'docker rm simple-go-1'
            }
        } 

          stage ('docker build app') {
            steps {
                echo "=============docker build app=============="
                sh 'docker build -t build-app:latest -f Dockerfile_build .'
            }
        } 
        
        stage ('copy arhiv') {
            steps {
                echo "=============copy arhiv=============="
                sh '[ -e "/home/slavik/file/build-app-archive.tar" ] && echo "Файл build-app-archive.tar уже существует." || docker save -o /home/slavik/file/build-app-archive.tar build-app:latest'
            }
        } 

            stage ('docker run app') {
            steps {
                echo "=============docker run app=============="
                sh 'docker run -d --name build_app1 -p 8085:8080 build-app:latest'
            }
        } 
    }


 
}
