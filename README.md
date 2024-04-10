# IM-cli-chat server

This is a simple chat server implemented in Go, using WebSockets for real-time communication. It allows multiple clients to connect and exchange messages. 

<img width="1438" alt="Screenshot 2024-04-10 at 12 52 13" src="https://github.com/ivermoka/IM-cli-chat-server/assets/119415554/7a8b5e7b-705d-47f9-9071-db1b1ec13821">


This project is primarily made for IM at Elvebakken Upper Secondary School, and hosted at a VM on our server. The downside of this, is that it can only be accessed to people on the same network as the VM, or by using a VPN. You could open a port to the public, but for security reasons, I chose to have it run locally on our network. 

This is the server side part of the IM-cli-chat project. The client side repo is [here](https://github.com/ivermoka/IM-cli-chat).

### Running it on your own

#### Prerequisites
* Have [Docker](https://www.docker.com/) installed on your system. **(If you choose to deploy using docker, which is recommended)**
* Have [go](https://go.dev/) version **1.22.1** (may work on different versions) installed on your system. **(If you choose to deploy using go)**
* Have [git](https://git-scm.com/) installed on your system.
* Use a terminal with a Unix shell concept that derives from Bourne shell (Linux, MacOS). If you use Windows, you may need to use WSL in order to follow the guide.

  
#### Deploying 


* Clone this repo (```git clone https://github.com/ivermoka/IM-cli-chat-server```)

##### Docker version (recommended)
* Build an image (```sudo docker build -t im-cli-chat-server .```)
* Run using your recently built image (```sudo docker run -d -p 8080:8080 im-cli-chat-server```)

##### Go version (not recommended)
* Make sure to be in the correct directory
* **(Not recommended)** You can just run the go program like normal, but then the program will stop once you close your terminal window (```go run .``` or ```go build . ./your_executable_name```)
* **(Recommended)** Use nohup to run the program detatched (```nohup go run .```). If you need to check if the process is running, you can use ```ps aux | grep program_name```. The numbers next to your user is the process ID. If you need to kill process, use ```kill your_process_id(PID)``` (you find the ID with the previous command). PS: after the ```nohup``` command has been written, be sure to exit the terminal tab in order to not kill process accidentally.

The server will start running on port 8080 by default.


### Usage

* Clients can connect to the server using a WebSocket connection.
* Once connected, clients can send messages to the server, which will be broadcasted to all connected clients.
* I have also made a client side program: [IM-cli-chat](https://github.com/ivermoka/IM-cli-chat/), which can be used to connect to this program. 


### Features

* WebSocket-based communication
* Real-time messaging
* Seeing amount of peers connected (not usable when running a detached program)
* CLI program (used in terminal)



### Dependencies
* [gorilla/websocket](https://github.com/gorilla/websocket): A WebSocket implementation for Go.

### Contributing

Contributions are welcome! If you'd like to contribute to this project, feel free to open an issue or submit a pull request.
