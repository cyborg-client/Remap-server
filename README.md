# EiT-Client

## Robotserver API
### Get Data
```HTTP
get localhost:<port>/data/<every-ms>
```
> `every-ms` specifies the interval (in millisec) of received data packages

### Returned Data Format
The server returns an array with 60 elements (one per multielectrode channel).
```
[<ch1>, <ch2>,<ch3>, ... , <ch60>]
```
Example:
```
[0, 1, 1, 0, 0, ... , 1, 0, 0]
```
> A `1`/`0` means that a spike occured, or not, in that time segment.
