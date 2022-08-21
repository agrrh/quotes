FROM python:3.10-alpine

WORKDIR /app

RUN apk add --no-cache --virtual .build-deps \
  gcc \
  musl-dev

COPY ./requirements.txt ./
RUN pip install -r requirements.txt

RUN apk del .build-deps

COPY ./ ./

CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "3000"]
