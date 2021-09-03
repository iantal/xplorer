# Design notes

* tool that creates dependency trees (gradle, maven, sbt, npm, etc)
  - input: repository
  - output: file.gradle, file.maven, etc (text file with dependency tree)

* tool to extract libraries
  - input: dependency tree file
  - output: file that contains list of libraries
  - ? what is both npm and gradle are used -> produce a csv file to specify the lib + build tool

* tool to download libraries from various repositories (maven central, npm, etc)
  - input: list of libraries
  - output list of downloaded libraries in a specified directory

* tool to analyze vulnerabilities for the given libraries

* tool that matches library with a particular file in the project

* tool to view results



xplorer --repository /opt/data/project --output /tmp
