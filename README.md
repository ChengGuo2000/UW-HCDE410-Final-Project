# Cheng Guo - HCDE 410 Final Project - May 2023

## Abstract
This project is about building machine-learning models that can predict the edibility of mushrooms based on their appearance. The reason for doing this project is that I have friends who live in areas where mushrooms are abundant, and I am interested in knowing how to select edible mushrooms based on their appearance in the wild, so I want to build this model based on a data set I found. I used data provided by mushroom researcher Dennis Wagner, which contains 17 nominal appearance and 3 quantitative appearance, and edibility information for 61069 mushroom species. All data are generated from information on 173 species of mushrooms. Overall, I have explored the relationship between mushroom edibility and its appearance, and I found that size and color are not the deterministic features of edibility, with larger mushrooms tending to be edible, contradicting our hypothesis. During data modeling, models with more features tend to perform better than models with only size and color features. In future research on mushrooms, we should consider more characteristics, including habitat and season, so that we can better predict mushroom edibility. Additionally, mushrooms that look similar tend to have the same edibility, so when building a model, k nearest neighbors would be a better choice.

## The Goal of the Project
I am exploring the relationship between mushroom's appearance with whether it is poisonous or not. The project consists of a data exploration phrase, where I make visualizations of the mushroom appearances data, and a data modeling phrase, where I build a model that can predict whether a mushroom is poisonous or not from its appearances. I am using the data on 61069 mushrooms with 17 nominal and 3 quantitative appearance features and a target (edible or poisonous) provided from https://archive.ics.uci.edu/ml/datasets/Secondary+Mushroom+Dataset

## List of the License of the Source Data and the Project
1. The data provided from https://archive.ics.uci.edu/ml/datasets/Secondary+Mushroom+Dataset has a Creative Commons LICENSE, and the license file can be accessed here: https://mushroom.mathematik.uni-marburg.de/files/. 
2. This project has a MIT LICENSE, which can be accessed from the "LICENSE" file in this folder.

## Final Files
1. Final_Project_Report.ipynb: The main file including all the codes
2. MushroomDataset folder: The folder contains all the data files from the above source. The data I used for this project is in the secondary_data.csv file in this folder.
3. Generated_Graphics folder: All the 12 data visualizations generated from the code in the Final_Project_Report.ipynb file.
4. LICENSE: The MIT LICENSE of this project.
5. README.md: This file.

## Links to all the References in the Assignment
1. Wagner, D., Heider, D., & Hattab, G. (2021). Mushroom data creation, curation, and simulation to support classification tasks. *Scientific Reports, 11*.
2. Schlimmer, J.S. (1987). Concept Acquisition Through Representational Adjustment (Technical Report 87-19). Doctoral disseration, Department of Information and Computer Science, University of California, Irvine.
3. Stack Overflow. (2017, August 30). Plotting two dictionaries in one bar chart with Matplotlib. Stack Overflow. https://stackoverflow.com/questions/45968359/plotting-two-dictionaries-in-one-bar-chart-with-matplotlib 
4. Kumar, A. (2023, April 19). Histogram plots using Matplotlib &amp; Pandas: Python. Data Analytics. https://vitalflux.com/histogram-plots-using-matplotlib-pandas-python/ 
5. scikit-learn developers. (2023). Sklearn.metrics.confusionmatrixdisplay. scikit. https://scikit-learn.org/stable/modules/generated/sklearn.metrics.ConfusionMatrixDisplay.html 
