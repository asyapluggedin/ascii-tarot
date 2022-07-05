# ASCII TAROT 
## Description

I started this project on May 2, 2016 as part of my creative portfolio to apply to [UCLA Design | Media Arts](ucla.dma.edu). You can read a bit about the motivations either on the projects [Behance](https://www.behance.net/gallery/42369763/ASCII-TAROT?locale=en_US) page, or in the original [manifesto]() I drafted back then.

In 2022, I revisited my work on the project to turn it into a CLI tarot reading app. 


## Progress Tracking

This section will eventually become a separate .md file that anybody can access to see how my workflow was going during the making of this program.

#### To Do List
- [ ] Add .gitignore
- [ ] write specs for simple interaction user flow
- [ ] look at intro to Go (~1hr) grok types and pointers frsure

### 06-05-2022
- Decided on using tcell (confidence level re: decision: 75%)
- Committed to writing daily progress logs in the README


## Installation
 ` lol `

## Usage
 ` lol `

## Collaborators

Benjamin Bascom wrote the extraction script that ripped the text elements from the .psd's I made in 2016.

Mischa Snook helped me decide which technologies to use for the project, and paired with me during the programming process.

## Physical Deck

to come ;)







##  How to run extraction script
` pip install psd_tools
python3 bin/extract_text.py --input_dir $INPUT_DIR --output_dir $OUTPUT_DIR `
e.g. 
` python3 bin/extract_txt.py --input_dir img/psd/ --output_dir txt/extracted_txt/ `


