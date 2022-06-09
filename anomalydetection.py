# python outlier detection
# pip3 install pyod
from cProfile import label
from calendar import c
import csv
import warnings
import numpy as np
import pandas as pd
import random
from pyod.models.mad import MAD
from pyod.models.knn import KNN
from pyod.models.lof import LOF
import matplotlib.pyplot as plt
from sklearn.ensemble import IsolationForest

# read information data
informations = []
with open('information.csv') as csv_file:
    csv_reader = csv.reader(csv_file, delimiter=',')
    for x in csv_reader:
        informations.append(x)

#read from csv
usage = []
with open('usage.csv') as csv_file:
    csv_reader = csv.reader(csv_file, delimiter=',')
    for row in csv_reader:
        for x in range(1, len(row)):
            row[x]=int(row[x])
        usage.append(row)

print(usage)
size = ['time']+informations[0][1:]
sizeList = informations[0][1:]
data = pd.DataFrame(usage , columns=size)
print("data",data)

def fit_model(model, data, column=sizeList):
    for x in sizeList:
        # fit the model and predict it
        column = x
        df = data.copy()
        data_to_predict = data[column].to_numpy().reshape(-1, 1)
        predictions = model.fit_predict(data_to_predict)
        df['Predictions'] = predictions
    
    return df

def plot_anomalies(df, x='time', y=sizeList):
    f = plt.figure(figsize=(24, 10))
    colors = ["black", "blue", "yellow", "purple", "magenta", "gray", "pink", "cyan"]
    for value in sizeList:
        pickedColor = []
        y = value
        # categories will be having values from 0 to n
        # for each values in 0 to n it is mapped in colormap
        categories = df['Predictions'].to_numpy()
        colormap = np.array(['g', 'r'])
        f = plt.title("Kyverno Automate Performance Test result")
        f = plt.scatter(df[x], df[y], c=colormap[categories])
        f = plt.xlabel(x)
        f = plt.ylabel("usage")
        f = plt.xticks(rotation=90)
        plt.plot(df[x], df[y], c = random.choice(colors), linestyle='solid', label=value)
        pickedColor.append(c)
        #plt.show()
    txt="Scales: "+informations[0][0]+"\n Kyverno restart count:"
    plt.figtext(0.5, 0.01, txt, wrap=True, horizontalalignment='center', fontsize=12)
    plt.legend(loc='upper left')
    plt.savefig("report.png")
    print("pickedcolor", pickedColor)

def find_anomalies(value, lower_threshold, upper_threshold):
    
    if value < lower_threshold or value > upper_threshold:
        return 1
    else: return 0

def iqr_anomaly_detector(data, column='amount', threshold=1.1):
    
    df = data.copy()
    quartiles = dict(data[column].quantile([.25, .50, .75]))
    quartile_3, quartile_1 = quartiles[0.75], quartiles[0.25]
    iqr = quartile_3 - quartile_1

    lower_threshold = quartile_1 - (threshold * iqr)
    upper_threshold = quartile_3 + (threshold * iqr)

    print(f"Lower threshold: {lower_threshold}, \nUpper threshold: {upper_threshold}\n")
    
    df['Predictions'] = data[column].apply(find_anomalies, args=(lower_threshold, upper_threshold))
    return df
  

#Isolation forest
iso_forest = IsolationForest(n_estimators=125)
iso_df = fit_model(iso_forest, data)
iso_df['Predictions'] = iso_df['Predictions'].map(lambda x: 1 if x==-1 else 0)
plot_anomalies(iso_df)