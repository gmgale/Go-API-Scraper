![Go logo](img.png)


# GoApiScraper
### A web-scraping REST API that makes use of Go's concurrency functionality for parralel GET requests.
___

## API
The input data for the endpoint is integer, which represents the number of threads/Goroutins to web pages to be scraped.
It then extracts a short title text from each page and save this text in a common global structure.

Requests can be sent in the form of a http GET to the endpoint **/webcall/{integer}**

Command line flags can be used:
* To vary the API port ("8081" default). Example: -port 1234.
* To set the host ("db" default). This will allow a connection to the Docker-PostgreSQL. Use "localhost" if running PostgreSQL on local machine outside of Docker.

Graceful shutdown can be initiated using the endpoint **/shutdown**

## Database and Results

Data is stored and retrieved via JSON to a PostgreSQL database.
All data can be printed using the endpoint **/results**

The Docker container will initialize PostgreSQL and create a persisting database called "webcalls" with a table "calls" if one does not already exist.

## Dockerfile and docker-compose
A Dockerfile and docker-compose is provided to build a Docker image.
The image will spin-up a PostgreSQL database and the API on an Alpine Linux OS.

## Makefile
A Makefile is provided with some simple commands such as "docker-new" which creates a completely fresh Docker image, useful for ensuring caches are cleared and code-changes are built, and "local-run" for running the built executable with the "-port localhost" command line flag. 
