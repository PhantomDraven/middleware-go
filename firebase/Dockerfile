FROM node:alpine

RUN apk add openjdk11 curl

RUN npm install -g firebase-tools

WORKDIR /app

CMD [ "firebase", "--project=middleware-go", "emulators:start", "--only", "auth,database,firestore" ]