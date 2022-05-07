# How to run extraction script
python3 -m venv
source venv/bin/activate
python3 setup.py install
python3 bin/extract_text.py $INPUT_DIR $OUTPUT_DIR

