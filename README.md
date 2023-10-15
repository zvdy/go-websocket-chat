# 📢 Chat App

This is a simple chat app built with Go and WebSocket. It allows multiple users to join a chat room and send messages to each other in real time.

## ⚙️ Installation

To install the chat app, you need to have Go installed on your computer. You can download and install Go from the official website: https://golang.org/dl/

Once you have Go installed, you can download the chat app source code by running the following command in your terminal:

```
git clone https://github.com/zvdy/go-websocket-chat.git
```


This will download the chat app source code to your current directory.

## 🧑‍💻 Usage

To start the chat app server, navigate to the `go-websocket-chat` directory in your terminal and run the following command:

```
go run main.go
```
> You can also build the chat app server by using the Dockerfile in the root directory of the project. To do this, run the following command in your terminal: `docker build -t go-websocket-chat .` This will build the chat app server and tag it as `go-websocket-chat`. You can then run the chat app server by running the following command: `docker run -p 8080:8080 go-websocket-chat`. [Here](https://hub.docker.com/r/zvdy/go-websocket-chat/tags) is the latest docker tag.

This will start the server on port 8080. You can access the chat app by going to `http://localhost:8080/chat` in your web browser.

To join a chat room, append the `room` query parameter to the URL. For example, to join a chat room with ID `test`, go to `http://localhost:8080/chat?room=test`.

![image](/images/sample.png)

To send a message to the chat room, type your message in the input field at the bottom of the chat window and press Enter. Your message will be broadcast to all users in the chat room.

You can also send messages programmatically using the WebSocket API. To do this, you need to open a WebSocket connection to the chat app server and send JSON messages with the following format:

```json
{
  "username": "Cristian",
  "text": "Hello, world!"
}
```
Replace Alice with your username and Hello, world! with your message text.

You can use the following curl command to send a message to the chat room:

```
curl -X POST -H "Content-Type: application/json" -d '{"username":"your_username","text":"your_message"}' http://localhost:8080/message?room=your_room_id
```


---


To open a WebSocket connection, send an HTTP GET request to the /ws endpoint with the room query parameter. For example, to join a chat room with ID test, run the following command in your terminal:

```
curl -i -N -H "Connection: Upgrade" -H "Upgrade: websocket" -H "Sec-WebSocket-Version: 13" -H "Sec-WebSocket-Key: SGVsbG8sIHdvcmxkIQ==" http://localhost:8080/ws?room=test
```

Replace the test with the ID of the chat room you want to join.

The output should look like this:

```
HTTP/1.1 101 Switching Protocols
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Accept: s3pPLMBiTxaQ9kYGzzhZRbK+xOo=
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
