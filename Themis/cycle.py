import matplotlib.pyplot as plt
import networkx as nx

def generate_node_orderings(local_sequences):
    # Get a list of all unique transactions
    all_transactions = sorted(set(tx for seq in local_sequences for tx in seq))
    num_transactions = len(all_transactions)
    
    # Create a dictionary to map transactions to indices
    transaction_indices = {tx: i for i, tx in enumerate(all_transactions)}
    
    # Initialize the node_orderings array with 'x'
    num_sequencers = len(local_sequences)
    node_orderings = [['x'] * num_transactions for _ in range(num_sequencers)]
    
    # Fill in the node_orderings array based on local_sequences
    for seq_idx, sequence in enumerate(local_sequences):
        for tx_idx, tx in enumerate(sequence):
            global_tx_idx = transaction_indices[tx]
            node_orderings[seq_idx][global_tx_idx] = tx_idx
    
    return node_orderings

def construct_graph(local_sequences):
    num_txs = len(local_sequences[0])  # Number of all transactions
    G = nx.DiGraph()
    G.add_nodes_from(range(num_txs))
    
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
                G.add_edge(first, second)
            elif first_counter < second_counter:
                G.add_edge(second, first)
    
    return G

def visualize_graph(G):
    nx.draw(G, with_labels=True, node_size=1000, node_color='lightblue', font_size=8)
    plt.show()

def run():
    pre_local_sequences = [
        ['tx0', 'tx1', 'tx2'],
        ['tx2', 'tx3', 'tx4'],
        ['tx4', 'tx2']
    ]
    local_sequences = generate_node_orderings(pre_local_sequences)
    print("Node Orderings:")
    print(local_sequences)
    
    G = construct_graph(local_sequences)
    print("Graph Edges:")
    print(G.edges)
    
    visualize_graph(G)
    print("Topological Sort:")
    print(list(nx.topological_sort(G)))

def main():
    run()

if __name__ == '__main__':
    main()