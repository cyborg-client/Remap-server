# Remap Server

A client providing MEA buckets of spike data, where each bucket contains the number of spikes since the last transmission. This is served through a websocket connection. The user can specify how often it should receive data, which again defines the resolution of the data. The client is written in golang.

## Usage
### Requirements
golang.org/x/net/websocket and guithub.com/satori/go.uuid are needed to run the application.
```bash
go get golang.org/x/net/websocket
go get github.com/satori/go.uuid
```

You also need to have an MEA server running. If you do not have one, a simulator can be found at https://github.com/cyborg-client/offline-server

### Installation
In order to install the client, simply use go get
```bash
go get github.com/cyborg-client/client
```

### Running
Assuming proper gopath setup, start the application with
```bash
go run $GOPATH/src/github.com/cyborg-client/client/main.go
```

### Config

Configuration is done by editing config/config.go. The entries should be self explanatory.

## Websocket API
In order to retrieve data, you must open a websocket connection to the server.
### Get Data
```HTTP
ws://<serverip>:<websocket-port>/data/<every-ms>/
```

> `websocket-port` is defined in config/config.go, *WebSocketPort*. Default is 6480
>
> `every-ms` specifies the interval (in millisec) of received data packages

### Returned Data Format
For *every-ms* milliseconds, the client sends an array with the number of spikes per time interval.
```
[<ch1>, <ch2>,<ch3>, ... , <ch60>]
```
**Example:**

If *every-ms* is 60 000, <ch1> represents the number of spikes in channel 1 the last 60 seconds, <ch2> the number of spikes in channel 2 the last 60 seconds and so on.

### Send stimulation

You can stimulate a given channel by sending a POST request.
```HTTP
POST /stimulate
Host: <serverip>

{
  "frequency": <frequency>,
  "duration": <duration>,
  "channel": <channel>
}
```

### Example websocket Client

If you wish to test if the client is running properly, you can use the following javascript code snippet.

```javascript
(() => {

let ws = new WebSocket('ws://129.241.187.141:6780/data/1000/');
ws.onmessage = (msg) => {
    console.log(JSON.parse(msg.data));
};

ws.onclose = () => {
    console.log("Server closed connection.")
};

// Closes the connection after 5 seconds
ws.onopen = () => {
    setTimeout(() => {
        ws.close();
        console.log('We closed connection.');
    }, 5000);
};
})();
```
Please note that this script **must** be run from a website not using https, as the client does not support SSL.

### Roadmap
- [x] Establish connection with MEA server
- [x] Process data from the MEA Server
- [x] Create a websocket interface serving the MEA data
- [x] Add support for stimulation of the MEA server
