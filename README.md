# mock-rfid-system

| URL | Output |
|:--|:--|
| `http://localhost:56000/users` | list all users |
| `http://localhost:56000/getuser?rfid=aa` | return all details for user with rfid tag aa |

All output is using json by default, to disable json simply append `json=n` to the end of the url

## Deployment

This can be deployed via heroku: [![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)
