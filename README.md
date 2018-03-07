# EiT-Client

## Robotserver API
### Get Data
```HTTP
get localhost:<port>/data/<every-ms>
```
> `every-ms` specifies the interval (in millisec) of received data packages

### Returned Data Format
The server returns an array with 60 count numbers of spikes per time interval (one value per multielectrode channel).
```
[<ch1>, <ch2>,<ch3>, ... , <ch60>]
```
Example:
```
[0, 1, 3, 0, 0, ... , 4, 0, 1]
```

