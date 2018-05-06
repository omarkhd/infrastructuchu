#!/usr/bin/env python3
import hashlib
import multiprocessing
import random
import string
from time import sleep

import pymongo

from settings import MONGODB_HOST, MONGODB_REPLICA_SET


def randstr():
	sysrand = random.SystemRandom()
	choices = [sysrand.choice(string.digits + string.ascii_letters) for i in range(40)]
	return ''.join(choices)

def md5(arg):
	md5 = hashlib.md5()
	md5.update(arg)
	return md5.hexdigest()


def infinite_insert(collection, doc, attempt=0):
	sleep(0.1)  # Just a delay to avoid bombing the server.
	try:
		doc_id = collection.insert_one(doc).inserted_id
		retry_msg = 'RETRIED {} times' if attempt > 0 else ''
		print('{} inserted as {} {}'.format(doc, doc_id, retry_msg))
	except pymongo.errors.AutoReconnect:
		print('AutoReconnect on {}, Retrying...'.format(doc))
		infinite_insert(collection, doc, attempt + 1)
	except pymongo.errors.OperationFailure:
		print('OperationFailure on {}, Retrying...'.format(doc))
		infinite_insert(collection, doc, attempt + 1)


def start(worker_index):
	client = pymongo.MongoClient(MONGODB_HOST, replicaset=MONGODB_REPLICA_SET)
	database = client.example

	while True:
		random_string = randstr()
		random_hash = md5(random_string.encode('utf-8'))
		collection = database[random_hash[:2]]
		infinite_insert(collection, {'random': random_string})


for i in range(5):
	process = multiprocessing.Process(target=start, args=[i])
	process.start()

