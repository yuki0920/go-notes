FROM --platform=linux/amd64 node:14.17.3-alpine

RUN apk add --no-cache python2 g++ make

ARG WORKDIR
ARG PORT
ARG API_URL

ENV HOME=/${WORKDIR} \
  LANG=C.UTF-8 \
  TZ=Asia/Tokyo \
  HOST=0.0.0.0  \
  API_URL=${API_URL}

WORKDIR ${HOME}

COPY package*.json ./
RUN yarn install

COPY . .

EXPOSE ${PORT}
