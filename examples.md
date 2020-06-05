| Locally Hosted | Heroku Hosted | Output |
|:--|:--|:--|
| http://localhost:56000/users | https://mock-rfid-system.herokuapp.com/users | list all users in json format |
| http://localhost:56000/users?json=n | https://mock-rfid-system.herokuapp.com/users?json=n | list all users in plain text format |
| http://localhost:56000/getuser?rfid=aa | https://mock-rfid-system.herokuapp.com/getuser?rfid=aa | return all details for user with rfid tag aa in json format|
| http://localhost:56000/getuser?rfid=aa | https://mock-rfid-system.herokuapp.com/getuser?rfid=aa&json=n | return all details for user with rfid tag aa in plain text format |
| http://localhost:56000/check?rfid=cc&device=laser | https://mock-rfid-system.herokuapp.com/?rfid=cc&device=laser | check if user with rfid cc is allowed access to device laser in json format |
| http://localhost:56000/check?rfid=cc&device=laser | https://mock-rfid-system.herokuapp.com/?rfid=cc&device=laser&json=n | check if user with rfid cc is allowed access to device laser in plain text format |
