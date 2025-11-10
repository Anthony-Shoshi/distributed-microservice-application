import pika
import json
import time
import os
import sys

# Force unbuffered output
sys.stdout.reconfigure(line_buffering=True)

rabbit_url = os.getenv("RABBITMQ_URL", "amqp://guest:guest@rabbitmq:5672/")
params = pika.URLParameters(rabbit_url)

# Retry until RabbitMQ is ready
for i in range(10):
    try:
        connection = pika.BlockingConnection(params)
        channel = connection.channel()
        channel.queue_declare(queue='southpark_messages', durable=True)
        print("Connected to RabbitMQ, waiting for messages...", flush=True)
        break
    except Exception as e:
        print(f"RabbitMQ not ready ({e}), retrying in 3s...", flush=True)
        time.sleep(3)
else:
    print("Failed to connect to RabbitMQ after retries.", flush=True)
    exit(1)

def callback(ch, method, properties, body):
    try:
        message = json.loads(body)
        author = message.get("author", "Unknown")
        text = message.get("body", "")
        print(f"{author} says: {text}", flush=True)
        sys.stdout.flush()
    except json.JSONDecodeError:
        print("Received invalid message:", body, flush=True)

channel.basic_consume(queue='southpark_messages', on_message_callback=callback, auto_ack=True)

print("Listening for messages...")
channel.start_consuming()