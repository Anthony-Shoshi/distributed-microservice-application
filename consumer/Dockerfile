FROM python:3.12-slim

WORKDIR /app
COPY main.py .
RUN pip install pika
CMD ["python", "main.py"]