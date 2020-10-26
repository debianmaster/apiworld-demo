FROM node:12.14.0 AS build-env
WORKDIR /app
ADD package.json /app
RUN yarn
ADD . /app
RUN GENERATE_SOURCEMAP=false npm run build

FROM nginx:alpine
COPY --from=build-env /app/build /usr/share/nginx/html
RUN apk add --no-cache bash
WORKDIR /usr/share/nginx/html
EXPOSE 3000
CMD ["/bin/bash", "-c", "nginx -g \"daemon off;\""]