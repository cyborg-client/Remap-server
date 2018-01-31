## HTTP Protocol
### Messages
There's two types of messages: _request_ and _response_.

#### Requests
Requests consist of the following elements:
* An HTTP method, usally a verb like `GET`, `GET` or a noun like `OPTIONS` or `HEAD` that defines the operation the client wants to perform. Typically fetch a resource using `GET`.
* The path of the resource to fetch; the URL of the resource stripped from elements that are obvious from the context, for example without the protocol (http://), the domain (dev.google.com), or the TCP port (e.g: 80).
* The version of the HTTP protocol.
* [Optional] Headers that convey additional information for the servers. Or a body, for some methods like `POST`, similar to those in responses which contain the resource sent.

Example request:
```HTTP
GET /RANDOM/PATH HTTP/1.1
Host: dev.google.com
```

#### Responses
Requests consist of the following elements:
* The version of the HTTP protocol they follow.
* A _status code_, indicating if the request has been successful, or not, and why.
* A _status message_, a non-authoritative short description of the status code.
* HTTP header, like those for requests.
* [Optional] A body containing the fetched resource.

Example response:
```HTTP
HTTP/1.1 200 OK
Date: Sat, 31 jan 2018 15:00:00 GMT
Server: Apache
Last-modified: Tue, 01 Dec 2017 20:00:00 GMT
ETag: "12328372"
Accept-Ranges: bytes
Content-Length: 29700
Content-Type: text/html