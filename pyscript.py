import json

# the file to be converted to 
# json format
filename = 'test.txt'
  
# dictionary where the lines from
# text will be stored
dict1 = {}

# fields in the sample file 
fields =['job_name', 'curent_job', 'log']

# p = 0
# creating dictionary
with open(filename) as fh:
    line_no = 0 
    for line in fh:
        
        # reads each line and trims of  the extra spaces 
        # and gives only the valid words
        description = list( line.strip().split('\t', 3))
        # print(description)
        # loop variable

        if len(description) < len(fields):
            dict1[line_no] = description
            line_no += 1
            continue
        
        i = 0
        # intermediate dictionary
        dict2 = {}
        while i<len(fields):
              
                # creating dictionary for each employee
                dict2[fields[i]]= description[i]
                i = i + 1
                  
        # appending the record of each employee to
        # the main dictionary
        dict1[line_no]= dict2
        # print(dict1)
        line_no += 1
        # p += 1

        # if p > 5:
        #     break
  
# creating json file
# the JSON file is named as test1

out_file = open("test_json.json", "w")
json.dump(dict1, out_file, indent = 4, sort_keys = False)
out_file.close()