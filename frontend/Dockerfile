FROM node:20-alpine AS builder

# Set the working directory
WORKDIR /app

ARG VITE_API_URL
ENV VITE_API_URL=${VITE_API_URL}

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application's source code
COPY . .

# Build the Svelte application
RUN npm run build

# Start a new stage from scratch
FROM node:20-alpine

# Set the working directory
WORKDIR /app

# Copy the build output from the builder stage
COPY --from=builder /app/build ./build
COPY --from=builder /app/package.json ./package.json

# We only need production dependencies
RUN npm install --omit=dev

# Expose port 3000 (default for sveltekit node adapter)
EXPOSE 3000

# The command to run the server
CMD [ "node", "build" ]