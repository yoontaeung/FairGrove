import random

def generate_random_test_case(num_transactions, num_sequencers, max_sequence_length):
    transactions = [f"tx{i}" for i in range(num_transactions)]
    pre_local_sequences = []
    
    for _ in range(num_sequencers):
        sequence_length = random.randint(1, max_sequence_length)
        sequence = random.sample(transactions, sequence_length)
        pre_local_sequences.append(sequence)
    
    return pre_local_sequences

def main():
    num_test_cases = 5
    num_transactions = 10
    num_sequencers = 4
    max_sequence_length = 5
    
    for i in range(num_test_cases):
        print(f"Test Case {i + 1}:")
        pre_local_sequences = generate_random_test_case(num_transactions, num_sequencers, max_sequence_length)
        print("[")
        for seq in pre_local_sequences:
            print(f"  {seq},")
        print("]")
        print()

if __name__ == '__main__':
    main()
