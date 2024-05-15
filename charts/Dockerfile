FROM nginxinc/nginx-unprivileged:stable-alpine

ARG IMAGE_EXPIRATION=Never

LABEL Name=charts \
      Release=https://github.com/appvia/wayfinder \
      Maintainer=wayfinder@appvia.io \
      Url=https://github.com/appvia/wayfinder \
      Help=https://github.com/appvia/wayfinder/issues \
      quay.expires-after=${IMAGE_EXPIRATION}

COPY index.yaml /usr/share/nginx/html/index.yaml.tmpl
COPY *.tgz /usr/share/nginx/html/

USER 0

RUN echo "Wayfinder Helm Chart Repository" > /usr/share/nginx/html/index.html

RUN touch /usr/share/nginx/html/index.yaml
RUN chown 101:101 /usr/share/nginx/html/index.yaml

COPY set-charts-url.sh /docker-entrypoint.d/99-set-charts-url.sh

RUN chmod +x /docker-entrypoint.d/99-set-charts-url.sh

USER 101
