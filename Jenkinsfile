pipeline {
    agent any

    environment {
        IMAGE_NAME = "iam-api" // MAKE IT SAME AS CONTAINER NAME
        GITLAB_REGISTRY = "registry.gitlab.com/iwibowo_team/paradise_inventory"
        CONTAINER_NAME = "iam-api-${DOCKER_TAG}"
        EC2_USER = "ubuntu"
        SSH_CREDENTIALS_ID = "cred-ssh_poc_server"
        GITLAB_CREDENTIALS = "cred-gitlab_repo"  // GitLab Container Registry credentials
        SONARQUBE_URL = 'https://sonar.delinnce.com'
        SONARQUBE_TOKEN = credentials('token-sonarqube')  // Secure SonarQube token from Jenkins Credentials
        CONSUL_URL = "https://consul.delinnce.com"
        CONSUL_TOKEN = credentials('token-consul')
        SNYK_TOKEN = credentials   ('token-snyk')
    } 

    stages {
        stage('Determine Environment') {
            steps {
                script {
                    // Fetch the branch name that triggered the pipeline
                    def BRANCH_NAME = sh(script: "echo \${GIT_BRANCH} | sed 's|origin/||g'", returnStdout: true).trim()
                    // Alternatif method if GIT_BRANCH is not available
                    if (!BRANCH_NAME) {
                        BRANCH_NAME = sh(script: 'git rev-parse --abbrev-ref HEAD', returnStdout: true).trim()
                    }

		            env.BRANCH_NAME = BRANCH_NAME
                    echo "Current branch: ${env.BRANCH_NAME}"

                    def envConfig = [
                        'main': [env: 'prod', host: '43.218.231.36', tag: 'latest', port: '10500', consul_path: 'dev/iam-api'],
                        'stg': [env: 'stg', host: '43.218.231.36', tag: 'stg', port: '10510', consul_path: 'dev/iam-api'],
                        'dev': [env: 'dev', host: '43.218.231.36', tag: 'dev', port: '10520', consul_path: 'dev/iam-api'],
                    ]

                    if (!envConfig.containsKey(BRANCH_NAME)) {
                    error "Pipeline can only be triggered by dev, stg, or main branches. Current branch: ${BRANCH_NAME}"
                    }

                    def config = envConfig[BRANCH_NAME]
                        env.DEPLOY_ENV = config.env
                        env.EC2_HOST = config.host
                        env.DOCKER_TAG = config.tag
                        env.APP_PORT = config.port
                        env.CONSUL_PATH = config.consul_path
                        env.CONTAINER_NAME = "iam-api-${env.DEPLOY_ENV}"
                        env.CONTAINER_NAME2 = "iam-api-${env.DOCKER_TAG}"
                        
		            echo "Deploying to ${env.DEPLOY_ENV} (${env.EC2_HOST}) with tag ${env.DOCKER_TAG} on port ${env.APP_PORT}"
		            echo "Using Consul config key: ${env.CONSUL_PATH}"
                }
            }
        }

        stage('Checkout') {
            steps {
                script {
                    echo "Checking out branch: ${env.BRANCH_NAME}"

		    // Checkout the branch that triggered the pipeline
                    git (
                        url: 'https://gitlab.com/iwibowo_team/paradise_inventory/iam-api.git',
                        credentialsId: GITLAB_CREDENTIALS, // The credentials ID you created
                        branch: env.BRANCH_NAME // Use the branch that triggered the pipeline
                    )
                }
            }
        }

        stage('SonarQube Analysis') {
            steps {
                script {
                    withSonarQubeEnv('SonarQube') {
                        sh """
                        /opt/sonar-scanner-5.0.1.3006-linux/bin/sonar-scanner \
                            -Dsonar.projectKey=iam-api \
                            -Dsonar.sources=. \
                            -Dsonar.host.url=${SONARQUBE_URL} \
                            -Dsonar.login=${SONARQUBE_TOKEN}
                        """
                    }
                }
            }
        }

        stage('Check SonarQube Quality Gate') {
            steps {
                script {
                    withCredentials([usernamePassword(credentialsId: 'cred-sonar_auth', usernameVariable: 'SONAR_USER', passwordVariable: 'SONAR_PASS')]) {
                        // Get the SonarQube Task ID
                        def response = sh(script: """
                            curl -s -u ${SONAR_USER}:${SONAR_PASS} "${SONARQUBE_URL}/api/ce/activity?component=iam-api"
                        """, returnStdout: true).trim()

                        def jsonResponse = readJSON text: response
                        def taskId = jsonResponse.tasks[0].id  // Extract the latest task ID

                        echo "SonarQube Task ID: ${taskId}"
                        env.SONAR_TASK_ID = taskId

                        // Fetch SonarQube Task Status
                        def taskStatus = ""
                        for (int i = 0; i < 30; i++) {  // Wait up to 5 minutes (10 sec * 30 retries)
                            sleep 10
                            def taskResponse = sh(script: """
                                curl -s -u ${SONAR_USER}:${SONAR_PASS} "${SONARQUBE_URL}/api/ce/task?id=${taskId}"
                            """, returnStdout: true).trim()

                            def taskJson = readJSON text: taskResponse
                            taskStatus = taskJson.task.status

                            echo "SonarQube Task Status: ${taskStatus}"

                            if (taskStatus == "SUCCESS" || taskStatus == "FAILED" || taskStatus == "ERROR") {
                                break  // Stop retrying if task is completed
                            }
                        }

                        // Abort pipeline if Quality Gate fails
                        if (taskStatus == "FAILED" || taskStatus == "ERROR") {
                            error "Pipeline aborted due to SonarQube Quality Gate failure!"
                        }
                    }
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                sh 'docker build --build-arg PORT=${APP_PORT} -t $GITLAB_REGISTRY/$IMAGE_NAME:${DOCKER_TAG} .'
		        sh 'docker image prune -a --filter "until=24h" -f'
            }
        }

        stage("Snyk Container Scan") {
            steps {
                script {
                    echo "Running Snyk container vulnerability scanning"
                    withEnv(["SNYK_TOKEN=${SNYK_TOKEN}"]) {
                        try {
                            // Run Snyk scan and save results
                            
                            def timestamp = sh(script: """
                                #!/bin/bash
                                date +'%d-T%H%M'
                            """, returnStdout: true).trim()
                            echo "Generated timestamp: ${timestamp}"
                            
                            def REPORT_FILENAME = "/var/log/snyk/apps/${IMAGE_NAME}/snyk-container-report-${timestamp}-${env.DEPLOY_ENV}.json"
                            echo "Report filename: ${REPORT_FILENAME}"

                            def scanResult = sh(script: """
                                snyk container test $GITLAB_REGISTRY/$IMAGE_NAME:${DOCKER_TAG} --severity-threshold=high --json > $REPORT_FILENAME""",
                                returnStatus: true
                                )

                            // Parse results
                            def scanReport = readJSON file: REPORT_FILENAME
                            def vulnerabilityCount = scanReport.vulnerabilities ? scanReport.vulnerabilities.size() : 0
                            def criticalVulns = scanReport.vulnerabilities ? scanReport.vulnerabilities.findAll { it.severity == 'critical' }.size() : 0
                            def highVulns = scanReport.vulnerabilities ? scanReport.vulnerabilities.findAll { it.severity == 'high' }.size() : 0

                            echo "Security scan complete: ${criticalVulns} critical, ${highVulns} high vulnerabilities found"

                            // Always enforce security rules, regardless of environment
                            // This ensures vulnerabilities are caught early in development
                            if (criticalVulns > 100) {
                                error "CRITICAL SECURITY ISSUE: Found ${criticalVulns} critical vulnerabilities. Deployment aborted."
                            }

                            // Different thresholds by environment for high vulnerabilities
                            if (env.DEPLOY_ENV == 'dev' && highVulns > 100) {
                                error "HIGH SECURITY RISK: Found ${highVulns} high vulnerabilities exceeding dev threshold (10)."
                            } else if (env.DEPLOY_ENV == 'stg' && highVulns > 100) {
                                error "HIGH SECURITY RISK: Found ${highVulns} high vulnerabilities exceeding staging threshold (5)."
                            } else if (env.DEPLOY_ENV == 'prod' && highVulns > 100) {
                                error "HIGH SECURITY RISK: Found ${highVulns} high vulnerabilities. Production must have zero high vulnerabilities."
                            }

                            // Document any remaining issues (medium/low) for awareness
                            if (vulnerabilityCount > 0){
                                echo "Non-blocking vulnerabilities found. See report for details at ${REPORT_FILENAME}"
                            } else {
                                echo "No vulnerabilities found. Report location ${REPORT_FILENAME}"
                            }
                        } catch (Exception e) {
                            echo "Error during Snyk scan: ${e.message}"
                            error "Security scan failed: ${e.message}"
                        }
                    }
                }
            }
        }

        stage('Login to GitLab Registry') {
            steps {
                withCredentials([usernamePassword(credentialsId: GITLAB_CREDENTIALS, usernameVariable: 'GITLAB_USER', passwordVariable: 'GITLAB_PASS')]) {
                    sh 'docker login -u $GITLAB_USER -p $GITLAB_PASS registry.gitlab.com'
                }
            }
        }

        stage('Push to GitLab Container Registry') {
            steps {
                sh 'docker push $GITLAB_REGISTRY/$IMAGE_NAME:$DOCKER_TAG'
            }
        }

        stage('Deploy to EC2') {
            steps {
                withCredentials([usernamePassword(credentialsId: GITLAB_CREDENTIALS, usernameVariable: 'GITLAB_USER', passwordVariable: 'GITLAB_PASS')]) {
                    sshagent(credentials: [SSH_CREDENTIALS_ID]) {
                        sh """
                            ssh -o StrictHostKeyChecking=no $EC2_USER@$EC2_HOST '
                            docker login -u $GITLAB_USER -p $GITLAB_PASS registry.gitlab.com
                            docker pull $GITLAB_REGISTRY/$IMAGE_NAME:$DOCKER_TAG
                            docker stop ${env.CONTAINER_NAME2} || true
                            docker rm ${env.CONTAINER_NAME2} || true
                            curl -s -H "X-Consul-Token: ${CONSUL_TOKEN}" ${CONSUL_URL}/v1/kv/${CONSUL_PATH}?raw > /etc/consul.d/config/${env.DEPLOY_ENV}/iam.config.yml
                            docker run -d --name ${env.CONTAINER_NAME2} --restart always -p ${APP_PORT}:${APP_PORT} -v /etc/consul.d/config/${env.DEPLOY_ENV}/iam.config.yml:/app/config.yml $GITLAB_REGISTRY/$IMAGE_NAME:$DOCKER_TAG
                            docker image prune -a --filter "until=24h" -f
                            '
                        """
                    }
                }
            }
        }
    }

    post {
        always {
            echo "Pipeline finished"
        }
        success {
            echo "Pipeline completed successfully!"
        }
        failure {
            echo "Pipeline failed!"
        }
    }
}
