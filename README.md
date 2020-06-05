# mock-rfid-system
[![Go Report Card](https://goreportcard.com/badge/github.com/smford/narcotk-hosts)](https://goreportcard.com/report/github.com/smford/narcotk-hosts) [![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

## Demo

A demonstration of the mock-rfid-system is available here [https://mock-rfid-system.herokuapp.com/](https://mock-rfid-system.herokuapp.com/) which is hosted on Heroku

---

## Usage
### Web API Examples

| URL | Output |
|:--|:--|
| `http://localhost:56000/users` | list all users |
| `http://localhost:56000/getuser?rfid=aa` | return all details for user with rfid tag aa |
| `http://localhost:56000/check?rfid=cc&device=laser` | check if user with rfid cc is allowed access to device laser |

All output is using json by default, to disable json simply append `json=n` to the end of the url

### Command Line

`./mock-rfid-system` or `mock-rfid-system` depending on how you installed it

| Command | Description | Example |
|:--|:--|:--|
| `--help` | Display help information | `mock-rfid-system --help` |
| `--listenip` | IP to listen on (default all) | `mock-rfid-system --listenip 192.168.1.45` |
| `--listenport` | Port to listen on (default 56000) | `mock-rfid-system --listenport 8090` |
| `--listusers` | List users and their permissions | `mock-rfid-system --listusers` |

---

## Installation 

### Heroku
This can be deployed via heroku: [![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

### Git
```bash
git clone git@github.com:smford/mock-rfid-system.git
cd mock-rfid-system
dep ensure
go build -o mock-rfid-system
./mock-rfid-system
```

### Go
```bash
go get -v github.com/smford/mock-rfid-system
mock-rfid-system
```
