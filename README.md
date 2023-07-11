# exercism-challenges
Repo contains my code for solving exercism challenges.

Refer to https://exercism.io/.

# Steps to add a new exercise to Go workspace.

Go Workspace is available at the root of the repository.

Run below command.
```shell
exercism download --exercise=resistor-color --track=go
go work use go/resistor-color
```

Executed below command to create the workspace and all modules.
```shell
go work init
for dir in $(ls go | xargs); do; echo $dir;go work use go/$dir; done
```