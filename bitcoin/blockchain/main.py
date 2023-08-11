import sys
import requests

address = '1GdK9UzpHBzqzX2A9JFP3Di4weBwqgmoQA'
# address = '1Cdid9KFAaatwczBwBttQcwXYCpvK8h7FK'

address = sys.argv[1]
res = requests.get('https://blockchain.info/unspent', params={'active': address})

data = res.json()

print(address)
print('script,tx_index,tx_hash,tx_output_n,value')

s = 0.0
for d in data['unspent_outputs']:
    asm = d['script']
    val = d['value']
    s += val/100000000.0
    print('{},{},{},{},{}'.format(d['script'], d['tx_index'], d['tx_hash'], d['tx_output_n'], d['value']/100000000.0))
    
print(s)




