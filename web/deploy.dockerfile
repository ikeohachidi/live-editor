FROM node:20.11.1-alpine AS BUILDER

WORKDIR /web

COPY . .

ENV VITE_API https://api.live-editor.ikeoha.xyz

RUN npm install -g serve \
    && npm run build

EXPOSE 8001

CMD ["serve", "dist", "-l", "8001"]