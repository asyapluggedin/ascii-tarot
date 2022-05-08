#!/usr/bin/env python

import glob
import argparse
from psd_tools import PSDImage
from psd_tools.api.layers import TypeLayer
from pprint import pprint

parser = argparse.ArgumentParser(description='Process PSD files and extract text layers')
parser.add_argument('--input_dir', help='path to PSD file directory')
parser.add_argument('--output_dir', help='path to write txt outfiles')

if __name__ == "__main__":
    args = parser.parse_args()
    files = glob.glob(args.input_dir + '*.psd')
    for name in files:
        psd = PSDImage.open(name)
        for layer in psd:
            if isinstance (layer, TypeLayer):
               filename = name.split('.')[0].strip().lower().replace(' ', '_')
               out = open(args.output_dir + filename, 'w')
               text = str(layer._data.text_data.get(b'Txt ')).replace('\\r','\n').strip("'").rstrip('\x00')
               out.write(text)
