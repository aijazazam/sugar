# sugar
IMDB architecture

Required Installations:
-------------------------
*. Please go through "IMDB Architecture Doccument.docx" in /docs folder to get complete overview of the architecture.
1. Please setup the GOPATH and GOBIN folders to this source code folder like in my system its GOPATH is "/home/think/sugar/go" and GOBIN is "/home/think/sugar/go/bin"
2. Install MariaDB with root password as "sugarbox". The mariaDB backup (imdb.sql) is stored in "data" folder. Please install mariaDB and import "imdb.sql"
3. Install Kafka and create topic TOPIC-MOVIE and TOPIC-AUTHENTICATION, using below script:
   ./bin/zookeeper-server-start.sh ./config/zookeeper.properties
   ./bin/kafka-server-start.sh ./config/server.properties
   ./bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic TOPIC-MOVIE
   ./bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic TOPIC-AUTHENTICATION

4. We are using https with TLS (using self signed certificate), so disable the SSL certificate verification in POST MAN.
   In PostMan goto Settings and disable the "SSL certificate verification"
5. For details how to use api's please see/import the PostMan samples in data folder (imdb api.postman_collection.json).    
6. cd to src folder and do a "make clean" than do "make all". I have also commited the executables, in case things don't work out (due to dependencies), please pull the executables from git "go/bin" and try running.
7. Open a new terminal and start "apigateway", again in another terminal start "search", again in another terminal start "recommendation", again in another terminal start "useraction".
8. Please make sure that following port numbers are unused since the port numbers used by micro services are apigateway(8000), search(8001), recommendation(8002), useraction(8003).
9. Please make sure that mariaDB and Kafka are running in localhost with default port numbers.
9. There are two pre populated users in the database, for https://localhost:8000/login POST request use any one of below
	{
		"username": "sugar",
		"password": "sugarbox"
	}
	{
		"username": "admin",
		"password": "adminpassword"
	}
	After /login after you get "jwt" in the response body, please copy that in "Autherization" Header as type "Bearer Token" for following apis:
	https://localhost:8000/useraction
	https://localhost:8000/add/rate
	https://localhost:8000/add/comment
	https://localhost:8000/add/movie  (requires an admin jwt)
   
10.Just for information, the movie data is copied from "https://data.world/promptcloud/imdb-data-from-2006-to-2016".

