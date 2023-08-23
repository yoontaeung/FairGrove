import matplotlib.pyplot as plt
import networkx as nx
from utils import tx_generate

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
    
    return all_transactions, node_orderings

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
                # print(all_transactions[first], all_transactions[second])
                G.add_edge(all_transactions[first], all_transactions[second])
            elif first_counter < second_counter:
                # print(all_transactions[second], all_transactions[first])
                G.add_edge(all_transactions[second], all_transactions[first])
    
    return G

def visualize_graph(G):
    # Use spring layout algorithm to arrange nodes
    pos = nx.circular_layout(G)

    # Draw the graph with the calculated positions
    nx.draw(G, pos, with_labels=True, node_size=500, node_color='lightblue', font_size=8)

    # Display the graph
    plt.show()

def run():
    pre_local_sequences = tx_generate(100, 10)
    # pre_local_sequences = [
    #     ['tx1', 'tx3', 'tx2'],
    #     ['tx4', 'tx1', 'tx5'],
    #     ['tx3', 'tx4', 'tx2']
    # ]

    all_transactions, local_sequences = generate_node_orderings(pre_local_sequences)
    # print("Node Orderings:")
    # print(local_sequences)
    
    G = construct_graph(all_transactions, local_sequences)
    # print("Graph Edges:")
    # print(G.edges)
    # print(G)
    # visualize_graph(G)
    # print("Topological Sort:")
    # print(list(nx.topological_sort(G)))

    scc = list(nx.strongly_connected_components(G))
    # print("Strongly Connected Components:")
    # print(scc)
    
    condensed_G = nx.condensation(G, scc)
    # print("Condensed Graph Edges:")
    # print(condensed_G.edges)
    
    topological_order = list(nx.topological_sort(condensed_G))
    # print("Topological Sort on Condensed Graph:")
    # print(topological_order)

    sorted_transactions = []
    for condensed_node in topological_order:
        original_vertices = scc[condensed_node]
        sorted_transactions.extend(sorted(original_vertices))
    
    print("Sorted Transactions:")
    print(sorted_transactions)

def main():
    run()

if __name__ == '__main__':
    main()