# Use the node image from official Docker Hub
FROM node:16.10.0-alpine3.13

# Install Nginx
RUN apk add --no-cache nginx

# Create the necessary directory for Nginx
RUN mkdir -p /run/nginx

# Set the working directory
WORKDIR /app

# Copy the working directory in the container
COPY package*.json ./

# Install the project dependencies
RUN npm install

# Copy the rest of the project files to the container
COPY . .

# Build the Vue.js application to the production mode to dist folder
RUN npm run build

# Remove default Nginx website
RUN rm -rf /usr/share/nginx/html/*

# Copy the build application to the Nginx web directory
COPY dist /usr/share/nginx/html

# Copy the nginx configuration file
COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf

# Expose the port 80
EXPOSE 80

# Start Nginx to serve the application
CMD ["nginx", "-g", "daemon off;"]
