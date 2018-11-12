FROM ubuntu

EXPOSE 8000

RUN mkdir /app

COPY plafws /app/plafws


CMD [ "/app/plafws"]