# TAKE HOME ASSIGNMENT 2

### Write a Shred(path) function in Go.

#### Alejandro Cendejas

# ABOUT Shred()

This function will overwrite the contents of a file with random data three times, and then 
delete the file.

If path does not exist, or is a directory, it will return the corresponding error.

If the user does not own path, it will return an error.

If path is a link, it will overwrite the target's content with random data three times, but 
it will not delete the target, just the link.

# USES

This is a useful tool for the safe deletion of sensitive information from a file system. It will ensure that random data is written to the same blocks three times, and then delete the file. 

Use with caution with SSDs since it may shorten its lifetime if used very frequently. A 
better alternative for SSDs might be to check the manufacturer's recommendations for 
safe erasing of data.

# NEXT STEPS 

Write tests iin order to ensure the proper behavior:

### Tests for writeRandomData

1. Pass a path which does not exist, get the corresponding error.
2. Pass a path which is a directory, get the corresponding error.
3. Pass a path for a file without write permissions for the current user/group, get the corresponding error.
4. Pass a path for a file owned by the current user, confirm it is overwritten.
5. Pass a path for a link, confirm the target is overwritten.

### Tests for Shred

1. Pass a path which does not exist, get the corresponding error.
2. Pass a path which is a directory, get the corresponding error.
3. Pass a path for a file without write permissions for the current user/group, get the corresponding error.
4. Pass a path for a file owned by the current user, confirm it is deleted.
5. Pass a path for a link, confirm the link is deleted and the target is overwritten.
