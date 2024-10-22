import spacy

def spacyLib(df, columns):
    nlp = spacy.load('es_core_news_sm')
    documents = {column: [] for column in columns}

    for index, row in df.iterrows():
        for column in columns:
            doc = nlp(row[column])
            documents[column].append(doc)

    for column in columns:
        print(f"\nAnálisis para la columna: {column}")
        for doc in documents[column]:
            print(f'Texto: {doc.text}')

            # Análisis de tokens
            for token in doc:
                print(f'Token: {token.text}, Lemma: {token.lemma_}, Parte del discurso: {token.pos_}')

            # Extracción de entidades
            print("Entidades reconocidas:")
            for ent in doc.ents:
                print(f'Entidad: {ent.text}, Etiqueta: {ent.label_}')

            print("-" * 40)  # Separador entre documentos
    return "resultados de spacy....."

def nltk():
    return "funcion NLTK"

def transformers():
    return "funcion Transformers"

def vader():
    return "funcion VADER"

def gensim():
    return "funcion Gensim"

textAnalysisMap = {
    'spacy': spacyLib,
    'nltk': nltk,
    'transformers': transformers,
    'vader': vader,
    'gensim': gensim
}