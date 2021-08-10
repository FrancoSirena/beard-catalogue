FROM node as build

COPY ./src src
COPY ./package.json package.json
COPY ./yarn.lock yarn.lock
RUN yarn install
RUN yarn build

FROM nginx

COPY --from=build dist /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/conf.d/default.conf
