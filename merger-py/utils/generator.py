import hashlib
import json
import random
import time

"""
This generator generates random transaction.
The transactions are used to construct a graph.
"""
# hash_length = 5

class Transaction:
    def __init__(self, from_addr, to_addr, value, gas_price, gas_limit, nonce, data):
        self.from_addr = from_addr
        self.to_addr = to_addr
        self.value = value
        self.gas_price = gas_price
        self.gas_limit = gas_limit
        self.nonce = nonce
        self.data = data

def calculate_transaction_hash(transaction):
    transaction_json = json.dumps(transaction.__dict__).encode('utf-8')
    hash_value = hashlib.sha256(transaction_json).hexdigest()
    return hash_value[:]

def generate_random_transaction():
    from_addr = "0x" + generate_random_hex(40)
    to_addr = "0x" + generate_random_hex(40)
    value = generate_random_hex(18)
    gas_price = generate_random_hex(10)
    gas_limit = generate_random_hex(6)
    nonce = generate_random_hex(8)
    data = generate_random_hex(64)
    return Transaction(from_addr, to_addr, value, gas_price, gas_limit, nonce, data)

def generate_random_hex(length):
    return ''.join(random.choice('0123456789abcdef') for _ in range(length))

def generate_transactions(k):
    transactions = []
    transaction_hashes = []
    for _ in range(k):
        transaction = generate_random_transaction()
        transactions.append(transaction)
        transaction_hash = calculate_transaction_hash(transaction)
        transaction_hashes.append(transaction_hash)
    return transactions, transaction_hashes

def tx_generate(k, n):
    # k: number of all transactions, n: number of sequences
    random.seed(time.time())

    transactions, transaction_hashes = generate_transactions(k)

    all_transaction_hashes = transaction_hashes

    max_transactions_per_file = 100
    min_transactions_per_sequence = 2

    if k < min_transactions_per_sequence:
        min_transactions_per_sequence = k

    used_indices = 0
    sequence_transaction_hashes = []
    for i in range(1, n + 1):
        shared_indices = list(range(k))
        random.shuffle(shared_indices)
        num_transactions = random.randint(min_transactions_per_sequence, max_transactions_per_file)

        if used_indices + num_transactions > k:
            random.shuffle(shared_indices)
            used_indices = 0

        selected_indices = shared_indices[used_indices:used_indices + num_transactions]
        used_indices += num_transactions

        selected_transaction_hashes = [transaction_hashes[index] for index in selected_indices]
        sequence_transaction_hashes.append(selected_transaction_hashes)

    return sequence_transaction_hashes


