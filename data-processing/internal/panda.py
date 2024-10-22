import pandas as pd
import os

def cleanFile(filename, columns):
    df = None
    try:
        name = '../files/' + filename
        df = pd.read_csv(name)
    except Exception as e:
        print("an error occurred loading file, remember to add <filename>.<extension> for example groups.csv: [ERROR]: ", e)

    try:
        for column in columns:
            df[column] = df[column].str.replace(r'[^\w\s]', '', regex=True)

        return df.drop_duplicates(subset=columns, keep='first')
    except Exception as e:
        print("error occurred cleaning columns {columns}, [ERROR]: ", e)

