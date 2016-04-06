# text-replace
Plain text string substitution with dictionary.

Takes a target plain text file and a dictionary plain text file "CSV"
and replaces all ocurrences found in the plain text from the dictionary.
One per line.

## Dictionary
The dictionary should have the form:

```
<target-string><separator><new-string>
<target-string><separator><new-string>
<target-string><separator><new-string>
...
<target-string><separator><new-string>
```

### example:

```csv
Hello,Hola
Hello,Hola
Hello,Hola
Hello,Hola
```

The default separator is TAB('\t') others must be specified on the
LoadDictionary function.
