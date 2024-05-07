# Ansible Notes (NO)



### To pass password (*not recommended*)

Updated `ansible.cfp` 

My user is `user`

```shell
become_allow_same_user=True
become=True
become_user=root
remote_user=user
```



Now use this command:

```shell
# Store your password:
export PASS="MyPass"

# Launch this command
ansible-playbook wapi.yml --extra-vars "ansible_password=$PASS"
```

