# getgithub
Get specific file or directory from the a github repo without cloning.

How to Install/Get
------------------

    $ go get github.com/praveenkumar/getgithub
    $ export PATH=$PATH:$GOPATH/bin
    $ getgithub -h
     Usage of ./getgithub:
       -branch string
         	branch or tag (default "master")
       -dest string
         	Destination directory (default "/Users/prkumar/work/github/go_practice/src/github.com/praveenkumar/getgithub/out")
       -dir string
         	Directory or file to download (default "/")
       -list
         	List the directory/file for provided Path
       -owner string
         	Name of the repo owner
       -repo string
         	Name of the repo

How to Use
----------

Get the list of file and directories from the top of repository.

     $ ./getgithub -repo minishift -owner minishift -list
    file 	 .gitignore 
    file 	 Gopkg.lock 
    file 	 Gopkg.toml 
    file 	 LICENSE 
    file 	 Makefile 
    dir 	 cmd 
    dir 	 test 
    dir 	 .circleci 
    file 	 .travis.yml 
    file 	 README.adoc 
    file 	 ROADMAP.adoc 
    file 	 centos_ci.sh 
    dir 	 addons 
    dir 	 scripts 
    dir 	 .github 
    file 	 .gitlab-ci.yml 
    file 	 CONTRIBUTING.adoc 
    file 	 appveyor.yml 
    dir 	 docs 
    file 	 gen_help_text.go 
    dir 	 pkg 
    
Get the list of files and directories from specified directory.

    $ ./getgithub -repo minishift -owner minishift -dir /pkg -list
    dir 	 pkg/version 
    dir 	 pkg/minikube 
    dir 	 pkg/minishift 
    dir 	 pkg/testing 
    dir 	 pkg/util

Get the content of the file and directories from specified directory keeping directory tree intact.

    $ ./getgithub -repo minishift -owner minishift -dir /pkg/minishift -dest /tmp/
    $ ls /tmp/pkg/minishift/
    addon		cluster		config		docker		network		openshift	provisioner	shell		update
    cache		clusterup	constants	hostfolder	oc		profile		registration	systemd		util

           
How to build
------------
 
    $ make build
    
Note
----

If you get rate limit error then set `GH_TOKEN` to environment variable.

`$ export GH_TOKEN=<my_token>`

- Github Token: https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/                 