# Remove files from the git history

sometimes I encountered I added/committed and pushed file(s) to GitHub. 
Then I have to remove files from the git history from the local and remote git history


## Remove from the local git history
```
$ gitt filter-branch --force --index-filter \
  "git rm --cached --ignore-unmatch git filter-branch --force --index-filter \
  "git rm --cached --ignore-unmatch INSERT_PATH_TO_FILE_HERE" \
  --prune-empty --tag-name-filter cat -- --all" \
  --prune-empty --tag-name-filter cat -- --all
```
for example, `INSERT_PATH_TO_FILE_HERE` will be like a `bin/test.exe`

## Then, Push it to the remote repository

```
$ git push origin --force --all
```

[Reference: How to remove sensitive files from GitHub commit history?](https://itnext.io/how-to-remove-sensitive-files-from-github-commit-history-638a3f291f74)

