FROM python:3-slim

WORKDIR /code

COPY requirements.txt ./requirements.txt
RUN pip3 install -r requirements.txt

COPY fastquote.py /code/fastquote.py
COPY lib /code/lib

ENV PYTHONUNBUFFERED=1
ENV PYTHONDONTWRITEBYTECODE=1

CMD ["python", "fastquote.py"]
