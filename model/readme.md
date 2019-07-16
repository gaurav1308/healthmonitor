In this project we have created a small service to monitor the health of external urls.
The system accepts a list of http/https urls and a crawl_timeout(milliseconds) and frequency(in seconds) and failure_threshold(count)


crawl_timeout : System will wait for this much time before giving up on the url
frequency : System will wait for this much time before retrying again.
failure_threshold : count of retries possible for that url


The system iterates over all the urls in the database and try to get a HTTP status of the URLs (wait for the crawl_timeout) seconds before giving up on the URL. 

There is a route that shows status (Health) of a particular URL in a prticular try.