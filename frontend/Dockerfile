FROM node:latest

COPY . /src
WORKDIR /src

RUN npm install --loglevel warn

EXPOSE 3000

CMD [ "npm", "start" ]