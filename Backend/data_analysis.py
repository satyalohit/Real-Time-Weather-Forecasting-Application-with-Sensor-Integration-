import requests
import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestRegressor
from sklearn.metrics import mean_squared_error

# API URL
api_url = 'https://live-weather-data-analysis-b6491cde3287.herokuapp.com/sensor'

# Fetch data from API
response = requests.get(api_url)
data= response.json()

df=pd.DataFrame(data['data'])

# Exclude records where Temp or Humidity is zero
df = df[(df['Temp'] != 0) & (df['Humidity'] != 0)]
# Convert Date and TimeOfDay to datetime and set as index
df['DateTime'] = pd.to_datetime(df['Date'] + ' ' + df['TimeOfDay'])
df = df.set_index('DateTime')
df.drop(['Date', 'TimeOfDay'], axis=1, inplace=True)


df['Hour'] = df.index.hour

# Split data into features and target for temperature and humidity
X = df[['Hour']]  # Only using 'Hour' as feature
y_temp = df['Temp']
y_humidity = df['Humidity']

# Split the data into training and test sets
X_train_temp, X_test_temp, y_train_temp, y_test_temp = train_test_split(X, y_temp, test_size=0.2, random_state=0)
X_train_humidity, X_test_humidity, y_train_humidity, y_test_humidity = train_test_split(X, y_humidity, test_size=0.2, random_state=0)

# Train Random Forest models
model_temp = RandomForestRegressor(random_state=0)
model_temp.fit(X_train_temp, y_train_temp)
model_humidity = RandomForestRegressor(random_state=0)
model_humidity.fit(X_train_humidity, y_train_humidity)

# Create a DataFrame for predictions (every hour of the current day)
current_date = pd.to_datetime("today").date()
prediction_hours = pd.date_range(start=f"{current_date} 00:00", end=f"{current_date} 23:00", freq='H')
prediction_df = pd.DataFrame(prediction_hours, columns=['DateTime'])
prediction_df['Hour'] = prediction_df['DateTime'].dt.hour

# Make predictions
prediction_df['Predicted_Temperature'] = model_temp.predict(prediction_df[['Hour']])
prediction_df['Predicted_Humidity'] = model_humidity.predict(prediction_df[['Hour']])

# Output the results

print(prediction_df[['DateTime', 'Predicted_Temperature', 'Predicted_Humidity']])

