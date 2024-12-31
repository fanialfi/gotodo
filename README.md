# GOTODO task tracker

Sample solution for the [task-tracker](https://roadmap.sh/projects/task-tracker) chalenge from [roadmap.sh](https://roadmap.sh)

## how to run

clone the repository adn run the following command :

```zsh
[user@hostname] $ git clone https://github.com/fanialfi/gotodo.git
[user@hostname] $ cd gotodo
```

run the following command to build and run the project

```zsh
[user@hostname] $ go build -o gotodo .
[user@hostname] $ ./gotodo --help # to see the list of available commands

# to add a task
[user@hostname] $ ./gotodo add --description "buy coffie"
[user@hostname] $ ./gotodo add -d "buy coffie" # shorthand

# to update a task
[user@hostname] $ ./gotodo update --id <id task> --description <new description task>

# to delete a task
[user@hostname] $ ./gotodo delete --id <id task>
[user@hostname] $ ./gotodo delete -i <id task> # shortand

# mark a task is in-progress
[user@hostname] $ ./gotodo mark-in-progress --id <id task>

# mark a task is done
[user@hostname] $ ./gotodo mark-done --id <id task>

# list all task
[user@hostname] $ ./gotodo list

# list alll task that are done
[user@hostname] $ ./gotodo list --done

# list all task that are todo
[user@hostname] $ ./gotodo list --todo

# list all task that are in-progress
[user@hostname] $ ./gotodo list --in-progress
```
