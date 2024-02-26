from gensim import corpora, models
import sys
import json
import csv

def perform_lda(documents):
    processed_docs = [doc.split() for doc in documents]
    dictionary = corpora.Dictionary(processed_docs)
    corpus = [dictionary.doc2bow(doc) for doc in processed_docs]
    lda_model = models.LdaModel(corpus, num_topics=20, id2word=dictionary, passes=1)
    topics = lda_model.print_topics(num_words=80)
    return topics

def save_topics_to_csv(topics, filename='lda_topics.csv'):
    # Abrindo o arquivo CSV para escrita
    with open(filename, 'w', newline='', encoding='utf-8') as csvfile:
        writer = csv.writer(csvfile)
        # Escrevendo o cabeçalho do CSV
        writer.writerow(['Topic Index', 'Words and Weights'])
        # Escrevendo os tópicos no arquivo CSV
        for topic in topics:
            writer.writerow([topic[0], topic[1]])

if __name__ == '__main__':
    documents = json.loads(sys.stdin.read())
    topics = perform_lda(documents)
    # Salvando os tópicos em um arquivo CSV
    save_topics_to_csv(topics)
    # Opcional: Retorna os tópicos como JSON na saída padrão
    print(json.dumps(topics))
