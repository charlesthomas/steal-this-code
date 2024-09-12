It's often useful to be able to do a shallow clone (eg `git clone --depth 1`), but when you do, `git` gets weird about pushing.

These are my notes on how to push from a shallow clone. I suspect not 100% of these commands are required, but I can't be bothered to try this over and over again leaving one / some of them out to confirm that.

# Cloning

```bash
git clone --depth 1 $repo_url
git remote set-branches origin $default_branch
git fetch --depth 1 origin $default_branch
git checkout $default_branch
```

# Pushing

```bash
git checkout -b $new_branch
git push origin $new_branch
git remote set-branches origin $new_branch
git fetch --depth 1 origin $new_branch
git branch --set-upstream-to origin/$new_branch
git add -A
git commit -m $message
git push
```
