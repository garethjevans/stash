# stash

name : String
Name of a stash. Should be a simple identifier akin to a job name.
allowEmpty : boolean (optional)
Create stash even if no files are included. If false (default), an error is raised when the stash does not contain files.
excludes : String (optional)
Optional set of Ant-style exclude patterns.
Use a comma separated list to add more than one expression.
If blank, no files will be excluded.
includes : String (optional)
Optional set of Ant-style include patterns.
Use a comma separated list to add more than one expression.
If blank, treated like **: all files.
The current working directory is the base directory for the saved files, which will later be restored in the same relative locations, so if you want to use a subdirectory wrap this in dir.
useDefaultExcludes : boolean (optional)
If selected, use the default excludes from Ant - see here for the list. Defaults to true.



