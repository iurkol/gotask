## SIMPLE WEB APPLICATION. Golang version

## Assignment 1:
* Setup the application and its dependencies
* Add an endpoint that listens on `/`
  this endpoint simply returns the text `Hello World!`

## Assignment 2:
* Log information about requests the application receives on `/` to a file
* The information should be stored in a text file (/tmp/requests.txt)
* Every new request is appended to the file
* Seperate requests with a `------------`
* Each pice of information goes into one line
* The following information should be stored in the file
  IP-address of the client
  Path (should always be `/`)
  Hostname of the machine received the request (`exec("hostname")` is good enough)
  Timestamp the request was received
* Example file content
  ```
  IP: 172.21.2.1
  Hostname: my-machine
  Path: /
  Timestamp: 2018-01-09 10:38:10 +0000
  ------------
  IP: 172.21.2.1
  Hostname: my-machine
  Path: /
  Timestamp: 2018-01-09 10:38:15 +0000
  ------------
  ```
* The endpoint `/` still returns `Hello World!`


## Assignment 3:
* Instead of storing the requests in a file, they should be stored in a datbase (postgres)
* Host information and credentials to connect to the database are read from environment variables:
  `POSTGRES_USER` (the username)
  `POSTGRES_PASSWORD` (the password)
  `POSTGRES_DB` (the name of the database_)
  `POSTGRES_HOST` (the hostname or ip of the machine running postgres)

* SQL to create the tables:
	```
	CREATE TABLE IF NOT EXISTS requests (
	  id serial primary key,
	  requested_at character varying,
	  ip character varying,
	  host character varying,
	  path character varying
	);
	```
* Example SQL to store a request:
	```
	INSERT INTO requests (ip, path, host, requested_at)
	VALUES ('127.21.0.4', '/', 'my-machine', '2018-01-09 10:38:10 +0000');
	```
* The endpoint `/` still returns `Hello World!`


## Assignment 4:
* Now we want to read the last 25 requests from the database and display them
* Requests are sorted by their ID descending
* Example HTML:
  ```
  <html>
    <body>
      <h1>The requests</h1>
      <table style="width: 100%;" border="1">
        <tr>
          <th>IP</th>
          <th>Path</th>
          <th>Container</th>
          <th>Timestamp</th>
        </tr>
        <!-- Requests go here -->
    </body>
  </html>
  ```
* Example SQL to retreive the requests:
  ```
  SELECT ip, path, host, requested_at
  FROM requests
  ORDER BY id DESC
  LIMIT 25;
  ```
