FROM node:latest

COPY . /src
WORKDIR /src

RUN npm install --loglevel warn

EXPOSE 4001

CMD [ "npm", "start" ]