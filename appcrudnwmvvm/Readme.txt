// go app container to perform CRUD with MYSQL container

// Main.go file is changed to cater for the CRUD


// Updated  : Ismail 
// Chg Date : 24/07/222

//  Clear images and container
docker image prune
docker container prune
docker compose down
///


// same private nw

1. Mysql container
   docker compose -f docker-compose-my-sql.yaml up

2. Go container
  docker compose -f docker-compose-web-app.yaml up

check network

Faziths-MacBook-Air:appcrudnwmvvm fazithmohamed$ docker network inspect mysql_network
[
    {
        "Name": "mysql_network",
        "Id": "e2f2f027e2dadf1c113df4d9f2135a112a39c6cfbd0ae64c98a03ffc40d6b8ef",
        "Created": "2022-07-24T04:51:23.610317375Z",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": null,
            "Config": [
                {
                    "Subnet": "172.27.0.0/16",
                    "Gateway": "172.27.0.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {
            "41505715c86cf6d8466c47a2175cedb044eb9d0eee0d76d16dbd8ec89ab19896": {
                "Name": "mysql-db",
                "EndpointID": "312cfb84b5eb59e3c7493f7c0adc19e7b173c401dd54031ccd7f68502f83c311",
                "MacAddress": "02:42:ac:1b:00:02",
                "IPv4Address": "172.27.0.2/16",
                "IPv6Address": ""
            }
        },
        "Options": {},
        "Labels": {
            "com.docker.compose.network": "network",
            "com.docker.compose.project": "appcrudnwmvvm",
            "com.docker.compose.version": "2.6.1"
        }
    }
]
Faziths-MacBook-Air:appcrudnwmvvm fazithmohamed$ docker network inspect mysql_network
[
    {
        "Name": "mysql_network",
        "Id": "e2f2f027e2dadf1c113df4d9f2135a112a39c6cfbd0ae64c98a03ffc40d6b8ef",
        "Created": "2022-07-24T04:51:23.610317375Z",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": null,
            "Config": [
                {
                    "Subnet": "172.27.0.0/16",
                    "Gateway": "172.27.0.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {
            "41505715c86cf6d8466c47a2175cedb044eb9d0eee0d76d16dbd8ec89ab19896": {
                "Name": "mysql-db",
                "EndpointID": "312cfb84b5eb59e3c7493f7c0adc19e7b173c401dd54031ccd7f68502f83c311",
                "MacAddress": "02:42:ac:1b:00:02",
                "IPv4Address": "172.27.0.2/16",
                "IPv6Address": ""
            },
            "edc5ba11f05b7f82adc87d1706ced03bbb511f9bdc1ad00039309ddebc8fd7be": {
                "Name": "web-app",
                "EndpointID": "5ce4ebc3303eaa434fa30279632d4347cf4ffebab31a677802573d13c189572c",
                "MacAddress": "02:42:ac:1b:00:03",
                "IPv4Address": "172.27.0.3/16",
                "IPv6Address": ""
            }
        },
        "Options": {},
        "Labels": {
            "com.docker.compose.network": "network",
            "com.docker.compose.project": "appcrudnwmvvm",
            "com.docker.compose.version": "2.6.1"
        }
    }
]
Faziths-MacBook-Air:appcrudnwmvvm fazithmohamed$ 

