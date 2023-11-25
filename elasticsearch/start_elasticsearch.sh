 docker run -v `pwd`/data:/usr/share/elasticsearch/data -d --rm --name elasticsearch -p 9200:9200 -p 5002:5002 -e "discovery.type=single-node" elasticsearch:7.9.3


