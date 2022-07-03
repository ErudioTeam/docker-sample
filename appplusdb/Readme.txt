Faziths-MacBook-Air:appplusdb fazithmohamed$ docker-compose up
Pulling db (mysql:8.0)...
8.0: Pulling from library/mysql
824b15f81d65: Pull complete
c559dd1913db: Pull complete
e201c19614e6: Pull complete
f4247e8f6125: Pull complete
dc9fefd8cfb5: Pull complete
af3787edd16d: Pull complete
b6bb40f875d3: Pull complete
75f6b647ddb1: Pull complete
a09ca0f0cb24: Pull complete
9e223e3cd2fd: Pull complete
2b038d826c65: Pull complete
d33ac6052fc9: Pull complete
Digest: sha256:a840244706a5fdc3c704b15a3700bfda39fdc069262d7753fa09de2d9faf5f83
Status: Downloaded newer image for mysql:8.0
Building go
[+] Building 9.2s (12/12) FINISHED                                                                      
 => [internal] load build definition from DockerFile                                               0.0s
 => => transferring dockerfile: 36B                                                                0.0s
 => [internal] load .dockerignore                                                                  0.0s
 => => transferring context: 2B                                                                    0.0s
 => [internal] load metadata for docker.io/library/golang:latest                                   2.9s
 => [1/7] FROM docker.io/library/golang:latest@sha256:a452d6273ad03a47c2f29b898d6bb57630e77baf839  0.0s
 => [internal] load build context                                                                  0.0s
 => => transferring context: 14.21kB                                                               0.0s
 => CACHED [2/7] RUN mkdir /app                                                                    0.0s
 => CACHED [3/7] WORKDIR /app                                                                      0.0s
 => [4/7] COPY go.mod go.sum ./                                                                    0.0s
 => [5/7] RUN go mod download                                                                      2.4s
 => [6/7] COPY . .                                                                                 0.0s
 => [7/7] RUN go build -o main .                                                                   3.0s
 => exporting to image                                                                             0.8s
 => => exporting layers                                                                            0.8s
 => => writing image sha256:aaa3fa0e98482233b180c4050961a50e343b86c661a34b7cb00ab27d8a3ac26e       0.0s
 => => naming to docker.io/library/appplusdb_go                                                    0.0s

Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them
WARNING: Image for service go was built because it did not already exist. To rebuild this image you must use `docker-compose build` or `docker-compose up --build`.
Creating appplusdb_db_1 ... done
Creating appplusdb_go_1 ... done
Attaching to appplusdb_db_1, appplusdb_go_1
db_1  | 2022-07-03 23:41:33+09:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 8.0.29-1debian10 started.
db_1  | 2022-07-03 23:41:34+09:00 [Note] [Entrypoint]: Switching to dedicated user 'mysql'
db_1  | 2022-07-03 23:41:34+09:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 8.0.29-1debian10 started.
db_1  | 2022-07-03 23:41:35+09:00 [Note] [Entrypoint]: Initializing database files
db_1  | 2022-07-03T14:41:35.288623Z 0 [System] [MY-013169] [Server] /usr/sbin/mysqld (mysqld 8.0.29) initializing of server in progress as process 100
db_1  | 2022-07-03T14:41:35.361877Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
db_1  | 2022-07-03T14:41:35.382606Z 1 [ERROR] [MY-012585] [InnoDB] Linux Native AIO interface is not supported on this platform. Please check your OS documentation and install appropriate binary of InnoDB.
db_1  | 2022-07-03T14:41:35.382946Z 1 [Warning] [MY-012654] [InnoDB] Linux Native AIO disabled.
db_1  | 2022-07-03T14:41:35.842937Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
db_1  | 2022-07-03T14:41:39.291970Z 6 [Warning] [MY-010453] [Server] root@localhost is created with an empty password ! Please consider switching off the --initialize-insecure option.
db_1  | 2022-07-03 23:41:44+09:00 [Note] [Entrypoint]: Database files initialized
db_1  | 2022-07-03 23:41:44+09:00 [Note] [Entrypoint]: Starting temporary server
db_1  | 2022-07-03T14:41:45.430414Z 0 [System] [MY-010116] [Server] /usr/sbin/mysqld (mysqld 8.0.29) starting as process 155
db_1  | 2022-07-03T14:41:45.548120Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
db_1  | 2022-07-03T14:41:45.579808Z 1 [ERROR] [MY-012585] [InnoDB] Linux Native AIO interface is not supported on this platform. Please check your OS documentation and install appropriate binary of InnoDB.
db_1  | 2022-07-03T14:41:45.580127Z 1 [Warning] [MY-012654] [InnoDB] Linux Native AIO disabled.
db_1  | 2022-07-03T14:41:45.847214Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
db_1  | 2022-07-03T14:41:46.944151Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
db_1  | 2022-07-03T14:41:46.944917Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
db_1  | 2022-07-03T14:41:46.951659Z 0 [Warning] [MY-011810] [Server] Insecure configuration for --pid-file: Location '/var/run/mysqld' in the path is accessible to all OS users. Consider choosing a different directory.
db_1  | 2022-07-03T14:41:47.083863Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Socket: /var/run/mysqld/mysqlx.sock
db_1  | 2022-07-03T14:41:47.086363Z 0 [System] [MY-010931] [Server] /usr/sbin/mysqld: ready for connections. Version: '8.0.29'  socket: '/var/run/mysqld/mysqld.sock'  port: 0  MySQL Community Server - GPL.
db_1  | 2022-07-03 23:41:47+09:00 [Note] [Entrypoint]: Temporary server started.
db_1  | Warning: Unable to load '/usr/share/zoneinfo/iso3166.tab' as time zone. Skipping it.
db_1  | Warning: Unable to load '/usr/share/zoneinfo/leap-seconds.list' as time zone. Skipping it.
db_1  | Warning: Unable to load '/usr/share/zoneinfo/zone.tab' as time zone. Skipping it.
db_1  | Warning: Unable to load '/usr/share/zoneinfo/zone1970.tab' as time zone. Skipping it.
db_1  | 2022-07-03 23:41:57+09:00 [Note] [Entrypoint]: Creating database go_database
db_1  | 2022-07-03 23:41:57+09:00 [Note] [Entrypoint]: Creating user go_test
db_1  | 2022-07-03 23:41:57+09:00 [Note] [Entrypoint]: Giving user go_test access to schema go_database
db_1  | 
db_1  | 2022-07-03 23:41:57+09:00 [Note] [Entrypoint]: Stopping temporary server
db_1  | 2022-07-03T14:41:57.611738Z 13 [System] [MY-013172] [Server] Received SHUTDOWN from user root. Shutting down mysqld (Version: 8.0.29).
db_1  | 2022-07-03T14:41:58.590603Z 0 [System] [MY-010910] [Server] /usr/sbin/mysqld: Shutdown complete (mysqld 8.0.29)  MySQL Community Server - GPL.
db_1  | 2022-07-03 23:41:58+09:00 [Note] [Entrypoint]: Temporary server stopped
db_1  | 
db_1  | 2022-07-03 23:41:58+09:00 [Note] [Entrypoint]: MySQL init process done. Ready for start up.
db_1  | 
db_1  | 2022-07-03T14:41:59.131031Z 0 [System] [MY-010116] [Server] /usr/sbin/mysqld (mysqld 8.0.29) starting as process 1
db_1  | 2022-07-03T14:41:59.224407Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
db_1  | 2022-07-03T14:41:59.252542Z 1 [ERROR] [MY-012585] [InnoDB] Linux Native AIO interface is not supported on this platform. Please check your OS documentation and install appropriate binary of InnoDB.
db_1  | 2022-07-03T14:41:59.252923Z 1 [Warning] [MY-012654] [InnoDB] Linux Native AIO disabled.
db_1  | 2022-07-03T14:41:59.535257Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
db_1  | 2022-07-03T14:42:00.354326Z 0 [Warning] [MY-010068] [Server] CA certificate ca.pem is self signed.
db_1  | 2022-07-03T14:42:00.355060Z 0 [System] [MY-013602] [Server] Channel mysql_main configured to support TLS. Encrypted connections are now supported for this channel.
db_1  | 2022-07-03T14:42:00.359820Z 0 [Warning] [MY-011810] [Server] Insecure configuration for --pid-file: Location '/var/run/mysqld' in the path is accessible to all OS users. Consider choosing a different directory.
db_1  | 2022-07-03T14:42:00.492348Z 0 [System] [MY-011323] [Server] X Plugin ready for connections. Bind-address: '::' port: 33060, socket: /var/run/mysqld/mysqlx.sock
db_1  | 2022-07-03T14:42:00.494233Z 0 [System] [MY-010931] [Server] /usr/sbin/mysqld: ready for connections. Version: '8.0.29'  socket: '/var/run/mysqld/mysqld.sock'  port: 3306  MySQL Community Server - GPL.
go_1  | ..........................
go_1  | DB connection successful
go_1  | [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
go_1  | 
go_1  | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
go_1  |  - using env:   export GIN_MODE=release
go_1  |  - using code:  gin.SetMode(gin.ReleaseMode)
go_1  | 
go_1  | panic: html/template: pattern matches no files: `templates/*.html`
go_1  | 
go_1  | goroutine 1 [running]:
go_1  | html/template.Must(...)
go_1  |         /usr/local/go/src/html/template/template.go:374
go_1  | github.com/gin-gonic/gin.(*Engine).LoadHTMLGlob(0x40005bc1a0, {0x59413b, 0x10})
go_1  |         /go/pkg/mod/github.com/gin-gonic/gin@v1.8.1/gin.go:251 +0x2a0
go_1  | main.main()
go_1  |         /app/main.go:17 +0x64
go_1  | exit status 2
appplusdb_go_1 exited with code 1
