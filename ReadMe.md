##ocdJanus Utilities

1) Put all selects into a file for later use...

grep SELECT inputStage1/janus_structs_20151023.txt > inputStage1/selectStatements.txt

2) Remove all the SELECT commands with:

sed  '/SELECT/d' janus_structs_20151023.txt > janus_structs_20151023_noselect.txt

3) Split the struct file with 

awk -f awkcode.awk structdata.txt 
awk -f ../awkcode.awk ../inputStage1/janus_structs_20151023_noselect.txt

4) clean up blanks and spaces

find . -name "part*" -exec sed -i '/^[[:space:]]*$/d;s/[[:space:]]*$//' {} \;
find . -name "part*" -exec sed -i ".bck" '/^ *$/d' {} \;

5) the rename.sh renames the partX.txt files based on the struct nameint he file
./rename.sh 

6) Next we remove all likes that container a { or a }

$ find . -name "Janus*" -exec sed -i '/}/d' {} \;                                                                                                                    
$ find . -name "Janus*" -exec sed -i '/{/d' {} \; 

