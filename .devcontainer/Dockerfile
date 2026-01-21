FROM golang

RUN apt-get update && apt-get install -y git ripgrep curl wget

RUN useradd -ms /bin/bash dev-user
USER dev-user

WORKDIR /home/dev-user
RUN mkdir -p .local/share
RUN go install golang.org/x/tools/gopls@latest

ENV PATH=$PATH:/opt/nvim-linux-x86_64/bin

