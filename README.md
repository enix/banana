# The banana project

Complete documentation is available here :
https://confluence.enix.org/pages/viewpage.action?pageId=31621154


## Setting up a working dev workspace

#### Create your .env file

Example `.env`:

```ini
API_ENDPOINT=https://object-storage.r1.nxs.enix.io
API_ACCESS_TOKEN=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
API_SECRET_TOKEN=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

#### Start the stack

```bash
docker-compose up -d
```

#### Generate the certs

There's a script that can do that for you. If you're lazy you can one-shot the config by adding the `--auto` switch.

```bash
export BANANA_COMPANY_NAME="bananagency"
./scripts/init_dev.sh $BANANA_COMPANY_NAME --auto
```

#### Trust the root CA

```bash
open security/ca/ca.pem
```

Select 'Always trust' in keychain or the page won't load!

#### Add hosts to /etc/hosts

Add the following line to `/etc/hosts`

```
127.0.0.1	banana.enix.io api.banana.enix.io
```

#### Issue and trust a client cert for your user

```bash
./scripts/init_user.sh $BANANA_COMPANY_NAME "king.kong"
open security/out/king.kong.p12
```

#### Issue a client cert for your agent

```bash
./scripts/init_agent.sh $BANANA_COMPANY_NAME "the.agent"
```

#### Test your setup

- For user access, open your browser on [https://api.banana.enix.io:7443](https://api.banana.enix.io:7443/ping)
- For agent access:

```bash
echo -n '{"foo": "bar"}' > payload.json
./scripts/sign_and_post.sh '/ping' payload.json 'the.agent'
```
