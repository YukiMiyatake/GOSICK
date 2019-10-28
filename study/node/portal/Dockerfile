FROM node:10

WORKDIR /GOSICK

COPY package*.json ./

RUN npm install
COPY . .
EXPOSE 3000

#ENV NODE_ENV="development"
CMD [ "npm", "start" ]