FROM golang:1.19.4-bullseye

RUN apt-get update && apt-get install curl -y
RUN curl -fsSL https://deb.nodesource.com/setup_lts.x | bash - 
RUN apt-get install -y nodejs  

WORKDIR /src 

# CMD ["/bin/bash", "-c", "npx nodemon & npx nodemon --config nodemon-svelte.json"]

# CMD ["/bin/bash", "-c", " npm run build --prefix=svelte & npx nodemon "]


CMD ["/bin/bash", "-c", "npx nodemon & npm run dev --prefix=svelte"]

# ENTRYPOINT ["tail"]
# CMD ["-f","/dev/null"]