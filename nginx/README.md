# Configuration via Volumes

Nginx configuration is done via files. You'll need to use volumes to mount your
own configuration within the container. Here are the configuration points for
Nginx on Ubuntu Trusty (the basis for this container).

* /etc/nginx/conf.d/*.conf
* /etc/nginx/sites-enabled/*

Additional files, such as SSL certificates, may be mounted as desired and
referenced in your configs.
