require:
+ go version 1.17
+ postgresql: 12

1. create postgresql info
     port: 5432
     userName: postgres
     password: root
     database: health
and change infomation database at: `health-service/config/config.yaml`
![root project config](config_file.png)

And change config at
![content yaml file config](config_file_yaml.png)

At root project health-service
Run command: `go run *.go`

Success as bellow:
![go_run_command](go_run_command.png)

![img_2.png](img_2.png)