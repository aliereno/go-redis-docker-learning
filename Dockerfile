FROM golang:1.15.8-buster

ARG UNAME=aliereno
ARG UID=1000
ARG GID=1000

RUN groupadd -g $GID -o $UNAME
RUN useradd -m -u $UID -g $GID -o -s /bin/bash $UNAME

RUN mkdir /code
COPY . /code/

CMD chown -R $UNAME:$UNAME /code

USER $UNAME
CMD /bin/bash

WORKDIR /code