 #!/usr/bin/sh

 go test -v ./...
 if [ $? -ne 0 ]; then
   echo "unit tests failed"
   exit 1
 fi