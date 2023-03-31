import requests
h = input("Enter the website your want to check :")
with open('list.txt') as f:
    for line in f:
        r = requests.get(h+line.strip())
        print(line.strip()+ '', r)
    