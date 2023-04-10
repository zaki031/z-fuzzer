import requests
import os
from termcolor import colored
os.system('color')
# asking for the website
h = input("Enter the website your want to check :")+'/'

#asking for the txt file path
dirs = input("Enter the dirs txt file you want to check PATH :")

try:
    #opening the txt fule
    with open(dirs) as f:
        # looping into it
        for line in f:
         #testing requests to the url
         r = requests.get(h+line.strip())
         #checking the request status code
         if r.status_code == 200:
            print('/'+line.strip()+ '', colored(r,'green')  ,'     |  size: ',len(r.content) , 'bytes')
         else:
            print('/'+line.strip()+ '', colored(r,'red')  ,'     |  size: ',len(r.content) , 'bytes')


except requests.exceptions.HTTPError as error:
  print(error)