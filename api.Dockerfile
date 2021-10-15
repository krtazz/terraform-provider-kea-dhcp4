FROM nginx
ADD test-data/nginx/.htpasswd /etc/nginx/.htpasswd
ADD test-data/nginx/*.conf /etc/nginx/conf.d/
EXPOSE 8080/tcp
CMD ["nginx", "-g", "daemon off;"]
