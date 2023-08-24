import matplotlib.pyplot as plt
import networkx as nx
from utils import tx_generate
import time
from tqdm import tqdm

# Function to generate the node orderings from local sequences
def generate_transaction_orderings(local_sequences):
    # Get a list of all unique transactions
    all_transactions = sorted(set(tx for seq in local_sequences for tx in seq))
    num_transactions = len(all_transactions)
    
    # Create a dictionary to map transactions to indices
    transaction_indices = {tx: i for i, tx in enumerate(all_transactions)}
    
    # Initialize the transaction_orderings array with 'x'
    num_sequencers = len(local_sequences)
    transaction_orderings = [['x'] * num_transactions for _ in range(num_sequencers)]
    
    # Fill in the transaction_orderings array based on local_sequences
    for seq_idx, sequence in enumerate(local_sequences):
        for tx_idx, tx in enumerate(sequence):
            global_tx_idx = transaction_indices[tx]
            transaction_orderings[seq_idx][global_tx_idx] = tx_idx
    
    return all_transactions, transaction_orderings

# Function to construct the directed graph from transaction sequences
def construct_graph(all_transactions, local_sequences):
    num_txs = len(local_sequences[0])  # Number of all transactions
    G = nx.DiGraph()
    G.add_nodes_from(all_transactions)
    
    for first in range(num_txs):
        for second in range(first + 1, num_txs):
            first_counter = 0
            second_counter = 0
            for sequencer in local_sequences:
                if sequencer[first] != 'x' and sequencer[second] != 'x':
                    if sequencer[first] < sequencer[second]:
                        first_counter += 1
                    elif sequencer[second] < sequencer[first]:
                        second_counter += 1
            
            if first_counter > second_counter:
                G.add_edge(all_transactions[first], all_transactions[second])
            elif first_counter < second_counter:
                G.add_edge(all_transactions[second], all_transactions[first])
    
    return G

# Function to visualize the graph
def visualize_graph(G):
    # Use circular layout algorithm to arrange nodes
    pos = nx.circular_layout(G)

    # Draw the graph with the calculated positions
    nx.draw(G, pos, with_labels=True, node_size=500, node_color='lightblue', font_size=8)

    # Display the graph
    plt.show()

# Main function to run the program
def run(pre_local_sequences):
    # pre_local_sequences = tx_generate(100000, 10)
    # #### TODO: the graph that has a big cycle. Do we use a Hamiltonian path????

    # pre_local_sequences = [
    #     ['sdf', 'tx1', 'tx3', 'tx2'],
    #     ['tx4', 'tx1', 'tx5', 'tqw', 'sda', 'sdf'],
    #     ['sdf', 'tqw', 'tx3', 'tx4', 'tx2'],
    #     ['abc', 'def', 'ghi', 'abc'],  # Another SCC
    #     ['ghi', 'def', 'jkl', 'ghi']   # Another SCC
    # ]
    start_time = time.time()
    all_transactions, local_sequences = generate_transaction_orderings(pre_local_sequences)
    
    G = construct_graph(all_transactions, local_sequences)
    # visualize_graph(G)

    # Identify strongly connected components (SCCs)
    scc = list(nx.strongly_connected_components(G))
    # print(scc)

    # Condense the graph and create a condensed graph
    condensed_G = nx.condensation(G, scc)
    # visualize_graph(condensed_G)

    # Perform topological sort on the condensed graph
    topological_order = list(nx.topological_sort(condensed_G))
    # print("TOPO order: ", topological_order)
    
    # Reconstruct the sorted transactions based on condensed node order
    sorted_transactions = []
    for condensed_node in topological_order:
        original_vertices = scc[condensed_node]
        sorted_transactions.extend(sorted(original_vertices))
  

    # print("Sorted Transactions:", flush=True)
    # print(sorted_transactions, flush=True)

    # End measuring the time
    end_time = time.time()
    elapsed_time = end_time - start_time

    # Print the elapsed time in seconds and milliseconds
    return elapsed_time

def main():
    total_time = 0
    num_iterations = 10

    for _ in tqdm(range(num_iterations)):
        pre_local_sequences = tx_generate(100000, 10)
        start_time = time.time()
        run(pre_local_sequences)
        end_time = time.time()
        iteration_time = end_time - start_time
        total_time += iteration_time

    average_time = total_time / num_iterations
    print("Average execution time:", average_time, "seconds")

if __name__ == '__main__':
    main()
