from google.cloud import pubsub_v1, firestore

# Initialize Firestore and Pub/Sub clients
db = firestore.Client()
subscriber = pubsub_v1.SubscriberClient()

# Define subscription and Firestore collection
project_id = "mlb-predictor-448606"
subscription_id = "mlb-live-feed-sub"
subscription_path = f"projects/{project_id}/subscriptions/{subscription_id}"
collection_name = "predictions"

def callback(message):
    print(f"Received message: {message.data.decode('utf-8')}")

    # Parse the message and write to Firestore
    data = message.data.decode("utf-8")
    doc_ref = db.collection(collection_name).document()
    doc_ref.set({"prediction": data})

    print("Message written to Firestore.")
    message.ack()

# Subscribe to the Pub/Sub subscription
future = subscriber.subscribe(subscription_path, callback=callback)

print(f"Listening for messages on {subscription_path}...")
try:
    future.result()
except KeyboardInterrupt:
    future.cancel()

