# EiT-Client

A client providing MEA timestamp data through a websocket connection. The user can specify how often it should receive data, which again defines the resolution of the data.
## Robotserver API
In order to retrieve data, you must open a websocket connection to the server.
### Get Data
```HTTP
ws://<serverip>:<port>/data/<every-ms>/
```

> `port` is defined in config/config.go, *WebSocketPort*. Default is 6480
>
> `every-ms` specifies the interval (in millisec) of received data packages

### Returned Data Format
For *every-ms* milliseconds, the client sends an arary with the number of spikes per time interval.
```
[<ch1>, <ch2>,<ch3>, ... , <ch60>]
```
**Example:**

If *every-ms* is 60 000, <ch1> represents the number of spikes the first second, <ch2> the number of spikes the next second and so on.

### Roadmap
- [x] Establish connection with MEA server
- [x] Process data from the MEA Server
- [x] Create a websocket interface serving the MEA data
- [ ] Add support for stimulation of the MEA server
