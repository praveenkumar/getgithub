# getgithub
Get specific file or directory from the a github repo without cloning.

How to Install/Get
------------------

::

    $ go get github.com/praveenkumar/getgithub
    $ export PATH=$PATH:$GOPATH/bin
    $ getgithub -h
      -branch string
            branch or tag (default "master")
      -dest string
            branch or tag (default "/Users/prkumar/work/github/go_practice/src/github.com/praveenkumar/getgithub")
      -dir string
            Directory or file to download
      -owner string
            Name of the repo owner
      -repo string
            Name of the repo
            
How to build
------------

:: 
    
    $ make build
    
Note
----

If you get rate limit error then set `GH_TOKEN` to environment variable.

`$ export GH_TOKEN=<my_token>`

- Github Token: https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/                 