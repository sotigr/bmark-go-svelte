FROM golang:1.21-bullseye

RUN apt-get update && apt-get install curl -y
RUN curl -fsSL https://deb.nodesource.com/setup_lts.x | bash - 
RUN apt-get install -y nodejs  

WORKDIR /src 
  
CMD ["/bin/bash", "run.sh"]
