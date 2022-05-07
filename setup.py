from setuptools import setup, find_packages

setup(
        name ='ascii-tarot',
        version ='0.0.1',
        author ='Anastasia Lewis',
        author_email ='anastasia.m.lewis@gmail.com',
        url ='https://github.com/xnastasia/ascii-tarot',
        package_data={
          '': ['*.txt', '*.psd']
        },
        install_requires = ['psd-tools'],
        scripts=['bin/extract_text.py']
)
