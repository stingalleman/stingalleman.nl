FROM node:lts-alpine

ENV NODE_ENV=production

WORKDIR /usr/stingalleman.dev
COPY ./ ./

RUN yarn install --prod
RUN yarn build

EXPOSE 3000

CMD [ "yarn", "start" ]
