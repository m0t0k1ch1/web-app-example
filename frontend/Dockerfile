FROM node:18-alpine

WORKDIR /app

COPY package.json pnpm-lock.yaml ./
RUN corepack enable pnpm && pnpm i

COPY . .

CMD ["pnpm", "ng", "serve", "--host", "0.0.0.0"]
