import requests

h = input("Enter the website your want to check :")
dirs = input("Enter the dirs txt file you want to check PATH :")

try:
    with open(dirs) as f:
        for line in f:
         r = requests.get(h+line.strip())
         print(line.strip()+ '', r)
    
except requests.exceptions.HTTPError as error:
  print(error)