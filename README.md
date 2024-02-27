<h1>cligpt</h1>

<h1>Installation</h1>

Go 1.17+
```
$ go install github.com/drpaij0se/cligpt@latest
```

test

```
❯ cligpt "Do a poem about Golang in Latin"


Nunc est tempus, ut in Golang gradimur
Mox nostris verbis linguam didicimus
Dum in codice nos ponimus,
Cuncta sic apte nos iungimus.

Verba nova nos scrutamur
Et in praxi illa provamus
Ut optimus sic codex fiat,
Hic scimus nos invenire quod quaerimus.

Golang saepe gloriosa est,
Mentem nostram stimulat et excitat
Nobis efficit ut mira creemus
Et magna in mundo fata sequamur.
```

<h1>Configuration</h1>
Run:

```
$ ./cligpt "create the config directory cuh"
```

- You are going to see a error message, so, insert the OpenAI token in your config directory that is located in `$HOME/.config/cligpt/cligpt.yml`

```haskell
auth: token
model: text-davinci-003
max_tokens: 256
```

<h1>Building</h1>

- Clone the repository.

`$ git clone https://github.com/drpaij0se/cligpt`

- Build the cli.

`./build.sh`

![image](https://user-images.githubusercontent.com/69026987/209194859-a2456a7d-796f-47e0-8a8e-062848e2cbaf.png)
# cligpt
