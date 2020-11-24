Goapiscraper implements the RESTful endpoint API, which makes concurrent calls to the following websites:

https://www.result.si/projekti/ 
https://www.result.si/o-nas/
https://www.result.si/kariera/ 
https://www.result.si/blog/

The input data for the endpoint is integer, which represents the number of threads/Goroutins to the above web pages (min 1 represents all consecutive calls, max 4 represents all concurrent calls).

It then extracts a short title text from each page and save this text in a common global structure.

The service then displays the number of successful calls, the number of failed calls and the saved titles from all web pages.

Requests can be sent in the form of a http GET to localhost:8080/.

A wecome screen can be found at localhost:8080/
