Cli made using cobra to manage dns records for your domain.

Supports only godaddy.

To install for ubuntu- 
```
    wget https://github.com/ThunderGod77/domain-manager/releases/download/v0.1.4/domain-manager_0.1.4_Linux_x86_64.tar.gz
    
    sudo tar -C /usr/local/bin -xzf domain-manager_0.1.4_Linux_x86_64.tar.gz
    
    sudo chmod +x /usr/local/bin/domain-manager
```
To use the cli-
1. First configure the provider using the command
    ```
   domain-manaager new provider
   ```
   Then you would have to choose the domain name registrar(currently only supports godaddy) and then enter the access key and the secret.You can find them on https://developer.godaddy.com/getstarted

2. Configure the domain using the command
    ```
   domain-manager new domain
   ```
   Then you would have to enter the domain,select the domain registrar/provider and then add a description for the domain

3. You can then get the records using the command
      ```
      domain-manager get records --domain ${domain name}
   ```
4. To add a new record
      ```
   domain-manager new record --domain ${domain name}
   ```
5. To change an existing record
      ```
      domain-manager update record --domain ${domain name}
   ```
6. To delete an existing record
      ```
      domain-manager delete record --domain ${domain name}
   ```





To do:
- [ ] Add support for aws route 53
- [ ] Add support for namecheap
- [ ] To add tests