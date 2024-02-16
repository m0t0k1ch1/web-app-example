FROM node:18-alpine

WORKDIR /app

COPY package.json pnpm-lock.yaml ./
RUN corepack enable pnpm && pnpm i

COPY . .
RUN pnpm build

CMD ["pnpm", "start"]
