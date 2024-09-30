FROM node:22.9-alpine AS base

WORKDIR /app

RUN apk update && apk add --no-cache curl bash libc6-compat
RUN corepack enable && corepack prepare pnpm@9.7.0 --activate
