# mock-rfid-system
[![Go Report Card](https://goreportcard.com/badge/github.com/smford/narcotk-hosts)](https://goreportcard.com/report/github.com/smford/narcotk-hosts) [![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)


| URL | Output |
|:--|:--|
| `http://localhost:56000/users` | list all users |
| `http://localhost:56000/getuser?rfid=aa` | return all details for user with rfid tag aa |
| `http://localhost:56000/check?rfid=cc&device=laser` | check if user with rfid cc is allowed access to device laser |

All output is using json by default, to disable json simply append `json=n` to the end of the url

## Deployment

This can be deployed via heroku: [![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)
