FROM nginx:alpine

# # Set working directory to nginx asset directory
# WORKDIR /usr/share/nginx/html
RUN echo | ls .

# Remove default nginx static assets
RUN rm -rf /usr/share/nginx/html/*

# Copy static assets over
COPY ./dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

ENTRYPOINT ["nginx", "-g", "daemon off;"]