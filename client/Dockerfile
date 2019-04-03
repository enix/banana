FROM node:8.12.0 as bundler

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY package.json /app
COPY src /app/src
COPY res/public /app/public

RUN yarn

CMD ["bash", "-c", "cd /app && yarn start"]

# ---------------------------

FROM node:8.12.0 as builder

WORKDIR /app

COPY --from=bundler /app /app

RUN ls

ENV PATH /app/node_modules/.bin:$PATH

RUN yarn run build

# ---------------------------

FROM nginx:1.15-alpine as runner

RUN rm -rf /etc/nginx/conf.d/*

COPY res/nginx.conf /etc/nginx/conf.d/nginx.conf
# COPY ./docker/tls/** /etc/nginx/tls/
COPY --from=builder /app/build /usr/share/nginx/html

EXPOSE 80 443

CMD ["nginx", "-g", "daemon off;"]
