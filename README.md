# line_robot

#### line_robot implements 3 APIs to receive, reply and send messages to Line's official account.
####  Use to CINNOX interview test

## About Projects

### Spec:
* commit #1 setup project
* commit #2 Makefile or a script for local setup and run MongoDB docker (version: 4.4)
* commit #3 setup necessary config of LINE, MongoDB
  - Line official account message integration (use go line sdk),
  - Create a test line dev official account
* commit #4 Create a Go package connect to mongoDB, create a model/DTO to save/query user message to MongoDB
* commit #5 Create a Gin API
  - receive message from line webhook, save the user info and message in MongoDB
* commit #6 Create a API send message back to line
* commit #7 Create a API query message list of the user from MongoDB 
* provide a demo video or steps of test (or postman or ...) and github repo link

## Getting Started

This is an instruction on how we set up line_robot locally. Running follow these simple example steps.

### Prerequisites

1. [Download and install Go](https://go.dev/doc/install)

2. Install Docker

      #### **For Mac:**
      Install Homebrew
      ```
      /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
      ```
      Install Docker
      ```
      brew cask install docker
      ```
      #### **For Amazon Linux 2:**
      Apply pending updates using the yum command
      ```
      sudo yum update
      ```
      Search for Docker package
      ```
      sudo yum search docker
      ```

3. Clone this repository

### Run

1. Open Docker

2. Go into this repository directory

3. Run MongoDB docker image
      ```
      sh run_mongo.sh 
      ```
      
4. Open 8080 port (if you just run with your own PC, you can install ngrok and run below)
      ```
      ngrok http 8080
      ```
5. Get your forwarding URL with 8080 port
      
6. Upload Webhook URL(forwarding URL) to your Line official account (don't forget to verify it works or not by "Verify" button offered by Line official)

7. Save your "Channel secret" and "Channel access token" from your Line official account into config.yaml

8. Run server in line_robot repository directory
      ```
      go run main.go
      ```
      
9. Send a message to your Line robot!

## Result

<img width="1031" alt="line_robot_demo_screenshot" src="https://user-images.githubusercontent.com/23217065/197409834-58250cbf-2ce0-4c5f-81a4-69236d0893b9.png">

## Contact

Victor Chen (Yoea) - cek2llm@gmail.com

