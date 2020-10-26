FROM node:10.22.1-alpine3.11 AS build-env
ADD . /app
WORKDIR /app

RUN npm i --only=production

FROM gcr.io/distroless/nodejs:10
COPY --from=build-env /app /app
WORKDIR /app
CMD ["server.js"]