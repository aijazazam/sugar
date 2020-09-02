# sugar
IMDB architecture

Required Installations:
-------------------------
*. Please go through "IMDB Architecture Doccument.docx" in /docs folder to get complete overview of the architecture.

Database Setup:
----------------
Install MariaDB with root password as "sugarbox". Go to "imdb.sql" in data folder and execute below command to import the data:
		mysql -u root -p < imdb.sql

Kafka Setup:
-------------
Install Kafka and create topic TOPIC-MOVIE, TOPIC-RATING and TOPIC-AUTHENTICATION, using below script:
1. ./bin/zookeeper-server-start.sh ./config/zookeeper.properties
2. ./bin/kafka-server-start.sh ./config/server.properties
3. ./bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic TOPIC-MOVIE
4. ./bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic TOPIC-RATING
5. ./bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic TOPIC-AUTHENTICATION

PostMan Setup:
---------------
1. We are using https with TLS (using self signed certificate), so disable the SSL certificate verification in POST MAN by going to Settings and disable the "SSL certificate verification".
2. Please import "imdb api.postman_collection.json" into PostMan for details on how to use the api's.

Environment SetUp
------------------
1. Please setup the GOPATH and GOBIN folders to this source code folder like in my system its GOPATH is "/home/think/sugar/go" and GOBIN is "/home/think/sugar/go/bin"
2. download the source code using "git clone https://github.com/aijazazam/sugar".
3. download all dependencies: go to "sugar/go/src" folder and do "make init", it will go get all dependencies.
4. compile the code: "make clean" then do "make all" from src folder. I have also commited the executables, in worst case if things don't work out (due to dependencies), please pull the executables from git "go/bin" and try running directly, they are compiled on Ubuntu 18.04
5. Before running all micro services please make sure that mariaDB and Kafka are running in localhost with default port numbers. with mariaDB root password set as "sugarbox"
8. Please make sure that following port numbers are unused since the port numbers used by micro services are apigateway(8000), search(8001), recommendation(8002), useraction(8003).
6. Run all micro services: do a "make run", this will start apigateway, search, recommendation and useraction micro services in background.
7. Hit the api's imported on PostMan from "imdb api.postman_collection.json"
8. to Kill all micro services: do a "make kill" to kill all running micro services.
9. There are two pre populated users in the database, for https://localhost:8000/login POST request use any one of below in body:
	{
		"username": "sugar",
		"password": "sugarbox"
	}
	{
		"username": "admin",
		"password": "adminpassword"
	}
	If you want to rate/comment or check your useraction then you have to pass the jwt in the "Autherization" Header of your request.
	You will get jwt after you hit the login api. copy that jwt into "Autherization" Header of your request for below api's
	https://localhost:8000/useraction
	https://localhost:8000/add/rate
	https://localhost:8000/add/comment
	https://localhost:8000/add/movie  (requires an admin jwt, you get it by logging in as admin credentials above)
	Please see PostMan examples for above api's usage.
   
10.Just for information, the movie data is copied from "https://data.world/promptcloud/imdb-data-from-2006-to-2016".

