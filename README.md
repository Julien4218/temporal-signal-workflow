# temporal-signal-workflow

This is a repo to evaluate starting and signaling a workflow without being a worker. The worker can be started from this repo https://github.com/Julien4218/money-transfer-project-template-go

## example

### creating a workflow

```bash
go run main.go start --name MoneyTransfer --input ewogICAiU291cmNlQWNjb3VudCI6Ijg1LTE1MCIsCiAgICJUYXJnZXRBY2NvdW50IjoiNDMtODEyIiwKICAgIkFtb3VudCI6MjUxLAogICAiUmVmZXJlbmNlSUQiOiIxMjM0NSIKfQ==
```

### signaling a workflow

```bash
go run main.go signal --token QmFja2dyb3VuZENoZWNrOnVzZXJuYW1lKzIwMjMxMDI1MTQzNEBleGFtcGxlLmNvbS8zYjgyZmE0Ni03YmI3LTQ1YmUtYTE3Yi1iMmNmOWRlMDM1MWE= --body x
```
