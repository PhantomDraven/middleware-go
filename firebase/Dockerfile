FROM node:alpine

RUN apk add openjdk11

RUN npm install -g firebase-tools

WORKDIR /app

CMD [ "firebase", "--project=middleware-go", "emulators:start", "--only", "database,firestore" ]