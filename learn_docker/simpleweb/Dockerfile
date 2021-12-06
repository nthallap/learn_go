FROM node:17-alpine3.12

RUN mkdir -p /home
WORKDIR /home
COPY ./ /home

RUN npm install

CMD ["npm", "start"]