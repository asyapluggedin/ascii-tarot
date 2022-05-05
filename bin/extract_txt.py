import glob
from psd_tools import PSDImage
from psd_tools.api.layers import TypeLayer
from pprint import pprint

if __name__ == "__main__":
    files = glob.glob('*.psd')
    for name in files:
        psd = PSDImage.open(name)
        for layer in psd:
            if isinstance (layer, TypeLayer):
               filename = name.split('.')[0].strip().lower().replace(' ', '_')
               out = open(filename, 'w')
               text = str(layer._data.text_data.get(b'Txt ')).replace('\\r','\n')
               out.write(text)
