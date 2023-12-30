CREATE ENV FILE : 

    HTTP_PORT=*

INSTALL GLOBALLY NODEMON : 

    - npm i -g nodemon
    

RUN : 

    - go run .

    RUN DEV : 
        
        - nodemon --exec "go run ." ./main.go --signal SIGTERM



NOTES : 

    dependencies commands : 

        - go install *name

        - go get *name