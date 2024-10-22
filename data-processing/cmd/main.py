from internal.clustering import clusteringMap
from internal.text_analysis import textAnalysisMap
from internal.panda import cleanFile
import os

def main():
    filename, columns, nlpAndTextAnalysisList = validateInputs()
    df = cleanFile(filename, columns)
    executeNlpAndTextAnalysis(df, columns, nlpAndTextAnalysisList)

def validateInputs():
    filename =  os.getenv('FILENAME', 'groups.csv')
    if filename == '':
        raise ValueError("file name is needed to be obtained")

    columns = os.getenv('COLUMNS', 'description, name')
    if columns == '':
        raise ValueError("at least one column is needed to be processed")

    columns = [col.strip() for col in columns.split(',')]

    nlpAndTextAnalysisList = os.getenv('nlpAndTextAnalysisList', 'spacy')
    if nlpAndTextAnalysisList == '':
        raise ValueError("at least one nlp and textAnalysis is needed to be used")

    nlpAndTextAnalysisList = [col.strip() for col in nlpAndTextAnalysisList.split(',')]
    return filename, columns, nlpAndTextAnalysisList

def executeNlpAndTextAnalysis(df, columns, nlpAndTextAnalysisList):
    results = {libName: [] for libName in nlpAndTextAnalysisList}
    for lib in nlpAndTextAnalysisList:
        textAnalysisMap.get(lib)(df, columns)
    return results

if __name__ == "__main__":
    main()