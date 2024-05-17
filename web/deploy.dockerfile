FROM node:20.11.1-alpine AS BUILDER

WORKDIR /web

COPY . .

ARG VITE_API https://api.live-editor.ikeoha.xyz

RUN npm install \
    && npm run build \
    && npm install -g serve

EXPOSE 8001

CMD ["serve", "dist", "-l", "8001"]