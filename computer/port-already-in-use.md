# When the port is already in use, how to find the using app and manage it

Reference link: https://stackoverflow.com/questions/39322089/node-js-port-3000-already-in-use-but-it-actually-isnt

## Find the app (in Windows)

`$ netstat -anob | findstr <:port>`

ex) `$ netstat -anob | findstr :21`


a means all

b means process name <in administartor mode>

n means ip and port name as number

0 means owner process id

## Kill the app

`$ tskill <PID>`

ex) `$ tskill 1234`



